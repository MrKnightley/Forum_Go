package admin

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/database"
	"net/http"
	"strconv"
)

func AddBadge(w http.ResponseWriter, r *http.Request, user database.User) {
	fmt.Println("hello")
	switch r.Method {
	case "POST":
		var statement_badge *sql.Stmt
		var received ReceivedData
		err := json.NewDecoder(r.Body).Decode(&received)
		if err != nil {
			panic(err)
		}
		id, err := strconv.Atoi(received.ID)
		if err != nil {
			panic(err)
		}
		badgeID, err := strconv.Atoi(received.Value)
		if err != nil {
			panic(err)
		}
		if received.Category == "add" {
			statement_badge, err = database.Db.Prepare("INSERT OR IGNORE INTO user_badge (user_id,badge_id) VALUES(?,?)")
		} else {
			statement_badge, err = database.Db.Prepare("DELETE FROM user_badge WHERE user_id = ? AND badge_id = ?")
		}

		if err != nil {
			panic(err)
		}
		_, err = statement_badge.Exec(id, badgeID)
		if err != nil {
			panic(err)
		}
		var msg = "{\"message\": \"Successfully Added\"}"
		w.Write([]byte(msg))
	}
}
