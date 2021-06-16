package request

import (
	"forum/database"
	"forum/toolbox"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {

	// 🍔 Méthode 'GET' — Lorsqu'on arrive sur la page login.html pour la 1ère fois :
	case "GET":
		err := MyTemplates.ExecuteTemplate(w, "login", nil)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

	// 🍔 Méthode 'POST' — Lorsqu'on sur le bouton 'Login' pour se connecter :
	case "POST":
		// Je récupère le username ou email et le mot de passe :
		identifier := strings.ToLower(r.FormValue("identifier")) // Username or Email
		password := r.FormValue("password")

		// Je vérifie que les identifiants saisis ne sont pas vides :
		if toolbox.IsEmptyString(identifier) || toolbox.IsEmptyString(password) {
			err := MyTemplates.ExecuteTemplate(w, "400", user)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
			// http.Error(w, "400 BAD REQUEST\nTHE TEXT YOU ENTERED IS EMPTY.", http.StatusBadRequest)
			// return
		}

		// Je récupère l'utilisateur grâce à ses identifiants (s'il n'existe pas dans la base de données, user.ID == 0) :
		user, err := database.GetUserByUsernameOrEmail(identifier)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Vérification du statut de l'utilisateur (DELETED, NORMAL, BANNED) :
		if user.State != database.NORMAL {
			log.Println("❌ LOGIN | Accès refusé : compte banni ou désactivé.")
			http.Error(w, "Compte banni ou désactivé.", http.StatusBadRequest)
			return
		}

		// (1) Comparaison entre user.Password et 'password' (access denied if err != nil) :
		// (2) Si user.ID == 0 (valeur par défaut), l'utilisateur n'existe pas :
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil || user.ID == 0 {
			MyTemplates.ExecuteTemplate(w, "login", "No matching account was found.") // On ré-exécute le template avec un div 'No matching account was found'.
			return
		}

		// Création et ajout d'une session (cookie) dans la base de données :
		AddBadgeIfUnlocked(user)
		err = database.AddSessionToDatabase(w, r, user)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Après s'être identifié, on est redirigé vers la page index :
		http.Redirect(w, r, "/", http.StatusFound)
		log.Println("✔️ LOGIN | Access granted.")
		log.Println("Successfully logged in: ", user)
	}
}

func Logout(w http.ResponseWriter, r *http.Request, user database.User) {
	// Suppression de la session de l'utilisateur dans la base de données :
	database.Db.Exec("DELETE FROM sessions WHERE user_id = $1", user.ID)

	// On récupère le cookie dont le nom est "session", et on modifie son MaxAge (nombre négatif) pour le faire expirer :
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1, // Fait expirer le cookie immédiatement
	}
	http.SetCookie(w, cookie) // Suppression du cookie

	// On est redirigé vers la page index :
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
