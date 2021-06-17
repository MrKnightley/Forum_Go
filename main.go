package main

import (
	"forum/admin"
	"forum/database"
	"forum/request"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func main() {

	// os.Remove("./database/database.db")

	// ⭐ Initialisation de la base de données :
	database.Initialize()

	// database.FillCategories()
	// database.FillHouses()
	// database.FillBadge()
	// database.FillAllUsers()
	// database.FillAllPosts()

	// database.FillUser("Tenebros", "Abc123", "virgil.nauleau@ynov.com", database.ADMIN, database.NORMAL, database.KRAKENS) // Username, Password, Email, Role, Statut, HouseID
	// database.FillUser("Tenebros2", "Abc123", "virgil2.nauleau@ynov.com", 1, database.NORMAL, database.GRIPHONS)           // Username, Password, Email, Role, Statut, HouseID
	// database.FillUser("John Doe", "Abc123", "john.doe@ynov.com", database.MEMBER, database.NORMAL, database.UNAFFILIATED) // Username, Password, Email, Role, Statut, HouseID

	// database.FillPost("World wide best post ever done", 1, "This is the content of the best post ever done", 1, database.PUBLISHED) // Titre, AuthorID, Content, CategoryID, Statut
	// database.FillPost("Deleted post", 1, "This post is deleted", 1, database.UNPUBLISHED)

	// database.FillComment(1, 1, "Best post ever done", database.PUBLISHED, "https://media.tenor.com/images/fcc716e6b70dc2e8e9839280369952a6/tenor.gif") // PostID, AuthorID, Content, Statut, GIF
	// database.FillComment(2, 1, "I agree with myself", database.PUBLISHED, "")
	// database.FillComment(1, 1, "This one is deleted", database.UNPUBLISHED, "")

	// database.FillPostLike(1, 1, "like")
	// database.FillCommentLike(1, 1, "like")

	//Ajout de tickets :

	// database.FillTicket(2, 1, "Problem", "I have a problem", time.Now(), 0)
	// database.FillAnswer(1, 1, "Tenebros", "Answer", time.Now(), 0)

	// ⭐ Serving files to the server :
	images := http.FileServer(http.Dir("./database/images"))
	static := http.FileServer(http.Dir("./static"))

	http.Handle("/images/", http.StripPrefix("/images/", images))
	http.Handle("/static/", http.StripPrefix("/static/", static))

	// ⭐ Ticket :
	http.HandleFunc("/ticket", request.Auth(request.Ticket, "active members only"))

	// ⭐ New Ticket :
	http.HandleFunc("/newTicket", request.Auth(request.NewTicket, "active members only"))

	// ⭐ Ticket Answer :
	http.HandleFunc("/ticket-answer", request.Auth(request.Ticket_Answer, "active members only"))

	// ⭐ Delete User :
	http.HandleFunc("/delete", request.Auth(admin.Delete, "active admins only"))

	// ⭐ Accès aux stats :
	http.HandleFunc("/stats", request.Auth(request.Stats, "everybody"))

	// ⭐ Suppression de post ou commentaire :
	http.HandleFunc("/delete-post", request.Auth(request.DeletePost, "active members only"))

	// ⭐ AddBadgeToUser :
	http.HandleFunc("/badge", request.Auth(admin.AddBadge, "everybody"))

	// ⭐ Accès aux profiles :
	http.HandleFunc("/Profile", request.Auth(request.ProfilePage, "everybody"))

	// ⭐ Accès à la page de modification du compte :
	http.HandleFunc("/edit-account", request.Auth(request.Account, "active members only"))

	// ⭐ modification d'un post :
	http.HandleFunc("/edit-post", request.Auth(request.EditPost, "active members only"))

	// ⭐ Accès aux pages d'inscription, de connexion et déconnexion :
	http.HandleFunc("/register", request.Auth(request.Register, "guests only"))
	http.HandleFunc("/login", request.Auth(request.Login, "guests only"))
	http.HandleFunc("/forgotten-password", request.Auth(request.ForgottenPassword, "guests only"))
	http.HandleFunc("/logout", request.Auth(request.Logout, "active members only"))

	// ⭐ Accès à la page de chaque catégorie :
	http.HandleFunc("/category/", request.Auth(request.Category, "everybody"))

	// ⭐ Accès à la page de chaque post :
	http.HandleFunc("/post/", request.Auth(request.Post, "everybody"))

	// ⭐ Accès à la page 'Nouveau Post' :
	http.HandleFunc("/new-post", request.Auth(request.NewPost, "active members only"))

	// ⭐ Accès à la page de réaction (liker/disliker) :
	http.HandleFunc("/reaction/", request.Auth(request.Reaction, "active members only"))

	// ⭐ Accès au tableau de bord 'Admin' :
	http.HandleFunc("/moderation", request.Auth(admin.Moderation, "active moderators only"))

	// ⭐ Accès au questionnaire pour rejoindre une Maison :
	http.HandleFunc("/join-house", request.Auth(request.JoinHouse, "unaffiliated members only"))

	// ⭐ Accès à la page d'accueil :
	http.HandleFunc("/", request.Auth(request.Index, "everybody"))

	go database.CleanExpiredSessions()

	log.Println("✔️ SERVER | Listening server at port 8000...")
	err := http.ListenAndServeTLS(":8000", "https-server.crt", "https-server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
