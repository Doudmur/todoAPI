package service

import (
	"context"
	"net/http"
	"todoAPI/etc/logger"
)

func (s *Server) taskAddHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.NewLogger()
	ctx := context.Background()
	res, err := s.db.TaskAdd(ctx, r)
	if err != nil {
		l.Error("Error with adding to db!", err)
	}
	w.Write([]byte(res))
}
