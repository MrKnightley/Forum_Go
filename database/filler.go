package database

import (
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Ajout des catégories dans la base de données :
func FillCategories() {
	statement_categories, err := Db.Prepare("INSERT OR IGNORE INTO categories (name,theme,description) VALUES (?,?,?)")
	if err != nil {
		log.Println("❌ ERREUR | Impossible de remplir la table “categories” dans la base de données.")
		panic(err)
	}
	defer statement_categories.Close()
	statement_categories.Exec("Café", "News", "The Café is the best place to converse with other fellow Fairfaxers.<br><br>Is there a piece of news you would like to discuss? A recent breakthrough or exciting discovery? Say no more!")
	statement_categories.Exec("Men's Club", "Lifestyle", "The Men's Club is the perfect spot to discuss men-focussed topics and share ideas to inspire all our male Fairfaxers.<br><br> Need a tip on how to dress? Advice to improve your relationship with your partner? Seek no more. You're in the right place.")
	statement_categories.Exec("Women's Club", "Lifestyle", "The Women's Club is Fairfax's largest sisterhood. If there is a women-related topic you wish to discuss, you're in the right place.<br><br> Do you want to share a personal advice, talk about a trend or share inspirational ideas? The Women's Club is the perfect spot for you.")
	statement_categories.Exec("Theater", "Cinema", "The Theater is every Fairfaxer's go-to meeting spot to discuss films, TV shows and plays.<br><br> Go and have look if you don't want to miss the latest news.")
	statement_categories.Exec("Arcade", "Video Games", "Come and have a chat in the Arcade to discuss all the latest video game news.")
	statement_categories.Exec("Library", "Literature", "Chances are all your book-loving friends are discussing their favourite (and less favourite) books in the Library right now.<br><br> Are you deciding what to read next? Tell us what titles you’ve enjoyed in the past, and you’ll get surprisingly insightful recommendations.")
	statement_categories.Exec("Vinyl Shop", "Music", "The Viny Shop is the perfect meeting spot to discuss anything related to music.<br><br> Be it your favourite recordings, the latest music news or gossip about your favourite artists, the Vinyl Shop is the place to be.")
	statement_categories.Exec("Museum", "Art & History", "Aaah, the Museum. The one and only place dedicated to history discussions and arts around the world.<br><br> Come and visit the Museum to discuss world history, historical periods, archaeology, arts and cultures with other passionate Fairfaxers.")
	statement_categories.Exec("Academy", "Knowledge & Education", "Are you learning new skills? A new language? Then the Academy is the right place for you.<br><br> In the Academy, you can talk to other students and professionals about courses, schools, skills or any subjects that may require some training and education.")
	statement_categories.Exec("Stadium", "Sports", "The Stadium is every Fairfaxer's go-to meeting place to talk about sports and fitness.<br><br> If you're looking for a place to get work-out training advice or discuss the upcoming World Cup, seek no more! You're in the right place.")
	log.Println("✔️ DATABASE | Filled table “categories” successfully.")
}

// Ajout des maisons dans la base de données :
func FillHouses() {
	statement, err := Db.Prepare("INSERT OR IGNORE INTO houses (name, image) VALUES (?, ?)")
	if err != nil {
		log.Println("❌ ERREUR | Impossible de remplir la table “houses” dans la base de données.")
		panic(err)
	}
	defer statement.Close()
	statement.Exec("Kingdsbridge Griphons", "/images/houses/house-1.png")
	statement.Exec("Westside Wildcats", "/images/houses/house-2.png")
	statement.Exec("Columbus Krakens", "/images/houses/house-3.png")
	statement.Exec("Syracuse Vipers", "/images/houses/house-4.png")

	log.Println("✔️ DATABASE | Filled table “houses” successfully.")
}

// Ajout d'un utilisateur dans la base de données :
func FillUser(username string, password string, email string, role int, state int, houseID int) {
	statement, err := Db.Prepare("INSERT OR IGNORE INTO users (username, password, email, role, date, state, avatar, secretQuestion, secretAnswer, house_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("❌ ERREUR | Impossible de remplir la table “users” dans la base de données.")
		panic(err)
	}
	defer statement.Close()

	var newUser User

	newUser.Username = strings.ToLower(username)
	cryptedPassword, err2 := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err2 != nil {
		panic(err)
	}
	newUser.Password = string(cryptedPassword)
	newUser.Email = strings.ToLower(email)
	newUser.Role = role
	newUser.Date = time.Now()
	newUser.State = state
	newUser.Avatar = "/images/avatars/defaultAvatar.jpg"
	newUser.SecretQuestion = ""
	newUser.SecretAnswer = ""
	newUser.House.ID = houseID

	statement.Exec(newUser.Username, newUser.Password, newUser.Email, newUser.Role, newUser.Date, newUser.State, newUser.Avatar, newUser.SecretQuestion, newUser.SecretAnswer, newUser.House.ID)
	log.Println("✔️ DATABASE | Filled table “users” successfully.")
	log.Println("User added: ", newUser)
}

// Ajout d'un post, comment et like dans la base de données :
func FillPost(title string, authorID int, content string, categoryID int, state int) {
	statement, err := Db.Prepare("INSERT INTO posts (title, author_id, content, category_id, date,image, state) VALUES (?, ?, ?, ?, ?,?, ?)")
	if err != nil {
		log.Println("❌ ERREUR | Impossible de remplir la table “posts” dans la base de données.")
		panic(err)
	}
	defer statement.Close()

	var newPost Post

	newPost.Title = title
	newPost.AuthorID = authorID
	newPost.Content = content
	newPost.CategoryID = categoryID
	newPost.Date = time.Now()
	newPost.State = state

	statement.Exec(newPost.Title, newPost.AuthorID, newPost.Content, newPost.CategoryID, newPost.Date, " ", newPost.State)
	log.Println("✔️ DATABASE | Filled table “posts” successfully.")
	log.Println("Post added: ", newPost)
}

func FillComment(postID int, authorID int, content string, state int, gif string) {
	statement, err := Db.Prepare("INSERT INTO comments (author_id, post_id, content, gif, date, state) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("❌ ERREUR | Impossible de remplir la table “comments” dans la base de données.")
		panic(err)
	}
	defer statement.Close()

	var newComment Comment

	newComment.AuthorID = authorID
	newComment.PostID = postID
	newComment.Content = content
	newComment.Date = time.Now()
	newComment.State = state

	statement.Exec(newComment.AuthorID, newComment.PostID, newComment.Content, gif, newComment.Date, newComment.State)
	log.Println("✔️ DATABASE | Filled table “comments” successfully.")
	log.Println("Post added: ", newComment)
}

func FillPostLike(postID int, userID int, reactionType string) {
	statement, err := Db.Prepare("INSERT INTO post_likes (post_id, user_id, type) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("❌ ERREUR | Impossible de remplir la table “post_likes” dans la base de données.")
		panic(err)
	}
	defer statement.Close()

	var newLike PostLike

	newLike.PostID = postID
	newLike.UserID = userID
	newLike.Type = reactionType

	statement.Exec(newLike.PostID, newLike.UserID, newLike.Type)
	log.Println("✔️ DATABASE | Filled table “post_likes” successfully.")
	log.Println("Like added: ", newLike)
}

func FillCommentLike(commentID int, userID int, reactionType string) {
	statement, err := Db.Prepare("INSERT INTO comment_likes (comment_id, user_id, type) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("❌ ERREUR | Impossible de remplir la table “comment_likes” dans la base de données.")
		panic(err)
	}
	defer statement.Close()

	var newLike CommentLike

	newLike.CommentID = commentID
	newLike.UserID = userID
	newLike.Type = reactionType

	statement.Exec(newLike.CommentID, newLike.UserID, newLike.Type)
	log.Println("✔️ DATABASE | Filled table “comment_likes” successfully.")
	log.Println("Like added: ", newLike)
}

// Ajout d'un ticket dans la base de données :
func FillTicket(authorID int, actual_admin int, title string, content string, date time.Time, state int) {
	statement, err := Db.Prepare("INSERT INTO tickets (author_id,actual_admin,title, content, date, state) VALUES (?, ?, ?, ?, ?,?)")
	if err != nil {
		log.Println("❌ ERREUR | Impossible de remplir la table “posts” dans la base de données.")
		panic(err)
	}
	defer statement.Close()

	var newTicket Ticket

	newTicket.Title = title
	newTicket.Author_id = authorID
	newTicket.Actual_Admin = actual_admin
	newTicket.Content = content
	newTicket.Date = time.Now()
	newTicket.State = state

	statement.Exec(newTicket.Author_id, newTicket.Actual_Admin, newTicket.Title, newTicket.Content, newTicket.Date, newTicket.State)
	log.Println("✔️ DATABASE | Filled table “tickets” successfully.")
	log.Println("Ticket added: ", newTicket)
}
func FillAnswer(authorID int, ticket_id int, name string, content string, date time.Time, state int) {
	statement, err := Db.Prepare("INSERT INTO ticket_answers (author_id, ticket_id,author_name,content, date, state) VALUES (?,?,?, ?, ?, ?)")
	if err != nil {
		log.Println("❌ ERREUR | Impossible de remplir la table “posts” dans la base de données.")
		panic(err)
	}
	defer statement.Close()

	var newAnswer Ticket_Answer

	newAnswer.Author_id = authorID
	newAnswer.Content = content
	newAnswer.Date = time.Now()
	newAnswer.State = state

	statement.Exec(newAnswer.Author_id, ticket_id, name, newAnswer.Content, newAnswer.Date, newAnswer.State)
	log.Println("✔️ DATABASE | Filled table “ticket_answer” successfully.")
	log.Println("Answer added: ", newAnswer)
}

// Ajout des badges dans la base de données :
func FillBadge() {
	statement_badge, _ := Db.Prepare("INSERT OR IGNORE INTO badge (type,image) VALUES(?,?)")
	statement_badge.Exec("5like", "/images/badges/bronze.png")
	statement_badge.Exec("10like", "/images/badges/silver.png")
	statement_badge.Exec("15like", "/images/badges/gold.png")
	statement_badge.Exec("20like", "/images/badges/platinum.png")
}

// Ajout de toute une collection d'utilisateurs dans la base de données :

func FillAllUsers() {

	// Username, Password, Email, Role, Statut, HouseID
	FillUser("Tenebros", "Abc123", "virgil.nauleau@ynov.com", ADMIN, NORMAL, KRAKENS)
	FillUser("Nicolas L.", "Abc123", "nicolas.lepinay@ynov.com", MODERATOR, NORMAL, VIPERS)
	FillUser("Alice Kepler", "Abc123", "alice.kepler@outlook.com", MEMBER, NORMAL, GRIPHONS)
	FillUser("James L. Wright", "Abc123", "james.wright@outlook.com", MEMBER, NORMAL, WILDCATS)
	FillUser("Lily Cavendish", "Abc123", "lily.cavendish@outlook.com", MEMBER, NORMAL, KRAKENS)
	FillUser("Donnie Bryant", "Abc123", "donnie.bryant@outlook.com", MEMBER, NORMAL, VIPERS)
	FillUser("Alin Mela", "Abc123", "alin.mela@outlook.com", MEMBER, NORMAL, GRIPHONS)
	FillUser("Cyprus S.", "Abc123", "cyprus.s@outlook.com", MEMBER, NORMAL, WILDCATS)
	FillUser("Mari Hashiba", "Abc123", "mari.hashiba@outlook.com", MEMBER, NORMAL, KRAKENS)
	FillUser("Giulio Favaro", "Abc123", "giulio.favaro@outlook.com", MEMBER, NORMAL, VIPERS)
	FillUser("Anna Kaufmann", "Abc123", "anna.kaufmann@outlook.com", MEMBER, NORMAL, GRIPHONS)
	FillUser("Rafaelo Neixeira", "Abc123", "rafaelo.neixeira@outlook.com", MEMBER, NORMAL, WILDCATS)
	FillUser("Joy Woodley", "Abc123", "joy.woodley@outlook.com", MEMBER, NORMAL, KRAKENS)
	FillUser("Jensen Acker", "Abc123", "jensen.acker@outlook.com", MEMBER, NORMAL, VIPERS)
	FillUser("Betty L. Berrie", "Abc123", "betty.berrie@outlook.com", MEMBER, NORMAL, GRIPHONS)
	FillUser("Jeanne Dulcy", "Abc123", "jeanne.dulcy@outlook.com", MEMBER, NORMAL, WILDCATS)
}
