package request

import (
	"fmt"
	"forum/database"
	"log"
	"net/http"
)

func JoinHouse(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {

	// 🍔 Méthode 'GET' — Arrivée sur le questionnaire :
	case "GET":
		err := MyTemplates.ExecuteTemplate(w, "join-house", user)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			log.Println("❌ ERREUR | Impossible d'exécuter le template “houses”.")
			fmt.Println(err)
			return
		}

	// 🍔 Méthode 'POST' — Envoi des réponses du questionnaire :
	case "POST":
		// Je récupère le résultat de chacune des 13 questions :
		q1 := r.FormValue("q1")
		q2 := r.FormValue("q2")
		q3 := r.FormValue("q3")
		q4 := r.FormValue("q4")
		q5 := r.FormValue("q5")
		q6 := r.FormValue("q6")
		q7 := r.FormValue("q7")
		q8 := r.FormValue("q8")
		q9 := r.FormValue("q9")
		q10 := r.FormValue("q10")
		q11 := r.FormValue("q11")
		q12 := r.FormValue("q12")
		q13 := r.FormValue("q13")

		// Je concatène toutes ces valeurs et crée des compteurs :

		result := q1 + q2 + q3 + q4 + q5 + q6 + q7 + q8 + q9 + q10 + q11 + q12 + q13

		var griphonsCounter, wildcatsCounter, krakensCounter, vipersCounter int

		for i := 0; i < len(result); i++ {
			if result[i] == '1' {
				griphonsCounter++
			} else if result[i] == '2' {
				wildcatsCounter++
			} else if result[i] == '3' {
				krakensCounter++
			} else if result[i] == '4' {
				vipersCounter++
			}
		}
		log.Println("📑 JOIN HOUSE | Griphons Points : ", griphonsCounter)
		log.Println("📑 JOIN HOUSE | Wildcats Points : ", wildcatsCounter)
		log.Println("📑 JOIN HOUSE | Krakens Points : ", krakensCounter)
		log.Println("📑 JOIN HOUSE | Vipers Points : ", vipersCounter)

		// J'ajoute les 4 compteurs dans un tableau :
		countersArr := []int{griphonsCounter, wildcatsCounter, krakensCounter, vipersCounter}

		// Je cherche la valeur MAX du tableau :
		max := countersArr[0]
		for _, value := range countersArr {
			if value > max {
				max = value
			}
		}

		// Détermination de l'ID de la maison :
		if max == griphonsCounter {
			user.House.ID = 1
		}
		if max == wildcatsCounter {
			user.House.ID = 2
		}
		if max == krakensCounter {
			user.House.ID = 3
		}
		if max == vipersCounter {
			user.House.ID = 4
		}
		log.Println("📑 JOIN HOUSE | House ID : ", user.House.ID)

		// Remplissage du champ House de la struct User :
		user.House = database.GetHouseByID(user.House.ID)

		// Mise à jour de l'utilisateur dans la base de données :
		err := user.UpdateInDatabase("house_id")
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}

		err = MyTemplates.ExecuteTemplate(w, "discover-your-house", user)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
