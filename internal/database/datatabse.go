package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"todoAPI/etc/logger"
	"todoAPI/internal/models"
)

type Storage struct {
	db *sql.DB
}

func New() (*Storage, error) {
	logger.SetErrorLevel(4)
	ctx := context.Background()

	if err := initConfig(); err != nil {
		logger.Errorf(ctx, "Error config ", err)
	}

	if err := godotenv.Load(); err != nil {
		logger.Errorf(ctx, "Error with environment variables", err)
	}

	host := viper.GetString("db.host")
	port := viper.GetInt("db.port")
	user := viper.GetString("db.user")
	password := os.Getenv("DB_PASSWORD")
	dbname := viper.GetString("db.name")
	dbconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbconn)
	if err != nil {
		logger.Errorf(ctx, "Error with connection to storage!", err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Shutdown() {
	s.db.Close()
}

func (s *Storage) GetTasks(ctx context.Context) ([]models.Task, error) {
	logger.SetErrorLevel(4)
	selectResult, err := s.db.Query("SELECT * FROM tasks;")
	if err != nil {
		logger.Errorf(ctx, "Error with selection!", err)
		return nil, err
	}

	var tasks []models.Task
	for selectResult.Next() {
		var task models.Task
		err = selectResult.Scan(&task.Id, &task.Name, &task.Date, &task.IsDone)
		if err != nil {
			logger.Errorf(ctx, "Error with scan!", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *Storage) TaskAdd(ctx context.Context, r *http.Request) (string, error) {
	logger.SetErrorLevel(4)
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		logger.Errorf(ctx, "Error with decoding!", err)
		return "", nil
	}
	_, err = s.db.Exec("insert into tasks(name, date, isdone) values($1, $2, $3)", task.Name, time.Now(), task.IsDone)
	if err != nil {
		logger.Errorf(ctx, "Error with insert to db!", err)
		return "", nil
	}
	return "Task " + task.Name + " was successfully added", nil
}

func (s *Storage) TaskDelete(ctx context.Context, r *http.Request) string {
	logger.SetErrorLevel(4)
	id := mux.Vars(r)["id"]
	_, err := s.db.Exec("delete from tasks where id=$1", id)
	if err != nil {
		logger.Errorf(ctx, "Error with deleting!", err)
		return "Error"
	}
	return "Task was successfully deleted"
}

func (s *Storage) TaskUpdate(ctx context.Context, r *http.Request) string {
	logger.SetErrorLevel(4)
	id := mux.Vars(r)["id"]
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	log.Println(err)
	if err != nil {
		logger.Errorf(ctx, "Error with decoding!", err)
		return ""
	}
	if task.Name == "" {
		_, err = s.db.Exec("update tasks set isdone=$1 where id=$2", task.IsDone, id)
		if err != nil {
			logger.Errorf(ctx, "Error with updating!", err)
			return ""
		}

		return "Task status was successfully changed to " + strconv.FormatBool(task.IsDone)
	} else {
		_, err = s.db.Exec("update tasks set name=$1 where id=$2", task.Name, id)
		if err != nil {
			logger.Errorf(ctx, "Error with updating!", err)
			return ""
		}
		return "Task name was successfully changed to " + task.Name
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
