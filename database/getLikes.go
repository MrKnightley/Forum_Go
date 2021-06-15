package database

import "log"

// Fonction récupératrice de likes/dislikes d'un post et du booléen 'liked/disliked par le user' en fonction de l'ID du post (et de l'ID du user) :
func GetLikesByPostID(ID int, currentUserID int) ([]PostLike, []PostLike, bool, bool) {
	var Likes, Dislikes []PostLike
	var Liked, Disliked bool

	rows, err := Db.Query("SELECT * FROM post_likes WHERE post_id = ?", ID)
	defer rows.Close()
	if err != nil {
		log.Println("❌ ERREUR | Impossible de sélectionner les colonnes de la table post_likes avec post_id = ", ID)
		panic(err)
	}

	for rows.Next() {
		var postLike PostLike
		rows.Scan(&postLike.PostID, &postLike.UserID, &postLike.Type, &postLike.Date)

		switch postLike.Type {
		case "like":
			Likes = append(Likes, postLike)
			if postLike.UserID == currentUserID {
				Liked = true
			}

		case "dislike":
			Dislikes = append(Dislikes, postLike)
			if postLike.UserID == currentUserID {
				Disliked = true
			}
		}
	}

	return Likes, Dislikes, Liked, Disliked // Nombre de likes/dislikes, et booléen 'liké/disliké par l'utilisateur connecté'
}

// Fonction récupératrice de likes/dislikes d'un commentaire et du booléen 'liked/disliked par le user' en fonction de l'ID du commentaire (et de l'ID du user) :
func GetLikesByCommentID(ID int, currentUserID int) ([]CommentLike, []CommentLike, bool, bool) {
	var Likes, Dislikes []CommentLike
	var Liked, Disliked bool

	rows, err := Db.Query("SELECT * FROM comment_likes WHERE comment_id = ?", ID)
	defer rows.Close()
	if err != nil {
		log.Println("❌ ERREUR | Impossible de sélectionner les colonnes de la table comment_likes avec comment_id = ", ID)
		panic(err)
	}

	for rows.Next() {
		var commentLike CommentLike
		rows.Scan(&commentLike.CommentID, &commentLike.UserID, &commentLike.Type, &commentLike.Date)
		switch commentLike.Type {
		case "like":
			Likes = append(Likes, commentLike)
			if commentLike.UserID == currentUserID {
				Liked = true
				Disliked = false
			}

		case "dislike":
			Dislikes = append(Dislikes, commentLike)
			if commentLike.UserID == currentUserID {
				Disliked = true
				Liked = false
			}
		}
	}

	return Likes, Dislikes, Liked, Disliked
}
