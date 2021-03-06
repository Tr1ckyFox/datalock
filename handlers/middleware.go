package handlers

import (
	"log"
	"net/http"

	"github.com/leominov/datalock/server"
	"github.com/leominov/datalock/utils"
)

const (
	logginFormat = "%s - \"%s %s %s\" %s\n"
)

func MiddlewareHandler(s *server.Server, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.FixReferer(r)
		ip := utils.RealIP(r)
		log.Printf(logginFormat, ip, r.Method, r.URL.Path, r.Proto, r.UserAgent())
		r.Header.Set("User-Agent", utils.RandomUserAgent())
		h.ServeHTTP(w, r)
	})
}
