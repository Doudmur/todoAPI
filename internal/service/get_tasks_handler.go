package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) getTasksHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	tasks, err := s.db.GetTasks(ctx)
	result, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("No result in tasks: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(result)
}
