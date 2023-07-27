package service

import (
	"context"
	"encoding/json"
	"net/http"
	"todoAPI/etc/logger"
)

func (s *Server) getTasksHandler(w http.ResponseWriter, r *http.Request) {
	logger.SetErrorLevel(4)
	ctx := context.Background()
	tasks, err := s.db.GetTasks(ctx)
	result, err := json.Marshal(tasks)
	if err != nil {
		logger.Errorf(ctx, "No result in tasks", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(result)
}
