package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
	"time"
	"todoAPI/internal/models"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "12345"
	dbname   = "todo"
)

type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	dbconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbconn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) GetTasks(ctx context.Context) ([]models.Task, error) {
	selectResult, err := s.db.Query("SELECT * FROM tasks;")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var tasks []models.Task
	for selectResult.Next() {
		var task models.Task
		err = selectResult.Scan(&task.Id, &task.Name, &task.Date, &task.IsDone)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *Storage) TaskAdd(ctx context.Context, r *http.Request) (string, error) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Println("Error with decoding")
	}
	_, err = s.db.Exec("insert into tasks(name, date, isdone) values($1, $2, $3)", task.Name, time.Now(), task.IsDone)
	if err != nil {
		log.Println("Error with insert task")
	}
	return "Task " + task.Name + " was successfully added", nil
}

func (s *Storage) TaskDelete(ctx context.Context, r *http.Request) string {
	id := mux.Vars(r)["id"]
	_, err := s.db.Exec("delete from tasks where id=$1", id)
	if err != nil {
		log.Println("Task was not found!")
		return "Error"
	}
	return "Task was successfully deleted"
}

func (s *Storage) TaskUpdate(ctx context.Context, r *http.Request) string {
	id := mux.Vars(r)["id"]
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	log.Println(err)
	if err != nil {
		log.Println("Something wrong with decoding")
		return ""
	}
	if task.Name == "" {
		_, err = s.db.Exec("update tasks set isdone=$1 where id=$2", task.IsDone, id)
		return "Task status was successfully changed to " + strconv.FormatBool(task.IsDone)
	} else {
		_, err = s.db.Exec("update tasks set name=$1 where id=$2", task.Name, id)
		return "Task name was successfully changed to " + task.Name
	}
}
