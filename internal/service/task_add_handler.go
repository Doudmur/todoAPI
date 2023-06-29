package service

import (
	"context"
	"net/http"
)

func (s *Server) taskAddHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res, _ := s.db.TaskAdd(ctx, r)
	w.Write([]byte(res))
}
