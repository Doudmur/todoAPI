package service

import (
	"context"
	"net/http"
)

func (s *Server) taskDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	str := s.db.TaskDelete(ctx, r)
	w.Write([]byte(str))
}
