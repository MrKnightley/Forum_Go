package database

func GetTicketByUserID(id int) []Ticket {
	var tickets []Ticket
	rows, _ := Db.Query("SELECT * FROM tickets WHERE author_id = ? ORDER BY id DESC", id)
	defer rows.Close()
	for rows.Next() {
		var ticket Ticket
		rows.Scan(&ticket.ID, &ticket.Author_id, &ticket.Actual_Admin, &ticket.Title, &ticket.Content, &ticket.Date, &ticket.State)
		tickets = append(tickets, ticket)
	}
	return tickets
}

func GetAllTickets() []Ticket {
	var tickets []Ticket

	rows, _ := Db.Query("SELECT * FROM tickets")
	defer rows.Close()
	for rows.Next() {
		var ticket Ticket
		rows.Scan(&ticket.ID, &ticket.Author_id, &ticket.Actual_Admin, &ticket.Title, &ticket.Content, &ticket.Date, &ticket.State)
		tickets = append(tickets, ticket)
	}
	return tickets
}

func GetTicketByID(id int) Ticket {
	var ticket Ticket
	Db.QueryRow("SELECT * FROM tickets WHERE id=?", id).Scan(&ticket.ID, &ticket.Author_id, &ticket.Actual_Admin, &ticket.Title, &ticket.Content, &ticket.Date, &ticket.State)
	ticket.Answer = GetAnswerOfTicket(id)
	return ticket
}

func GetAnswerOfTicket(id int) []Ticket_Answer {
	var answers []Ticket_Answer

	rows, _ := Db.Query("SELECT * FROM ticket_answers WHERE Ticket_id=?", id)
	defer rows.Close()
	for rows.Next() {
		var answer Ticket_Answer
		rows.Scan(&answer.ID, &answer.Ticket_id, &answer.Author_id, &answer.Author_name, &answer.Content, &answer.Date, &answer.State)
		answers = append(answers, answer)
	}
	return answers
}

func ResolveTicket(id string) {
	query := "UPDATE tickets SET state = 1 WHERE id=" + id
	_, err := Db.Exec(query)
	if err != nil {
		panic(err)
	}
}
