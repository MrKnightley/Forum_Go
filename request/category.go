package request

import (
	"encoding/json"
	"fmt"
	"forum/database"
	"forum/toolbox"
	"log"
	"net/http"
)

// HandleFunc pour la page catégorie :
func Category(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {
	case "GET":

		// (1) Récupération de l'ID de la catégorie depuis l'URL :
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

		// (2) Remplissage d'une struct Data pour chaque page de catégorie :

		/*  DataForCategory struct {
			- ID    int
			- Name  string
			- User  User
			- Posts []Post
		} */

		var dataForCategory database.DataForCategory

		myCategory, err := database.GetCategoryByID(ID) // Récupération du nom de la catégorie
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Println("❌ ERREUR | Impossible de récupérer la catégorie depuis l'ID : ", ID)
			return
		}

		dataForCategory.ID = ID // ID de la catégorie

		dataForCategory.User = user // Utilisateur actuel (dont la session est en cours)

		dataForCategory.Name = myCategory.Name // Nom de la catégorie
		if dataForCategory.Name == "" {
			err := MyTemplates.ExecuteTemplate(w, "404", user)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			// http.Error(w, "404 NOT FOUND", http.StatusNotFound)
			log.Println("❌ ERREUR | Le nom de la catégorie n°", ID, " est une string vide.")
			return
		}
		dataForCategory.Posts, err = database.GetPostsByCategoryID(ID) // Récupération de tous les posts appartenant à la catégorie
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		// (3) Exécution du template 'category' avec la Data :
		err = MyTemplates.ExecuteTemplate(w, "category", dataForCategory)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	case "POST":
		var p receivedData
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			panic(err)
		} else {
			_, err = database.Db.Exec("INSERT INTO categories (name) VALUES ($1)", p.NewValue)
			if err != nil {
				ERROR, _ := json.Marshal("ERROR")
				w.Write(ERROR)
				panic(err)
			}
			database.Db.QueryRow("SELECT id FROM categories WHERE name = ?", p.NewValue).Scan(&p.ID)
			fmt.Println(p.ID)
			var msg = `{"new": "` + p.NewValue + `" , "id" : "` + p.ID + `" }`
			w.Write([]byte(msg))
		}
	}
}
