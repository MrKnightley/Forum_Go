package database

import "log"

func GetCategoryByID(id int) (Category, error) {
	var myCategory Category

	row := Db.QueryRow("SELECT * FROM categories WHERE id = ?", id)
	row.Scan(&myCategory.ID, &myCategory.Name, &myCategory.Theme, &myCategory.Description)

	return myCategory, nil
}

func GetCategoriesList() []Category {
	rows, err := Db.Query("SELECT * FROM categories")
	defer rows.Close()
	if err != nil {
		log.Println("❌ DATABASE | ERREUR : Impossible de récupérer la liste des catégories.")
		panic(err)
	}

	var categories []Category
	for rows.Next() {
		var category Category
		rows.Scan(&category.ID, &category.Name, &category.Theme, &category.Description)
		categories = append(categories, category)
	}
	return categories
}
