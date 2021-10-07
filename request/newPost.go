package request

import (
	"forum/database"
	"forum/toolbox"
	"log"
	"net/http"
	"strconv"
	"time"
)

func NewPost(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {

	// 🍔 Méthode 'GET' — Lorsqu'on arrive sur la page :
	case "GET":
		// Remplissage d'une struct Data pour chaque page 'New Post' :

		/*  DataForNewPost struct {
			- User       User
			- Categories []Category
		} */

		var dataForNewPost database.DataForNewPost

		dataForNewPost.User = user
		dataForNewPost.Categories = database.GetCategoriesList()

		err := MyTemplates.ExecuteTemplate(w, "newpost", dataForNewPost)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Println("❌ ERREUR | Impossible d'exécuter le template “newpost”.")
			return
		}

	// 🍔 Méthode 'POST' — Lorsqu'on clique sur 'Publier' :
	case "POST":
		// (1) Récupération du titre, du contenu, et de la catégorie du post à publier :
		title := r.FormValue("title")
		title = toolbox.FormatString(title) // Formattage du titre

		content := r.FormValue("content")
		content = toolbox.FormatString(content) // Formattage du contenu

		categoryID, err := strconv.Atoi(r.FormValue("category"))
		if err != nil {
			log.Println("❌ POST | Impossible de récupérer l'ID de la catégorie du post à publier.")
			err := MyTemplates.ExecuteTemplate(w, "400", user)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
			// http.Error(w, "400 Bad Request", http.StatusBadRequest)
			// return
		}

		if toolbox.IsEmptyString(title) || toolbox.IsEmptyString(content) {
			log.Println("❌ POST | Impossible de publier le post : le titre ou le contenu est vide.")
			err := MyTemplates.ExecuteTemplate(w, "400", user)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
			// http.Error(w, "400 Bad Request\nThe text you added is empty.", http.StatusBadRequest)
			// return
		}

		// (2) Récupération de l'image uploadée par l'utilisateur pour son post :
		imagePath, err := toolbox.UploadImage(r, user.ID, "post")
		if err != nil && err.Error() != "http: no such file" {
			log.Println("❌ POST | Impossible de récupérer le path de l'image uploadée.")
			err := MyTemplates.ExecuteTemplate(w, "400", user)
			if err != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			}
			return
			// http.Error(w, "400 Bad Request\n"+err.Error(), http.StatusBadRequest)
			// return
		}

		// (3) Remplissage d'une struct 'Post' pour le post à publier :
		var post database.Post

		post.Title = title
		post.AuthorID = user.ID
		post.Content = content
		post.Date = time.Now()
		post.Image = imagePath
		post.CategoryID = categoryID
		post.State = database.PUBLISHED

		// (4) Insertion du post dans la base de données :
		postID, err := post.InsertIntoDatabase() // La méthode d'insertion dans la DB renvoie l'ID du post qui vient d'être inséré
		if err != nil || postID < 1 {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		// (5) Redirection vers la page du post :
		postURL := "/post/" + strconv.Itoa(postID)
		http.Redirect(w, r, postURL, http.StatusSeeOther)
	}
}
