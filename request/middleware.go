package request

import (
	"forum/database"
	"html/template"
	"log"
	"net/http"
)

type CustomFunc func(http.ResponseWriter, *http.Request, database.User)

var MyTemplates = template.Must(template.ParseGlob("./templates/*"))

// Middleware vérifiant le que l'utilisateur a l'autorisation (role) d'accéder à chacune des pages demandées :
func Auth(nextFunction CustomFunc, credentials string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := database.GetUserByCookie(w, r)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		if credentials == "guests only" && user.Role != database.GUEST {
			log.Println("⚠️ AUTH | Access denied. Guests only but user's role is : ", user.Role)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if credentials == "members only" && user.Role < database.MEMBER {
			log.Println("⚠️ AUTH | Access denied. Members only but user's role is : ", user.Role)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if credentials == "active members only" && (user.Role < database.MEMBER || user.State != database.NORMAL) {
			log.Println("⚠️ AUTH | Access denied. Active members only but user's role is ", user.Role, " and user's state is ", user.State)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if credentials == "active moderators only" && (user.Role < database.MODERATOR || user.State != database.NORMAL) {
			log.Println("⚠️ AUTH | Access denied. Active moderators only but user's role is ", user.Role, " and user's state is ", user.State)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if credentials == "active admins only" && (user.Role < database.ADMIN || user.State != database.NORMAL) {
			log.Println("⚠️ AUTH | Access denied. Active administrators only but user's role is ", user.Role, " and user's state is ", user.State)
			err := MyTemplates.ExecuteTemplate(w, "unauthorized", user)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
		}

		if credentials == "unaffiliated members only" && (user.House.ID != database.UNAFFILIATED || user.Role < database.MEMBER || user.State != database.NORMAL) {
			log.Println("⚠️ AUTH | Access denied. Unaffiliated members only but user's house ID is ", user.House.ID, " and user's role is ", user.Role)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		nextFunction(w, r, user)
	})
}
