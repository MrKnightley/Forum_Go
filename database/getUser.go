package database

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func GetUserByCookie(w http.ResponseWriter, r *http.Request) (User, error) {
	// Lecture du cookie :
	userCookie, err := r.Cookie("session")
	// Si le cookie n'existe pas, on le créé :
	if err != nil {
		sessionID := uuid.New()
		userCookie = &http.Cookie{
			Name:   "session",
			Value:  sessionID.String(),
			Secure: true,
		}
		userCookie.MaxAge = 60 * 60 * 24
		http.SetCookie(w, userCookie)
	}
	// Récupération de la session grâce au cookie :
	log.Println("🍪 Cookie | ", userCookie)
	session := GetSessionByUUID(userCookie.Value)

	// Récupération de l'utilisateur grâce à sa session :
	var user User
	user, err = GetUserByID(session.UserID)
	if err != nil {
		log.Println("❌ ERREUR | Impossible de récupérer l'utilisateur dont l'ID est ", session.UserID)
		return user, err
	}
	log.Println("🌐 Session de l'utilisateur n°", session.UserID)

	return user, nil
}

func GetUserByID(id int) (User, error) {
	var user User

	row := Db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Role, &user.Avatar, &user.Date, &user.State, &user.SecretQuestion, &user.SecretAnswer, &user.House.ID)
	user.Badges = GetBadgeByUserID(user.ID)
	user.House = GetHouseByID(user.House.ID)

	log.Println("🦸 Get User By ID | User : ", user)
	return user, nil
}

func GetUserByUsernameOrEmail(identifier string) (User, error) {
	var user User

	row := Db.QueryRow("SELECT * FROM users WHERE username = $1 OR email = $1", identifier)

	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Role, &user.Avatar, &user.Date, &user.State, &user.SecretQuestion, &user.SecretAnswer, &user.House.ID)
	user.Badges = GetBadgeByUserID(user.ID)
	user.House = GetHouseByID(user.House.ID)

	log.Println("🦸 Get User By Username Or Email | User : ", user)
	return user, nil

	// Si le username ou l'email n'existe pas, user.ID == 0 (car par défaut, variable de type int = 0)
}
