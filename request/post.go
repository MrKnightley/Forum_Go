package request

import (
	"fmt"
	"forum/database"
	"forum/toolbox"
	"log"
	"net/http"
	"strings"
	"time"
)

func Post(w http.ResponseWriter, r *http.Request, user database.User) {

	// (1) Récupération de l'ID du post depuis l'URL :
	ID, err := toolbox.ParseURL(w, r)

	if err != nil || ID < 1 {
		err := MyTemplates.ExecuteTemplate(w, "404", user)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		// http.Error(w, "404 NOT FOUND", http.StatusNotFound)
		return
	}
	switch r.Method {

	// 🍔 Méthode 'GET' (arrivée sur l'URL /𝐩𝐨𝐬𝐭/{𝐈𝐃}) :
	case "GET":
		// (2) Remplissage d'une struct Data pour chaque page de post :

		/*  DataForPost struct {
			- User     User
			- Post     Post
			- Comments []Comment
		} */

		var dataForPost database.DataForPost

		dataForPost.User = user
		dataForPost.Post, err = database.GetPostByID(ID, user.ID) // (ID du post, ID de l'utilisateur loggé)
		if dataForPost.Post.State == 0 {
			dataForPost.Comments, err = database.GetCommentsByPostID(ID, user.ID)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				log.Println("❌ ERREUR | Impossible de récupérer le post ou les commentaires du post dont l'ID est ", ID)
				return
			}
		}
		err = MyTemplates.ExecuteTemplate(w, "post", dataForPost)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Println("❌ ERREUR | Impossible d'exécuter le template “post”.")
			fmt.Println(err)
			return
		}

	// 🍔 Méthode 'POST' (publication d'un commentaire) :
	case "POST":
		fmt.Println("⭐ METHOD POST")
		// (1) Vérification du rôle et du statut de l'utilisateur :
		if user.Role < database.MEMBER || user.State > database.NORMAL {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			log.Println("❌ POST | Publication d'un commentaire refusée. L'utilisateur a un rôle ou statut inapproprié.")
			log.Println("Utilisateur : ", user)
			return
		}
		// (2) Récupération du commentaire :
		content := r.FormValue("comment")

		// (3) Vérification de la validité du commentaire :
		if toolbox.IsEmptyString(content) || len(content) < 3 {
			err := MyTemplates.ExecuteTemplate(w, "404", user)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
			// http.Error(w, "404 BAD REQUEST\nYour reply must be at least 3 characters.", http.StatusBadRequest)
			// log.Println("❌ POST | Publication d'un commentaire refusée. Le contenu du commentaire est invalide.")
			// log.Println("Contenu du commentaire : “", content, "”")
		}

		// (4) Récupération du GIF :
		gif := r.FormValue("gif")
		// (5) Vérification qu'il s'agit bien d'un gif tenor
		if !strings.Contains(gif, "https://media.tenor.com/images/") || gif == "Choose a gif" {
			gif = ""
		}
		var comment database.Comment

		comment.AuthorID = user.ID
		comment.PostID = ID
		comment.Content = content
		comment.Gif = gif
		comment.Date = time.Now()
		comment.State = database.PUBLISHED

		err = comment.InsertIntoDatabase()
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		fullPath := r.URL.Path + "#contact"

		// (6) Actualisation de la page (c-à-d redirection vers la page du post) :
		http.Redirect(w, r, fullPath, http.StatusSeeOther) // r.URL.Path = '/𝐩𝐨𝐬𝐭/{𝐈𝐃}'
	}

}
