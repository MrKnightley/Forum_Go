package request

import (
	"fmt"
	"forum/database"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/google/uuid"
)

// Fonction handleFunc pour la page 'Forgotten Password' :
func ForgottenPassword(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {

	// 🍔 Méthode 'GET' — Lorsqu'on arrive sur la page register.html pour la 1ère fois :
	case "GET":
		err := MyTemplates.ExecuteTemplate(w, "forgotten-password", "")
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

	// 🍔 Méthode 'POST' — Lorsqu'on sur le bouton 'Create your account' pour s'enregistrer :
	case "POST":
		// (1) Je récupère l'email saisi :
		email := r.FormValue("email")

		// (2) Je recherche l'utilisateur grâce à son email :
		user, err := database.GetUserByUsernameOrEmail(email)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		// (3) Si l'utilisateur n'existe pas...
		if user.ID == 0 {
			message := "INVALID EMAIL"
			MyTemplates.ExecuteTemplate(w, "forgotten-password", message) // On ré-exécute le template en affichant cette fois une div "Email non-reconnu".
			log.Println("❌ FORGOTTEN PASSWORD | Email inconnu.")
			return
		}

		// (4) ...Sinon, je mets à jour le mot de passe dans la base de données :

		newPassword := GenerateNewPassword()
		// fmt.Println("NEW PASSWORD : ")
		// fmt.Println(newPassword)
		user.Password = newPassword // Le cryptage se fait dans la fonction UpdateInDatabase !

		err = user.UpdateInDatabase("password")
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		// (5) J'envoie un e-mail à l'utilisateur avec son nouveau mot de passe :
		err = SendEmail(email, user.Username, newPassword)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		// (6) Je ré-exécute le template :
		message := "VALID EMAIL"
		MyTemplates.ExecuteTemplate(w, "forgotten-password", message) // On ré-exécute le template en affichant cette fois une div "Un email a été envoyé".
		log.Println("✔️ FORGOTTEN PASSWORD | Email envoyé.")
		return
	}
}

// Fonction de génération d'un mot de passe aléatoire :
func GenerateNewPassword() string {
	uuid := uuid.New().String()
	randomPassword := uuid[0:14] + "Z" // Les 15 premiers caractères de l'UUID + un 'Z' pour avoir une majuscule
	return randomPassword
}

// Fonction d'envoi d'un email de confirmation de changement de mot de passe :
func SendEmail(toEmail, toUsername, newPassword string) error {
	// Coordonnées de l'expéditeur :
	from := "forum.fairfax@gmail.com"
	password := "Abcd1234?"

	// Coordonnées du destinataire :
	to := []string{toEmail}

	// Serveur SMTP :
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	// Contenu de l'email :
	subject := "Subject: FAIRFAX | Your password has been changed\n"

	body := fmt.Sprintf(`Hello %s,

A request to reset your Fairfax account password was sent today.

Your new randomly generated password is:

	%s

Please keep in mind it is only intended as a temporary password and should be customised in your Profile page.

- Nicolas, Administrator
	`, strings.Title(toUsername), newPassword)

	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host) // Identity, username, pwd, host
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		log.Println("❌ FORGOTTEN PASSWORD | L'envoi de l'email a échoué.")
		fmt.Println(err)
		return err
	}

	return nil
}
