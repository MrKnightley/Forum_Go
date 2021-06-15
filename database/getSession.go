package database

import "log"

func GetSessionByUUID(uuid string) Session {
	var session Session

	// Dans la table 'session', je stocke toutes les variables de la ligne dont l'UUID est 'uuid' dans la variable 'row' :
	row := Db.QueryRow("SELECT * FROM sessions WHERE uuid = $1", uuid)
	// On stocke ces variables dans la struct 'session' :
	row.Scan(&session.ID, &session.UserID, &session.UUID, &session.Date)
	log.Println("🌐 GET SESSION (ID, UserID, UUID, Date) | ", session)

	return session
}
