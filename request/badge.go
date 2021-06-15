package request

import (
	"forum/database"
)

func AddBadgeIfUnlocked(user database.User) {
	total := getTotalLikeOfUser(user.ID)
	if total > 0 {
		database.Db.Exec("INSERT OR IGNORE INTO user_badge (user_id,badge_id) VALUES(?,?)", user.ID, 1)
	}
	if total > 10 {
		database.Db.Exec("INSERT OR IGNORE INTO user_badge (user_id,badge_id) VALUES(?,?)", user.ID, 2)
	}
	if total > 15 {
		database.Db.Exec("INSERT OR IGNORE INTO user_badge (user_id,badge_id) VALUES(?,?)", user.ID, 3)
	}
	if total > 20 {
		database.Db.Exec("INSERT OR IGNORE INTO user_badge (user_id,badge_id) VALUES(?,?)", user.ID, 4)
	}

}

func getTotalLikeOfUser(id int) int {
	var result int
	database.Db.QueryRow(`SELECT count(*) as Compte FROM posts p 
	INNER JOIN post_likes pl ON pl.post_id = p.id
	WHERE p.author_id = $1 and pl.type = "like"`, id).Scan(&result)
	return result
}
