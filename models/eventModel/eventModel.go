package eventModel

import (
	"fmt"
	"time"

	"github.com/githubak2002/golang-event-api/db"
)

type Event struct {
	Id          int64
	Name        string 	`binding:"required"`  // struct tag in Go, Gin-specific validation tag
	Description string  `binding:"required"`
	Location    string	`binding:"required"`
	DateTime    time.Time	`binding:"required"`
	UserId      int
}

func (e Event) Save() error { 
	// add it to a SQL lite DB
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?,?,?,?,?)`

	// VALUES (?,?,?,?,?)   
	// SQL statements helps protect against SQL injection.

	stmt,err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserId)
	if err != nil{
		return err
	}
	id, err := result.LastInsertId()
	e.Id = id
	return err
	// events = append(events, e)
}

func GetAllEvents () ([]Event, error){

	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil{
		return nil, err
	}
	
	// ensures the database connection is closed when the function ends.
	defer rows.Close()
	var events []Event
	for rows.Next(){
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil{
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById (id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil{
		return nil, err
	}
	return &event, nil
}





func (e Event) UpdateEvent() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)
	return err
}

