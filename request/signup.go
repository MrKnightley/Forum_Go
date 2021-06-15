package request

import (
	"database/sql"
	"forum/database"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Username       string
	SecretQuestion string
	ID             int
	Problem        string
}

var tmpl *template.Template

func Signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		//Je récupère les arguments donner par la méthode post
		r.ParseForm()
		//Partie 1 du formulaire
		password := r.Form.Get("password")
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		date := time.Now().Unix()
		//Partie 2 du formulaire
		secretQuestion := r.Form.Get("question")
		secretAnswer := r.Form.Get("answer")
		//
		var Data Data
		var err error
		var id = r.Form.Get("id")
		//open database
		database.Db, err = sql.Open("sqlite3", "./database/database.db")
		defer database.Db.Close()
		if err != nil {
			panic(err)
		}
		//
		if id == "" {
			verif := validAccount(database.Db, email, username)
			if verif == "" {
				//preparation database
				statement, err := database.Db.Prepare("INSERT INTO users (username,password,email,role,date,state) VALUES (?,?,?,?,?,?)")
				if err != nil {
					log.Fatal(err)
				}
				//
				cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
				_, err = statement.Exec(username, cryptedPassword, email, 1, date, 0, 0, 0)
				if err != nil {
					panic(err)
				}
				rows, err := database.Db.Query("SELECT id,password,username,email,date FROM users")
				defer rows.Close()
				if err != nil {
					panic(err)
				}
				for rows.Next() {
					rows.Scan(&Data.ID, &password, &username, &email, &date)
				}
				tmpl, _ = template.ParseFiles("./templates/signup.html")
				tmpl.ExecuteTemplate(w, "signup", Data)
			} else if verif == "mail" {
				Data.Problem = "alreadyUsedEmail"
				tmpl, _ = template.ParseFiles("./templates/signup.html")
				tmpl.ExecuteTemplate(w, "signup", Data)
			} else {
				Data.Problem = "alreadyUsedUsername"
				tmpl, _ = template.ParseFiles("./templates/signup.html")
				tmpl.ExecuteTemplate(w, "signup", Data)
			}
		} else {
			protectedAnswer, _ := bcrypt.GenerateFromPassword([]byte(secretAnswer), 14)
			_, err = database.Db.Exec("UPDATE users SET secretQuestion = ?,secretAnswer = ? WHERE id=?", secretQuestion, protectedAnswer, id)
			if err != nil {
				panic(err)
			}
		}

	case "GET":
		tmpl, _ = template.ParseFiles("./templates/signup.html")
		tmpl.ExecuteTemplate(w, "signup", nil)
	}

}

func validAccount(database *sql.DB, mail string, username string) string {
	var email string
	var username2 string
	var err error
	rows, err := database.Query("SELECT email FROM users")
	defer rows.Close()
	if err != nil {
		return ""
	}
	for rows.Next() {
		rows.Scan(&email)
		if email == mail {
			return "mail"
		}
	}
	rows, _ = database.Query("SELECT username FROM users")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&username2)
		if username2 == username {
			return "username"
		}
	}
	return ""
}
