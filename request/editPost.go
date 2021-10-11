package request

import (
	"encoding/json"
	"forum/database"
	"net/http"
)

func EditPost(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {
	case "POST":
		var p ReceivedData
		var query string
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			panic(err)
		}
		if p.Table == "gif" {
			query = `UPDATE comments SET gif = "" WHERE id = "` + p.ID + `"`
		} else {
			query = `UPDATE ` + p.Table + ` SET content = "` + p.NewValue + `" WHERE id = "` + p.ID + `"`
		}
		_, err = database.Db.Exec(query)
		if err != nil {
			panic(err)
		}
		msg := `{"newVal":"` + p.NewValue + `"}`
		w.Write([]byte(msg))
	}
}
