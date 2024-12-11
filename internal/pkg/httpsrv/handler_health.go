package httpsrv

import (
	"net/http"
)

func (s *Server) handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
