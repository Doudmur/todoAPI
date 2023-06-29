package service

import (
	"context"
	"net/http"
)

func (s *Server) taskUpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	str := s.db.TaskUpdate(ctx, r)
	w.Write([]byte(str))
}
