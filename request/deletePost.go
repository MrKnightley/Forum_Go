package request

import (
	"encoding/json"
	"forum/database"
	"net/http"
)

func DeletePost(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {
	case "POST":
		var received database.ReceivedData
		err := json.NewDecoder(r.Body).Decode(&received)
		if err != nil {
			ERROR, _ := json.Marshal("ERROR WHILE DECODING JSON")
			w.Write(ERROR)
			panic(err)
		}
		query := "UPDATE " + received.Table + " SET state = 1 WHERE id=" + received.ID
		_, err = database.Db.Exec(query)
		if err != nil {
			ERROR, _ := json.Marshal("ERROR")
			w.Write(ERROR)
			panic(err)
		}
		var msg = "{\"message\": \"Successfully Deleted\"}"
		w.Write([]byte(msg))
	}
}
