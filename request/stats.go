package request

import (
	"forum/database"
	"log"
	"net/http"
)

type DataStats struct {
	Categories []database.Category
	//nombre de posts
	LifePostResult    []int
	MonthlyPostResult []int
	WeeklyPostResult  []int
	DailyPostResult   []int
	//nombre de commentaires
	LifeCommentResult    []int
	MonthlyCommentResult []int
	WeeklyCommentResult  []int
	DailyCommentResult   []int
	//nombre de likes
	LifeLikeResult    []int
	MonthlyLikeResult []int
	WeeklyLikeResult  []int
	DailyLikeResult   []int
	//nombre de dislikes
	LifeDislikeResult    []int
	MonthlyDislikeResult []int
	WeeklyDislikeResult  []int
	DailyDislikeResult   []int
	//On pourrais rajouter un nombre de visite de chaque personne (une session crée = une visite possible sur chaque catégories)
}

// HandleFunc pour la page catégorie :
func Stats(w http.ResponseWriter, r *http.Request, user database.User) {
	var data DataStats
	data.Categories = database.GetCategoriesList()
	//Nombre de posts de chaque catégorie
	for y := 0; y < len(data.Categories); y++ {
		//append new cell
		data.LifePostResult = append(data.LifePostResult, 0)
		data.MonthlyPostResult = append(data.MonthlyPostResult, 0)
		data.WeeklyPostResult = append(data.WeeklyPostResult, 0)
		data.DailyPostResult = append(data.DailyPostResult, 0)
		data.LifeCommentResult = append(data.LifeCommentResult, 0)
		data.MonthlyCommentResult = append(data.MonthlyCommentResult, 0)
		data.WeeklyCommentResult = append(data.WeeklyCommentResult, 0)
		data.DailyCommentResult = append(data.DailyCommentResult, 0)
		data.LifeLikeResult = append(data.LifeLikeResult, 0)
		data.MonthlyLikeResult = append(data.MonthlyLikeResult, 0)
		data.WeeklyLikeResult = append(data.WeeklyLikeResult, 0)
		data.DailyLikeResult = append(data.DailyLikeResult, 0)
		data.LifeDislikeResult = append(data.LifeDislikeResult, 0)
		data.MonthlyDislikeResult = append(data.MonthlyDislikeResult, 0)
		data.WeeklyDislikeResult = append(data.WeeklyDislikeResult, 0)
		data.DailyDislikeResult = append(data.DailyDislikeResult, 0)
		//
		i := y + 1
		//Post side
		database.GetNumberOfPostByDateAndPostCategory(i, &data.LifePostResult[y], &data.MonthlyPostResult[y], &data.WeeklyPostResult[y], &data.DailyPostResult[y])
		//Com side
		database.GetNumberOfCommentByDateAndPostCategory(i, &data.LifeCommentResult[y], &data.MonthlyCommentResult[y], &data.WeeklyCommentResult[y], &data.DailyCommentResult[y])
		//Like side
		database.GetNumberOfReactionByDate(i, "like", &data.LifeLikeResult[y], &data.MonthlyLikeResult[y], &data.WeeklyLikeResult[y], &data.DailyLikeResult[y])
		//Dislike side
		database.GetNumberOfReactionByDate(i, "dislike", &data.LifeDislikeResult[y], &data.MonthlyDislikeResult[y], &data.WeeklyDislikeResult[y], &data.DailyDislikeResult[y])
	}

	err := MyTemplates.ExecuteTemplate(w, "stats", data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("❌ ERREUR | Impossible d'exécuter le template “post”.")
		return
	}
}
