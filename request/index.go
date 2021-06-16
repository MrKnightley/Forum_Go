package request

import (
	"fmt"
	"forum/database"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, user database.User) {
	if r.URL.Path != "/" {
		err := MyTemplates.ExecuteTemplate(w, "400", user)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
		// http.Error(w, "404 NOT FOUND", http.StatusNotFound)
	}

	// Remplissage d'une struct Data pour la page d'accueil :

	/*  DataForIndex struct {
	    - User  				User
	    - Categories 			[]Category
		- MostLikedPost 		Post
		- MostCommentedPost 	Post
		- MostRecentPost 		Post
		- PromotedPost			Post
	} */

	var err, err2, err3, err4 error
	var dataForIndex database.DataForIndex

	dataForIndex.User = user
	dataForIndex.Categories = database.GetCategoriesList()

	dataForIndex.MostLikedPost, err = database.GetMostLikedPostOfTheWeek()
	dataForIndex.MostLikedPost.Author, _ = database.GetUserByID(dataForIndex.MostLikedPost.AuthorID)

	dataForIndex.MostCommentedPost, err2 = database.GetMostCommentedPostOfTheWeek()
	dataForIndex.MostCommentedPost.Author, _ = database.GetUserByID(dataForIndex.MostCommentedPost.AuthorID)

	dataForIndex.MostRecentPost, err3 = database.GetMostRecentPost()
	dataForIndex.MostRecentPost.Author, _ = database.GetUserByID(dataForIndex.MostRecentPost.AuthorID)

	dataForIndex.PromotedPost, err4 = database.GetPromotedPost()
	dataForIndex.PromotedPost.Author, _ = database.GetUserByID(dataForIndex.PromotedPost.AuthorID)

	if err != nil || err2 != nil || err3 != nil || err4 != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("❌ ERREUR | Impossible de récupérer les 3 ou l'un des 3 posts pour la page Index")
		fmt.Println(err, err2, err3)
		return
	}

	log.Println("👩 INDEX | dataForIndex.User : ", dataForIndex.User)

	err5 := MyTemplates.ExecuteTemplate(w, "index", dataForIndex)
	if err5 != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
