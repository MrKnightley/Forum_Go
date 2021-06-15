package admin

import (
	"encoding/json"
	"forum/database"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request, user database.User) {
	switch r.Method {
	case "POST":
		var received receivedData
		err := json.NewDecoder(r.Body).Decode(&received)
		if err != nil {
			ERROR, _ := json.Marshal("ERROR WHILE DECODING JSON")
			w.Write(ERROR)
			panic(err)
		}
		id, err := strconv.Atoi(received.ID)
		err = database.DeleteFromDatabase(id, received.Table)
		if err != nil {
			ERROR, _ := json.Marshal("ERROR")
			w.Write(ERROR)
			panic(err)
		}
		var msg = "{\"message\": \"Successfully Deleted\"}"
		w.Write([]byte(msg))
	}
}
