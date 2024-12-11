package httpsrv

import (
	"log"
	"net/http"
)

func (s *Server) error(w http.ResponseWriter, code int, err error) {
	log.Printf("error: %+v\n", err)
	http.Error(w, http.StatusText(code), code)
}
