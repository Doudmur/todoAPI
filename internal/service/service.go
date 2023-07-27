package service

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"todoAPI/internal/models"
)

type Server struct {
	db Storage
}

type Storage interface {
	GetTasks(ctx context.Context) ([]models.Task, error)
	TaskAdd(ctx context.Context, r *http.Request) (string, error)
	TaskDelete(ctx context.Context, r *http.Request) string
	TaskUpdate(ctx context.Context, r *http.Request) string
	Shutdown()
}

func New(db Storage) (*Server, error) {
	return &Server{db: db}, nil
}

func (s *Server) Run() error {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/tasks", s.getTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/add", s.taskAddHandler).Methods("POST")
	router.HandleFunc("/tasks/delete/{id}", s.taskDeleteHandler).Methods("POST")
	router.HandleFunc("/tasks/update/{id}", s.taskUpdateHandler).Methods("POST")

	return http.ListenAndServe(viper.GetString("port"), router)
}

func (s *Server) Shutdown() error {
	s.db.Shutdown()
	return nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
