package request

import (
	"database/sql"
	"forum/database"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func MotDePasseOublier(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		//chope les valeurs
		r.ParseForm()
		var Data Data
		var err error
		Data.SecretQuestion = r.Form.Get("answer")
		Data.Username = r.Form.Get("username")
		var temp string
		//ouvre la database
		database.Db, _ = sql.Open("sqlite3", "./database/database.db")
		defer database.Db.Close()
		if Data.Username != "" {
			//cherche la question de cet user
			rows, _ := database.Db.Query("SELECT secretQuestion FROM users WHERE username=?", Data.Username)
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&temp)
			}
		}
		if Data.SecretQuestion != "" {
			answer := r.Form.Get("answer")
			psw, _ := bcrypt.GenerateFromPassword([]byte(r.Form.Get("psw")), 14)
			var realAnswer string
			rows, _ := database.Db.Query("SELECT secretAnswer FROM users WHERE username=?", Data.Username)
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&realAnswer)
				if bcrypt.CompareHashAndPassword([]byte(realAnswer), []byte(answer)) == nil {
					database.Db.Exec("UPDATE users SET password = ? WHERE username=?", psw, Data.Username)
				}
			}
		} else {
			Data.SecretQuestion = temp
		}
		//
		tmpl, err = template.ParseFiles("./templates/MotDePasseOublier.html")
		if err != nil {
			panic(err)
		}
		err = tmpl.ExecuteTemplate(w, "motdepasseoublier", Data)
		if err != nil {
			panic(err)
		}
	case "GET":
		tmpl, _ = template.ParseFiles("./templates/MotDePasseOublier.html")
		tmpl.ExecuteTemplate(w, "motdepasseoublier", nil)
	}
}
