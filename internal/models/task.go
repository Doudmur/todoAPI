package models

import "time"

type Task struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
	IsDone bool      `json:"isDone"`
}
