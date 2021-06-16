package request

import (
	"fmt"
	"forum/database"
	"forum/toolbox"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TicketData struct {
	User    database.User
	Tickets []database.Ticket
	Bool    int
}

func Ticket(w http.ResponseWriter, r *http.Request, user database.User) {
	var data TicketData
	switch r.Method {
	// 🍔 Méthode 'GET' — Lorsqu'on arrive sur la page :
	case "GET":
		particular, _ := r.URL.Query()["id"]
		if particular != nil {
			data.Bool = 1
			id, _ := strconv.Atoi(particular[0])
			data.Tickets = append(data.Tickets, database.GetTicketByID(id))
			database.Db.QueryRow("SELECT username FROM users WHERE id=?", data.Tickets[0].Author_id).Scan(&data.Tickets[0].Author_name)
		} else {
			if user.Role < database.ADMIN {
				data.Tickets = database.GetTicketByUserID(user.ID)
			} else {
				data.Tickets = database.GetAllTickets()
			}
			data.Bool = 0
		}
		data.User = user

		err := MyTemplates.ExecuteTemplate(w, "ticket", data)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Println("❌ ERREUR | Impossible d'exécuter le template ticket.")
			fmt.Println(err)
			return
		}
	case "POST":
		id := r.FormValue("id")
		database.ResolveTicket(id)
		if user.Role < database.ADMIN {
			data.Tickets = database.GetTicketByUserID(data.User.ID)
		} else {
			data.Tickets = database.GetAllTickets()
		}
		err := MyTemplates.ExecuteTemplate(w, "ticket", data)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Println("❌ ERREUR | Impossible d'exécuter le template ticket.")
			fmt.Println(err)
			return
		}
	}
}

func NewTicket(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {

	// 🍔 Méthode 'GET' — Lorsqu'on arrive sur la page :
	case "GET":
		var data database.DataForNewTicket
		data.User = user

		err := MyTemplates.ExecuteTemplate(w, "newticket", data)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Println("❌ ERREUR | Impossible d'exécuter le template “newpost”.")
			return
		}
	case "POST":
		// (1) Récupération du titre, du contenu, et de la catégorie du post à publier :
		title := r.FormValue("title")
		content := r.FormValue("content")
		if toolbox.IsEmptyString(title) || toolbox.IsEmptyString(content) {
			log.Println("❌ POST | Impossible de publier le post : le titre ou le contenu est vide.")
			err := MyTemplates.ExecuteTemplate(w, "400", user)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
			// http.Error(w, "400 Bad Request\nThe text you added is empty.", http.StatusBadRequest)
			// return
		}
		// (3) Remplissage d'une struct 'Post' pour le post à publier :
		var ticket database.Ticket

		ticket.Title = title
		ticket.Author_id = user.ID
		ticket.Actual_Admin = 0
		ticket.Content = content
		ticket.Date = time.Now()
		ticket.State = database.UNRESOLVED

		// (4) Insertion du post dans la base de données :
		ticketID, err := ticket.InsertIntoDatabase() // La méthode d'insertion dans la DB renvoie l'ID du post qui vient d'être inséré
		if err != nil || ticketID < 1 {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		// (5) Redirection vers la page du post :
		postURL := "/ticket?id=" + strconv.Itoa(ticketID)
		http.Redirect(w, r, postURL, http.StatusSeeOther)
	}
}

func Ticket_Answer(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {
	case "POST":
		var received receivedData
		received.Value = r.FormValue("answer")
		received.ID = r.FormValue("id")
		if len(received.Value) > 1 {
			_, err := database.Db.Exec("INSERT INTO ticket_answers (ticket_id,author_id,author_name,content) VALUES(?,?,?,?)", received.ID, user.ID, user.Username, received.Value)
			if err != nil {
				panic(err)
			}
			if user.Role >= 3 {
				database.Db.Exec("UPDATE ticket_answers SET actual_admin = ? WHERE id = ?", user.ID, received.ID)
			}
		}
		postURL := "ticket?id=" + received.ID
		http.Redirect(w, r, postURL, http.StatusSeeOther)
	}
}
