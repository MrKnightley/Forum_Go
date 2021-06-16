package request

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var profileTmpl = template.Must(template.ParseGlob("./templates/*"))

type receivedData struct {
	ID       string `json:"id"`
	Category string `json:"cat"`
	Value    string `json:"val"`
	NewValue string `json:"newVal"`
	Table    string `json:"table"`
	Reason   string `json:"reason"`
}

// HandleFunc pour la page profile de l'utilisateur :
func ProfilePage(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {
	// 🍔 Méthode 'GET' — Lorsqu'on arrive sur la page register.html pour la 1ère fois :
	case "GET":
		var err1, err2, err3, err4, err5 error
		var data database.DataForProfile

		/* type DataForProfile struct {
			User          User
			Profile       User
			Posts         []Post
			Comments      []Comment
			LikedPosts    []Post
			LikedComments []Comment
		} */

		profile := r.URL.Query().Get("user")

		data.User = user
		data.Profile, err1 = database.GetUserByUsernameOrEmail(strings.ToLower(profile))
		data.Profile.Badges = database.GetBadgeByUserID(data.Profile.ID)
		data.Posts, err2 = database.GetPostsFromUserByID(data.Profile.ID)
		data.Comments, err3 = database.GetCommentFromUserByID(data.Profile.ID)
		data.LikedPosts, err4 = database.GetPostsLikedByUser(data.Profile.ID)
		data.LikedComments, err5 = database.GetCommentsLikedByUser(data.Profile.ID)

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Println("❌ ERREUR | Impossible de récupérer l'utilisateur pour la page Profile")
			fmt.Println(err1, err2, err3)
			return
		}

		if data.Profile.ID == 0 {
			err := MyTemplates.ExecuteTemplate(w, "404", user)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			// http.Error(w, "404 NOT FOUND", http.StatusNotFound)
			log.Println("❌ ERREUR | Impossible de récupérer l'utilisateur n°", data.Profile.ID, " pour la page Profile")
			return
		}

		err := profileTmpl.ExecuteTemplate(w, "profile", data)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

	case "POST":
		var p receivedData
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.Write([]byte(`{"message":"ERROR"}`))
			panic(err)
		} else {
			deleteAccount(w, r, user, p)
		}
	}
}

func deleteAccount(w http.ResponseWriter, r *http.Request, user database.User, p receivedData) {
	_, err := database.Db.Exec("UPDATE users SET state = 2 WHERE id = ?", p.ID)
	if err != nil {
		ERROR, _ := json.Marshal("ERROR WHILE DELETE")
		w.Write(ERROR)
		panic(err)
	}
	database.Db.Exec("DELETE FROM sessions WHERE user_id = $1", p.ID)
	// On récupère le cookie dont le nom est "session", et on modifie son MaxAge (nombre négatif) pour le faire expirer :
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1, // Fait expirer le cookie immédiatement
	}
	http.SetCookie(w, cookie) // Suppression du cookie
	w.Write([]byte(`{"message":"deleted"}`))
	http.Redirect(w, r, "/", http.StatusFound)
}
