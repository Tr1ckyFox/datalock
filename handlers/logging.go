package handlers

import (
	"log"
	"net/http"

	"github.com/leominov/datalock/utils"
)

const (
	LogginFormat = "%s - \"%s %s %s\" %s\n"
)

func LoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := utils.RealIP(r)
		log.Printf(LogginFormat, ip, r.Method, r.URL.Path, r.Proto, r.UserAgent())
		h.ServeHTTP(w, r)
	})
}
