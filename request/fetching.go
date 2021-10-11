package request

import (
	"encoding/json"
	"forum/admin"
	"forum/database"
	"forum/toolbox"
	"net/http"
)

var userCanDelete = []string{"posts", "comments", "gif", "images", "users", "tickets", "ticket_answers"}
var userCanAlterate = []string{"posts", "comments", "users", "tickets", "ticket_answers"}

//requete de modification de la bdd
func Fetching(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {
	//savoir quelle action est demander
	case "POST":
		var ERROR []byte
		var received database.ReceivedData
		err := json.NewDecoder(r.Body).Decode(&received)
		if err != nil {
			ERROR, _ = json.Marshal("ERROR WHILE DECODING JSON")
			w.Write(ERROR)
			panic(err)
		}
		switch received.Action {
		case "CREATE":

		case "UPDATE":
			if isLegal(received, user) {
				performAction(received)
				ok, _ := json.Marshal(received)
				w.Write(ok)
			} else {
				ERROR, _ = json.Marshal("DOESNT HAVE PERMISSION.")
				w.Write(ERROR)
			}
			break
		case "DELETE":
			if isLegal(received, user) {
				performAction(received)
				ok, _ := json.Marshal(received)
				w.Write(ok)
			} else {
				ERROR, _ = json.Marshal("DOESNT HAVE PERMISSION.")
				w.Write(ERROR)
			}
		}
	}
}

//Tout est vérifier je fait ce que la requête demande
func performAction(r database.ReceivedData) {
	var query string
	switch r.Action {
	case "UPDATE":
		query = "UPDATE " + r.Table + " SET " + r.What + " = \"" + r.NewValue + "\""
		break
	case "DELETE":
		if r.Is == "cell" { //si c'est une cellule c'est un update sur une valeur null
			query = "UPDATE " + r.Table + " SET " + r.What + " = \"\" WHERE id" + " = \"" + r.ID + "\""
		} else { //une table je la supprime
			query = "DELETE FROM " + r.Table + " WHERE id = \"" + r.ID + "\""
		}
		break
	case "CREATE":

	}
	_, err := database.Db.Exec(query)
	if err != nil {
		panic(err)
	}
}

//a les droits administrateur ou demande de toucher quelque chose qui lui appartient
func isLegal(received database.ReceivedData, user database.User) bool {
	var answer bool
	if user.Role > 2 || user.IsAuthor(received.ID, received.Table) {
		switch received.Action {
		case "UPDATE":
			answer = canUpdate(user.Role, received, user)
		case "DELETE":
			answer = canDelete(user.Role, received, user)
		case "CREATE":
		}
	} else {
		answer = false
	}
	return answer
}

//Si Delete est quelque chose que les droits de l'user permet de modifier ou qu'il est administrateur
func canDelete(role int, r database.ReceivedData, user database.User) bool {
	var answer bool
	if (admin.TableExist(r.Table)) && (user.Role > 2 || toolbox.Contain(userCanDelete, r.Table)) {
		answer = true
	} else {
		answer = false
	}

	return answer
}

//Si update est quelque chose que les droits de l'user permet de modifier ou qu'il est administrateur
func canUpdate(role int, r database.ReceivedData, user database.User) bool {
	var answer bool
	if (admin.ColExist(r.Table, r.What)) && (user.Role > 2 || toolbox.Contain(userCanAlterate, r.What)) {
		answer = true
	} else {
		answer = false
	}

	return answer
}
