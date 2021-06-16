package request

import (
	"forum/database"
	"forum/toolbox"
	"log"
	"net/http"
	"strings"
)

//Gère les comptes des utilisateurs, Get affiche les info, Post reçois celle a modifier.
func Account(w http.ResponseWriter, r *http.Request, user database.User) {

	/* type DataForSettings struct {
		User  User
		Error ErrorData
	} */

	var dataForSettings database.DataForSettings
	dataForSettings.User = user

	switch r.Method {

	// 🍔 Méthode 'GET' — Lorsqu'on arrive sur la page register.html pour la 1ère fois :
	case "GET":
		err := profileTmpl.ExecuteTemplate(w, "account", dataForSettings)
		if err != nil {
			log.Println("❌ ERREUR | Impossible d'afficher le template Account")
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

	// 🍔 Méthode 'POST' — Lorsqu'on clique sur 'Apply changes' pour mettre à jour son compte :
	case "POST":

		// (1) Je récupère l'email, le nom d'utilisateur, le mot de passe et l'avatar mis à jour :
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		imagePath, err := toolbox.UploadImage(r, user.ID, "avatar")
		if err != nil && err.Error() != "http: no such file" {
			log.Println("❌ EDIT AVATAR | Impossible de récupérer le path de l'image uploadée.")
			err := MyTemplates.ExecuteTemplate(w, "400", user)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
			// http.Error(w, "400 Bad Request\n"+err.Error(), http.StatusBadRequest)
		}

		// (2) J'ajoute ces valeurs dans une struct userUpdated :
		userUpdated := user
		userUpdated.Username = strings.ToLower(username)
		userUpdated.Password = password
		userUpdated.Email = strings.ToLower(email)

		// (3) Je vérifie si l'email ou username existe déjà dans la base de données :
		dataForSettings.Error = CheckNewAccount(userUpdated)
		if dataForSettings.Error.Account != nil || dataForSettings.Error.Username != nil || dataForSettings.Error.Email != nil {
			log.Println("❌ EDIT ACCOUNT SETTINGS | Request denied : ", dataForSettings.Error)
			MyTemplates.ExecuteTemplate(w, "account", dataForSettings) // On ré-exécute le template 'Account' en affichant cette fois une div "Username déjà pris", etc.
			return
		}

		// (4) Je mets à jour la base de données avec les nouvelles valeurs, SI et seulement si elles sont valides :
		if len(username) >= 3 {
			user.Username = userUpdated.Username
			user.UpdateInDatabase("username")
		}

		if len(email) >= 5 {
			user.Email = userUpdated.Email
			user.UpdateInDatabase("email")
		}

		if len(password) >= 6 {
			user.Password = userUpdated.Password
			user.UpdateInDatabase("password")
		}

		if len(imagePath) > 0 {
			user.Avatar = imagePath
			user.UpdateInDatabase("avatar")
		}

		err2 := profileTmpl.ExecuteTemplate(w, "account-success", user)
		if err2 != nil {
			log.Println("❌ ERREUR | Impossible d'afficher le template Account-Success")
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
