package admin

import (
	"encoding/json"
	"forum/database"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
	// "golang.org/x/crypto/bcrypt"
)

var profileTmpl = template.Must(template.ParseGlob("./templates/*"))

type Data struct {
	User     []database.User
	Post     []database.Post
	Comment  []database.Comment
	Category []database.Category
	Self     database.User
}

type ReceivedData struct {
	ID       string `json:"id"`
	Category string `json:"cat"`
	Value    string `json:"val"`
	NewValue string `json:"newVal"`
	Table    string `json:"table"`
	Reason   string `json:"reason"`
}

var results []string

var tmpl *template.Template

func Moderation(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {
	case "POST":
		var p ReceivedData
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			ERROR, _ := json.Marshal("ERROR")
			w.Write(ERROR)
			panic(err)
		}
		//Si la demande est de promote un post alors on le promote ici
		if p.Table == "promote" {
			database.Db.Exec("insert or replace into promoted_post (post_id) values (?)", p.ID)
		} else {
			//Sinon c'est pour update un élément de la base de donnée.
			if isDatabaseTable(p.Table) && ColExist(p.Table, p.Category) {
				querry := "UPDATE " + p.Table + " SET " + p.Category
				_, err = database.Db.Exec(querry+` = ? WHERE id = ?`, p.NewValue, p.ID)
				if err != nil {
					panic(err)
				}
			}
		}
		var msg = "{\"message\": \"" + p.NewValue + "\"}"
		w.Write([]byte(msg))
	case "GET":
		//preparation de la bdd pour afficher les utilisateurs
		var Data Data
		Data = GetClientList(Data)
		Data = GetCommentList(Data)
		Data = GetPostList(Data)
		Data.Self = user
		Data.Category = database.GetCategoriesList()
		err := profileTmpl.ExecuteTemplate(w, "moderation", Data)
		if err != nil {
			panic(err)
		}
	}
}
func TableExist(table string) bool {
	query := "SELECT * FROM " + table
	cols, err := database.Db.Query(query)
	rows, err := cols.Columns()
	if err != nil {
		panic(err)
	}
	return len(rows) > 0
}
func ColExist(table string, cat string) bool {
	query := "SELECT * FROM " + table
	cols, err := database.Db.Query(query)
	rows, err := cols.Columns()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(rows); i++ {
		if rows[i] == cat {
			return true
		}
	}
	return false
}
func isDatabaseTable(table string) bool {
	var newtable string
	rows, err := database.Db.Query("SELECT name FROM sqlite_master WHERE type ='table' AND name NOT LIKE 'sqlite_%';")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&newtable)
		if newtable == table {
			return true
		}
	}
	return false
}

//Récupère tout les users
func GetClientList(Data Data) Data {
	rows, _ := database.Db.Query("SELECT * FROM users")
	defer rows.Close()
	for rows.Next() {
		var newUser database.User
		err := rows.Scan(&newUser.ID, &newUser.Username, &newUser.Password, &newUser.Email, &newUser.Role, &newUser.Avatar, &newUser.Date, &newUser.State, &newUser.SecretQuestion, &newUser.SecretAnswer, &newUser.House.ID)
		if err != nil {
			panic(err)
		}
		Data.User = append(Data.User, newUser)
	}
	return Data
}

//Récupère tout les commentaires
func GetCommentList(Data Data) Data {
	rows, _ := database.Db.Query("SELECT * FROM comments")
	defer rows.Close()
	for rows.Next() {
		var newComment database.Comment
		err := rows.Scan(&newComment.ID, &newComment.AuthorID, &newComment.PostID, &newComment.Content, &newComment.Gif, &newComment.Date, &newComment.State, &newComment.Reason)
		if err != nil {
			panic(err)
		}
		Data.Comment = append(Data.Comment, newComment)
	}
	return Data
}

//Récupère tout les posts
func GetPostList(Data Data) Data {
	rows, _ := database.Db.Query("SELECT * FROM posts")
	defer rows.Close()
	for rows.Next() {
		var newPost database.Post
		err := rows.Scan(&newPost.ID, &newPost.Title, &newPost.AuthorID, &newPost.Content, &newPost.CategoryID, &newPost.Date, &newPost.Image, &newPost.State, &newPost.Reason)
		if err != nil {
			panic(err)
		}
		Data.Post = append(Data.Post, newPost)
	}
	return Data
}
