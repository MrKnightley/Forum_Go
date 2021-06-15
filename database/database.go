/*Diagramme Merise :
Boite carré : table
Boite arrondie : relation (avec des verbes à l'infinitif)
Attribut souligné : clé primaire
Voir : looping-mcd
*/

package database

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func Initialize() {
	// Déclaration de toutes les tables de la base de données :
	dbTables := []string{
		`CREATE TABLE IF NOT EXISTS "users" (
			"id"				INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			"username"			TEXT UNIQUE NOT NULL,
			"password"			TEXT NOT NULL,
			"email"				TEXT UNIQUE NOT NULL,
			"role"				INTEGER DEFAULT 0,
			"avatar"			TEXT DEFAULT "/images/avatars/defaultAvatar.jpg",
			"date"				DATETIME DEFAULT CURRENT_TIMESTAMP,
			"state"				INTEGER DEFAULT 0,		
			"secretQuestion"	TEXT DEFAULT '',
			"secretAnswer"		TEXT DEFAULT '',
			"house_id"			INTEGER DEFAULT 0
		)`,

		`CREATE TABLE IF NOT EXISTS "posts" (
			"id"			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			"title"			TEXT NOT NULL,
			"author_id"		INTEGER NOT NULL,
			"content"		TEXT NOT NULL,
			"category_id"	INTEGER NOT NULL,
			"date"			DATETIME DEFAULT CURRENT_TIMESTAMP,
			"image"			TEXT DEFAULT '',
			"state"			INTEGER DEFAULT 0,
			"reason"		TEXT DEFAULT "Supprimer par l'utilisateur lui même",
			FOREIGN KEY(author_id) REFERENCES "users"(id), 
			FOREIGN KEY(category_id) REFERENCES "categories"(id) 
		)`,

		`CREATE TABLE IF NOT EXISTS "comments" (
			"id"			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			"author_id"		INTEGER NOT NULL,
			"post_id"		INTEGER NOT NULL,
			"content"		TEXT NOT NULL,
			"gif"			TEXT NOT NULL DEFAULT '',
			"date"			DATETIME DEFAULT CURRENT_TIMESTAMP,
			"state"			INTEGER DEFAULT 0,
			"reason"		TEXT DEFAULT "Supprimer par l'utilisateur lui même",
			FOREIGN KEY(author_id) REFERENCES "users"(id),
			FOREIGN KEY(post_id) REFERENCES "posts"(id)
		)`,

		`CREATE TABLE IF NOT EXISTS "sessions" (
			"id"		INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			"user_id"	INTEGER NOT NULL,
			"uuid"		TEXT NOT NULL,
			"date"		DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES "users"(id)
		)`,

		`CREATE TABLE IF NOT EXISTS "categories" (
			"id"	INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			"name"	TEXT NOT NULL UNIQUE,
			"theme" TEXT NOT NULL DEFAULT ' ',
			"description"	TEXT NOT NULL DEFAULT ' '
		)`,

		`CREATE TABLE IF NOT EXISTS "post_likes" (
			"post_id"		INTEGER NOT NULL,
			"user_id"		INTEGER NOT NULL,
			"type"			TEXT NOT NULL,	
			"date"			DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(post_id) REFERENCES "posts"(id),
			FOREIGN KEY(user_id) REFERENCES "users"(id),
			PRIMARY KEY(post_id, user_id)
		)`,

		`CREATE TABLE IF NOT EXISTS "comment_likes" (
			"comment_id"	INTEGER NOT NULL,
			"user_id"		INTEGER NOT NULL,
			"type"			TEXT NOT NULL,	
			"date"			DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(comment_id) REFERENCES "comments"(id),
			FOREIGN KEY(user_id) REFERENCES "users"(id),
			PRIMARY KEY(comment_id, user_id)
		)`,

		`CREATE TABLE IF NOT EXISTS "badge" (
			"id"			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			"type"			TEXT NOT NULL,
			"image"			TEXT NOT NULL
		)`,

		`CREATE TABLE IF NOT EXISTS "user_badge" (
			"user_id"		INTEGER NOT NULL,
			"badge_id"		INTEGER NOT NULL,
			FOREIGN KEY(user_id) REFERENCES "users"(id),
			FOREIGN KEY(badge_id) REFERENCES "badge"(id),
			PRIMARY KEY(user_id, badge_id)
		)`,

		`CREATE TABLE IF NOT EXISTS "tickets" (
			"id"			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			"author_id"		INTEGER NOT NULL,
			"actual_admin" INTERGER NOT NULL,
			"title"		TEXT NOT NULL,
			"content"		TEXT NOT NULL,
			"date"			DATETIME DEFAULT CURRENT_TIMESTAMP,
			"state"			INTEGER DEFAULT 0,
			FOREIGN KEY(author_id) REFERENCES "users"(id),
			FOREIGN KEY(actual_admin) REFERENCES "users"(id)
		)`,

		`CREATE TABLE IF NOT EXISTS "ticket_answers" (
			"id"			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			"ticket_id"		INTEGER NOT NULL,
			"author_id"		INTEGER NOT NULL,
			"author_name" STRING NOT NULL,
			"content"		TEXT NOT NULL,
			"date"			DATETIME DEFAULT CURRENT_TIMESTAMP,
			"state"			INTEGER DEFAULT 0,
			FOREIGN KEY(ticket_id) REFERENCES "ticket"(id),
			FOREIGN KEY(author_name) REFERENCES "users"(username),
			FOREIGN KEY(author_id) REFERENCES "users"(id)
		)`,

		`CREATE TABLE IF NOT EXISTS "promoted_post" (
			"id"			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			"post_id"		INTEGER NOT NULL,
			FOREIGN KEY(post_id) REFERENCES "posts"(id)
		)`,

		`CREATE TABLE IF NOT EXISTS "houses" (
			"id"			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			"name"			TEXT NOT NULL,
			"image"			TEXT NOT NULL
		)`,
	}

	var err error
	Db, err = sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Println("❌ ERREUR | Impossible de créer le fichier database.db")
		panic(err)
	}

	// Création de chaque table de la base de données :
	for _, table := range dbTables {
		err := createDatabase(table)
		if err != nil {
			panic(err)
		}
	}
	log.Println("✔️ DATABASE | Database created and initialized successfully.")
}

func createDatabase(table string) error {
	statement, err := Db.Prepare(table)
	defer statement.Close()
	if err != nil {
		log.Println("❌ ERREUR | Impossible de créer les tables de la base de données.")
		return err
	}
	statement.Exec()
	return nil
}

func AddSessionToDatabase(w http.ResponseWriter, r *http.Request, user User) error {
	// Je supprime la session précédente de l'utilisateur et en créé une nouvelle :
	Db.Exec("DELETE FROM sessions WHERE user_id = $1", user.ID)
	log.Println("Adding session to database with user's ID : ", user.ID)

	sessionID := uuid.New()
	cookie := &http.Cookie{
		Name:   "session",
		Value:  sessionID.String(),
		Secure: true,
	}
	cookie.MaxAge = 60 * 60 * 24 // 24 heures
	http.SetCookie(w, cookie)

	// Insertion des valeurs de la session dans la table 'sessions' :
	statement, err := Db.Prepare("INSERT INTO sessions (user_id, uuid, date) VALUES (?, ?, ?)")
	defer statement.Close()
	if err != nil {
		log.Println("❌ ERREUR | Impossible d'insérer la session dans la base de données.")
		log.Println("Hypothèse : Mauvaise syntaxe du statement SQLite “INSERT INTO sessions (user_id, uuid, date) VALUES (", user.ID, sessionID, time.Now().Add(24*time.Hour), ")”")
		return err
	}

	statement.Exec(user.ID, sessionID, time.Now().Add(24*time.Hour))
	// statement.Exec(user.ID, sessionID, time.Now().Add(60*time.Minute)) // Heure actuelle + 60 minutes
	return nil
}

func CleanExpiredSessions() {
	for {
		Db.Exec("DELETE FROM sessions WHERE date < $1", time.Now())
		time.Sleep(10 * time.Minute)
	}
}
