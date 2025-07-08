package db

import (

	"database/sql"
	
	// not used directly
	// _ "github.com/mattn/go-sqlite3"


	_ "modernc.org/sqlite"

)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite", "api.db")
 
    if err != nil {
		// as we are using the gin default engin we can use panic and it will log the err 
        panic("Could not connect to database.")
    }
 
	
	// At most 10 active DB queries can run at the same time.
	// If 15 users hit the API at the same time:
		// First 10 get DB connections.
		// Remaining 5 will wait (not fail) — Go's database/sql queues them.
    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)
 
    createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	);
	`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create Events Table: " + err.Error())
		// panic("Could not create Events Table")
	}
}



 
