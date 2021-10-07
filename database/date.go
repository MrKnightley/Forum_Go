package database

import (
	"database/sql"
	"strconv"
	"time"
)

//Remplie des ints correspondant a 4 périodes: a vie, du mois,de la semaine,des dernières 24heures. Nombre de posts au cours de ces dates.
func GetNumberOfPostByDateAndPostCategory(cat int, life *int, month *int, week *int, day *int) {
	monthly := time.Now().AddDate(0, -1, 0)
	weekly := time.Now().AddDate(0, 0, -7)
	daily := time.Now().AddDate(0, 0, -1)
	rows, err := Db.Query("SELECT date FROM posts p WHERE p.category_id = ?", cat)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		rows.Scan(&post.Date)
		*life++
		if post.Date.After(daily) {
			*month++
			*week++
			*day++
		} else if post.Date.After(weekly) {
			*month++
			*week++
		} else if post.Date.After(monthly) {
			*month++
		}
	}
}

//Remplie des ints correspondant a 4 périodes: a vie, du mois,de la semaine,des dernières 24heures. Nombre de commentaires au cours de ces dates.
func GetNumberOfCommentByDateAndPostCategory(cat int, life *int, month *int, week *int, day *int) {
	monthly := time.Now().AddDate(0, -1, 0)
	weekly := time.Now().AddDate(0, 0, -7)
	daily := time.Now().AddDate(0, 0, -1)
	rows, err := Db.Query("SELECT c.date FROM comments c INNER JOIN posts p ON p.id = c.post_id WHERE p.category_id = ?", cat)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		rows.Scan(&comment.Date)
		*life++
		if comment.Date.After(daily) {
			*month++
			*week++
			*day++
		} else if comment.Date.After(weekly) {
			*month++
			*week++
		} else if comment.Date.After(monthly) {
			*month++
		}
	}
}

//Remplie des ints correspondant a 4 périodes: a vie, du mois,de la semaine,des dernières 24heures. Nombre de likes OU dislike au cours de ces dates. Reaction correspond a like ou dislike.
func GetNumberOfReactionByDate(cat int, reaction string, life *int, month *int, week *int, day *int) {
	monthly := time.Now().AddDate(0, -1, 0)
	weekly := time.Now().AddDate(0, 0, -7)
	daily := time.Now().AddDate(0, 0, -1)
	rows, err := Db.Query(`SELECT cl.date FROM comment_likes cl INNER JOIN comments c ON c.id = cl.comment_id INNER JOIN posts p ON p.id = c.post_id and p.category_id = ? WHERE cl.type = ? 
						   UNION ALL 
						   SELECT pl.date FROM post_likes pl INNER JOIN posts p ON p.id and p.id = pl.post_id and p.category_id = ? WHERE pl.type = ?`, cat, reaction, cat, reaction)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		//sa prend aussi les post likes
		var comment CommentLike
		rows.Scan(&comment.Date)
		*life++
		if comment.Date.After(daily) {
			*month++
			*week++
			*day++
		} else if comment.Date.After(weekly) {
			*month++
			*week++
		} else if comment.Date.After(monthly) {
			*month++
		}
	}

}

//Renvoie le post le plus like (ne prend pas en compte les dislike) de la semaine et une erreur si une erreur parviens lors de l'appel de db
func GetMostLikedPostOfTheWeek() (Post, error) {
	var res Post
	var rows *sql.Rows
	var err error

	for day := 7; res.Title == ""; day = day + 7 {
		rows, err = Db.Query(`SELECT *,count(case post_id WHEN l.type = "like" then 1 else 0 end) AS amount FROM posts p 
							INNER JOIN post_likes l ON l.post_id = p.id
							WHERE p.date > datetime('now', '-` + strconv.Itoa(day) + ` day')
							GROUP BY l.post_id
							ORDER BY amount DESC
							LIMIT 1`)
		defer rows.Close()
		for rows.Next() {
			var postID int
			var userID int
			var myType string
			var date time.Time
			var amount int
			rows.Scan(&res.ID, &res.Title, &res.AuthorID, &res.Content, &res.CategoryID, &res.Date, &res.Image, &res.State, &res.Reason, &postID, &userID, &myType, &date, &amount)
		}
	}

	return res, err
}

//Renvoie le post le plus commenter de la semaine et une erreur si une erreur parviens lors de l'appel de db
func GetMostCommentedPostOfTheWeek() (Post, error) {
	var res Post
	var rows *sql.Rows
	var err error

	for day := 7; res.Title == ""; day = day + 7 {
		rows, err = Db.Query(`SELECT *,count(post_id) AS amount FROM posts p 
								INNER JOIN comments c ON c.post_id = p.id
								WHERE p.date > datetime('now', '-` + strconv.Itoa(day) + ` day')
								GROUP BY c.post_id
								ORDER BY amount DESC
								LIMIT 1`)
		defer rows.Close()
		for rows.Next() {
			var id int
			var authorID int
			var postID int
			var content string
			var gif string
			var date time.Time
			var state int
			var reason string
			var amount int
			rows.Scan(&res.ID, &res.Title, &res.AuthorID, &res.Content, &res.CategoryID, &res.Date, &res.Image, &res.State, &res.Reason, &id, &authorID, &postID, &content, &gif, &date, &state, &reason, &amount)
		}
	}
	return res, err
}

//Renvoie le post le plus récent, une erreur est fournis si l'appel de Db plante
func GetMostRecentPost() (Post, error) {
	var res Post
	rows, err := Db.Query(`SELECT * FROM posts p 
	ORDER BY id DESC
	LIMIT 1 `)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&res.ID, &res.Title, &res.AuthorID, &res.Content, &res.CategoryID, &res.Date, &res.Image, &res.State, &res.Reason)
	}
	return res, err
}

//Renvoie le post promus par un modérateur ou administrateur
func GetPromotedPost() (Post, error) {
	var res Post
	rows, err := Db.Query(`SELECT p.id,title,author_id,content,category_id,date,image,state,reason FROM promoted_post pp INNER JOIN posts p ON p.id = pp.post_id`)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&res.ID, &res.Title, &res.AuthorID, &res.Content, &res.CategoryID, &res.Date, &res.Image, &res.State, &res.Reason)
	}
	return res, err
}
