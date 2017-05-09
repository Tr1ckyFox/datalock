package handlers

import (
	"net/http"
	"strings"

	"github.com/leominov/datalock/seasonvar"
	"github.com/leominov/datalock/utils"
)

func JavaScriptHandler(s *seasonvar.Seasonvar) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := s.AbsoluteLink(r.URL.RequestURI())
		b, err := utils.HttpGet(url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b = strings.Replace(b, seasonvar.Hostname, r.Header.Get("X-HOSTNAME"), -1)
		b = strings.Replace(b, "swichHDno:", "swichHDdisabled:", -1)
		b = strings.Replace(b, "swichHD:", "swichHDno:", -1)
		w.Write([]byte(b))
	})
}
