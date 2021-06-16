package request

import (
	"forum/database"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Reaction(w http.ResponseWriter, r *http.Request, user database.User) {
	// L'URL sera de la forme /𝐫𝐞𝐚𝐜𝐭𝐢𝐨𝐧/{𝐭𝐲𝐩𝐞}/{𝐫𝐞𝐜𝐞𝐢𝐯𝐞𝐫}/{𝐈𝐃}
	// Par exemple : '/𝐫𝐞𝐚𝐜𝐭𝐢𝐨𝐧/𝐝𝐢𝐬𝐥𝐢𝐤𝐞/𝐜𝐨𝐦𝐦𝐞𝐧𝐭/𝟐𝟕'

	pathArray := strings.Split(r.URL.Path[1:], "/") // Par exemple, pathArray = ["reaction", "like", "post", "13"]
	if len(pathArray) != 4 {
		err := MyTemplates.ExecuteTemplate(w, "404", user)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		// http.Error(w, "404 NOT FOUND", http.StatusNotFound)
		return
	}

	reactType := pathArray[1]             // 'like' ou 'dislike'
	receiverType := pathArray[2]          // 'post' ou 'comment'
	ID, err := strconv.Atoi(pathArray[3]) // ID du post ou commentaire auquel réagir

	if err != nil || ID < 1 || (reactType != "like" && reactType != "dislike") || (receiverType != "post" && receiverType != "comment") {
		err := MyTemplates.ExecuteTemplate(w, "400", user)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
		// http.Error(w, "400 Bad Request", http.StatusBadRequest)
		// return
	}

	switch receiverType {
	case "post":
		post, _ := database.GetPostByID(ID, user.ID)

		var like database.PostLike
		like.PostID = ID
		like.UserID = user.ID
		like.Date = time.Now()

		switch reactType {
		case "like":
			if post.Disliked { // Si le user avait DISLIKÉ le post auparavant...
				like.Type = "dislike"
				err = like.DeleteFromDatabase() // ... on supprime le DISLIKE d'abord
				like.Type = "like"
				err = like.InsertIntoDatabase() // ... puis on ajoute le LIKE
			} else {
				if post.Liked { // Si le user avait déjà LIKÉ le post auparavant...
					like.Type = "like"
					err = like.DeleteFromDatabase() // ... on supprime le LIKE pour l'annuler

				} else {
					like.Type = "like"
					err = like.InsertIntoDatabase() // ... sinon, on ajoute simplement le LIKE à la base de données
				}
			}
		case "dislike":
			if post.Liked { // Si le user avait LIKÉ le post auparavant...
				like.Type = "like"
				err = like.DeleteFromDatabase() // ... on supprime le LIKE
				like.Type = "dislike"
				err = like.InsertIntoDatabase() // ... puis on ajoute le DISLIKE...

			} else {
				if post.Disliked { // Si le user avait déjà DISLIKÉ le post auparavant...
					like.Type = "dislike"
					err = like.DeleteFromDatabase() // ... on supprime le DISLIKE pour l'annuler

				} else {
					like.Type = "dislike"
					err = like.InsertIntoDatabase() // ... sinon, on ajoute simplement le DISLIKE à la base de données
				}
			}
		}
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Redirection vers la page du post :
		postURL := "/post/" + strconv.Itoa(post.ID)
		http.Redirect(w, r, postURL, http.StatusSeeOther)

	case "comment":
		comment, _ := database.GetCommentByID(ID, user.ID)

		var like database.CommentLike
		like.CommentID = ID
		like.UserID = user.ID
		like.Date = time.Now()

		switch reactType {
		case "like":
			if comment.Disliked { // Si le user avait DISLIKÉ le commentaire auparavant...
				like.Type = "dislike"
				err = like.DeleteFromDatabase() // ... on supprime le DISLIKE d'abord
				like.Type = "like"
				err = like.InsertIntoDatabase() // ... puis on ajoute le LIKE

			} else {
				if comment.Liked { // Si le user avait déjà LIKÉ le commentaire auparavant...
					like.Type = "like"
					err = like.DeleteFromDatabase() // ... on supprime le LIKE pour l'annuler

				} else {
					like.Type = "like"
					err = like.InsertIntoDatabase() // ... sinon, on ajoute simplement le LIKE à la base de données
				}
			}
		case "dislike":
			if comment.Liked { // Si le user avait LIKÉ le commentaire auparavant...
				like.Type = "like"
				err = like.DeleteFromDatabase() // ... on supprime le LIKE
				like.Type = "dislike"
				err = like.InsertIntoDatabase() // ... puis on ajoute le DISLIKE...

			} else {
				if comment.Disliked { // Si le user avait déjà DISLIKÉ le commentaire auparavant...
					like.Type = "dislike"
					err = like.DeleteFromDatabase() // ... on supprime le DISLIKE pour l'annuler

				} else {
					like.Type = "dislike"
					err = like.InsertIntoDatabase() // ... sinon, on ajoute simplement le DISLIKE à la base de données
				}
			}
		}
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Redirection vers la page du post :
		postURL := "/post/" + strconv.Itoa(comment.PostID)
		http.Redirect(w, r, postURL, http.StatusSeeOther)
	}
}
