package models

import (
	"time"
)

type Event struct {
	Id          int
	Name        string 	`binding:"required"`  // struct tag in Go, Gin-specific validation tag
	Description string  `binding:"required"`
	Location    string	`binding:"required"`
	DateTime    time.Time	`binding:"required"`
	UesrId      int
}

var events = []Event{}

func (e Event) Save(){
	// add it to a DB
	events = append(events, e)
}

func GetAllEvents () []Event{
	return events
}

