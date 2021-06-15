package database

import (
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// ⭐ Méthode pour insérer un utilisateur dans la base de données :
func (user *User) InsertIntoDatabase() error {

	addStatement, err := Db.Prepare("INSERT INTO users (username, password, email, role, avatar, date, state, secretQuestion, secretAnswer) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);")

	defer addStatement.Close()

	// Gestion de l'erreur :
	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible d'insérer l'utilisateur dans la base de données.")
		log.Println("Hypothèse : Mauvaise syntaxe du statement SQLite suivant :")
		fmt.Println("—————————————————————————")
		fmt.Println("INSERT INTO users")
		fmt.Println("	(username, password, email, role, avatar, date, state, secretQuestion, secretAnswer)")
		fmt.Println("VALUES")
		fmt.Println("	", user.Username, user.Password, user.Email, user.Role, user.Avatar, user.Date, user.State, user.SecretQuestion, user.SecretAnswer)
		fmt.Println("—————————————————————————")
		return err
	}

	cryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(cryptedPassword)
	user.Role = MEMBER
	user.Avatar = "/images/avatars/defaultAvatar.jpg"
	user.Date = time.Now()
	user.State = NORMAL
	user.SecretQuestion = ""
	user.SecretAnswer = ""

	addStatement.Exec(user.Username, user.Password, user.Email, user.Role, user.Avatar, user.Date, user.State, user.SecretQuestion, user.SecretAnswer)

	// ⚠️ IMPORTANT : Envoyer l'ID automatiquement attribué par la base de données à l'intérieur de user.ID :
	row := Db.QueryRow("SELECT id FROM users WHERE username = ? OR email = ?", user.Username, user.Email)
	row.Scan(&user.ID)
	log.Println("✔️ DATABASE | Inserted user “", user.Username, "” into database successfully.")
	log.Println("User's complete information : ", user)

	return nil
}

// ⭐ Méthode pour modifier un champ de l'utilisateur dans la base de données :
func (user *User) UpdateInDatabase(column string) error {

	statement := "UPDATE users SET " + column + " = ? WHERE id = ?"
	updateStatement, err := Db.Prepare(statement)

	defer updateStatement.Close()

	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible de mettre à jour la base de données.")
		log.Println("Hypothèse : Mauvaise syntaxe du statement SQLite “UPDATE users SET ", column, " WHERE id = ", user.ID, "”")
		return err
	}

	switch column {
	case "username":
		updateStatement.Exec(user.Username, user.ID)
	case "password":
		cryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		user.Password = string(cryptedPassword)
		updateStatement.Exec(user.Password, user.ID)
	case "email":
		updateStatement.Exec(user.Email, user.ID)
	case "role":
		updateStatement.Exec(user.Role, user.ID)
	case "avatar":
		updateStatement.Exec(user.Avatar, user.ID)
	case "state":
		updateStatement.Exec(user.State, user.ID)
	case "secretAnswer":
		updateStatement.Exec(user.SecretAnswer, user.ID)
	case "secretQuestion":
		updateStatement.Exec(user.SecretQuestion, user.ID)
	case "house_id":
		updateStatement.Exec(user.House.ID, user.ID)
	default:
		log.Println("❌ DATABASE | ERREUR : Impossible de mettre à jour la base de données.")
		log.Println("La colonne “", column, "” n'existe pas dans la table “USERS”.")
		return errors.New("ERREUR | La colonne à mettre à jour n'existe pas dans la table “USERS”.")
	}
	return nil
}

// ⭐ Méthode pour insérer un commentaire dans la base de données :
// Renvoie l'ID du post venant d'être inséré :
func (comment *Comment) InsertIntoDatabase() error {
	addStatement, err := Db.Prepare("INSERT INTO comments (author_id, post_id, content, gif, date, state) VALUES (?, ?, ?, ?, ?,?)")

	defer addStatement.Close()

	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible d'insérer le commentaire dans la base de données.")
		log.Println("Hypothèse : Mauvaise syntaxe du statement SQLite suivant :")
		log.Println("“INSERT INTO comments (author_id, post_id, content, date) VALUES (", comment.AuthorID, comment.PostID, comment.Content, comment.Gif, comment.Date, comment.State, ")”")
		return err
	}

	addStatement.Exec(comment.AuthorID, comment.PostID, comment.Content, comment.Gif, comment.Date, comment.State)
	return nil
}

// ⭐ Méthode pour insérer un post dans la base de données :
func (ticket *Ticket) InsertIntoDatabase() (int, error) {
	addStatement, err := Db.Prepare("INSERT INTO tickets (title, author_id,actual_admin, content, date, state) VALUES (?,?, ?, ?, ?, ?)")

	defer addStatement.Close()

	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible d'insérer le post dans la base de données.")
		log.Println("Hypothèse : Mauvaise syntaxe du statement SQLite suivant :")
		log.Println("“INSERT INTO tickets (title, author_id, content, category_id, date, image, state) VALUES (", ticket.Title, ticket.Author_id, ticket.Content, ticket.Date, ticket.State, ")”")
		return 0, err
	}

	execution, err := addStatement.Exec(ticket.Title, ticket.Author_id, ticket.Actual_Admin, ticket.Content, ticket.Date, ticket.State)
	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible d'insérer le tickets dans la base de données.")
		log.Println("Hypothèse : Mauvaise exécution du statement SQLite.")
		log.Println("Ticket à insérer : ", ticket)
		return 0, err
	}
	// Récupération de l'ID du post qui vient d'être ajouté :
	id64, _ := execution.LastInsertId() // ID du dernier post inséré (sous forme de int64)
	id := int(id64)                     // ID casté en int standard

	return id, nil
}

// ⭐ Méthode pour insérer un post dans la base de données :
func (post *Post) InsertIntoDatabase() (int, error) {
	addStatement, err := Db.Prepare("INSERT INTO posts (title, author_id, content, category_id, date, image, state) VALUES (?, ?, ?, ?, ?, ?, ?)")

	defer addStatement.Close()

	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible d'insérer le post dans la base de données.")
		log.Println("Hypothèse : Mauvaise syntaxe du statement SQLite suivant :")
		log.Println("“INSERT INTO posts (title, author_id, content, category_id, date, image, state) VALUES (", post.Title, post.AuthorID, post.Content, post.CategoryID, post.Date, post.Image, post.State, ")”")
		return 0, err
	}

	execution, err := addStatement.Exec(post.Title, post.AuthorID, post.Content, post.CategoryID, post.Date, post.Image, post.State)
	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible d'insérer le post dans la base de données.")
		log.Println("Hypothèse : Mauvaise exécution du statement SQLite.")
		log.Println("Post à insérer : ", post)
		return 0, err
	}
	// Récupération de l'ID du post qui vient d'être ajouté :
	id64, _ := execution.LastInsertId() // ID du dernier post inséré (sous forme de int64)
	id := int(id64)                     // ID casté en int standard

	return id, nil
}

// ⭐ Méthode pour insérer un like de post dans la base de données :
func (like *PostLike) InsertIntoDatabase() error {
	addStatement, err := Db.Prepare("INSERT INTO post_likes (post_id, user_id, type, date) VALUES (?, ?, ?, ?)")
	defer addStatement.Close()
	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible d'insérer le like de post dans la base de données.")
		return err
	}

	addStatement.Exec(like.PostID, like.UserID, like.Type, like.Date)
	return nil
}

// ⭐ Méthode pour insérer un like de commentaire dans la base de données :
func (like *CommentLike) InsertIntoDatabase() error {
	addStatement, err := Db.Prepare("INSERT INTO comment_likes (comment_id, user_id, type, date) VALUES (?, ?, ?, ?)")
	defer addStatement.Close()
	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible d'insérer le like de commentaire dans la base de données.")
		return err
	}

	addStatement.Exec(like.CommentID, like.UserID, like.Type, like.Date)
	return nil
}

// ⭐ Méthode pour supprimer un like de post dans la base de données :
func (like *PostLike) DeleteFromDatabase() error {
	_, err := Db.Exec("DELETE FROM post_likes WHERE post_id = $1 AND user_id = $2 AND type = $3", like.PostID, like.UserID, like.Type)

	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible de supprimer le like de post dans la base de données.")
		return err
	}
	return nil
}

// ⭐ Méthode pour supprimer un like de commentaire dans la base de données :
func (like *CommentLike) DeleteFromDatabase() error {
	_, err := Db.Exec("DELETE FROM comment_likes WHERE comment_id = $1 AND user_id = $2 AND type = $3", like.CommentID, like.UserID, like.Type)

	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible de supprimer le like de commentaire dans la base de données.")
		return err
	}
	return nil
}

// ⭐ Méthode pour supprimer une row et toutes ces données si user en fonction de sont ID :
func DeleteFromDatabase(ID int, table string) error {
	var err error
	if table == "users" {
		_, err = Db.Exec("DELETE FROM users WHERE id = $1", ID)
		_, err = Db.Exec("DELETE FROM posts WHERE author_id = $1", ID)
		_, err = Db.Exec("DELETE FROM comments WHERE author_id = $1", ID)
		_, err = Db.Exec("DELETE FROM session WHERE user_id = $1", ID)
		_, err = Db.Exec("DELETE FROM post_likes WHERE user_id = $1", ID)
		_, err = Db.Exec("DELETE FROM comment_likes WHERE user_id = $1", ID)
		_, err = Db.Exec("DELETE FROM user_badge WHERE user_id = $1", ID)
	} else {
		_, err = Db.Exec("DELETE FROM $1 WHERE id = $1", table, ID)
	}

	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible de supprimer les données de l'utilisateur.")
		return err
	}
	return nil
}
