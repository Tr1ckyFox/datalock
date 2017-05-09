package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/leominov/datalock/seasonvar"
	"github.com/leominov/datalock/utils"
)

func MeHandler(s *seasonvar.Seasonvar) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := utils.RealIP(r)
		switch r.Method {
		case "POST":
			decoder := json.NewDecoder(r.Body)
			var u *seasonvar.User
			err := decoder.Decode(&u)
			if err != nil {
				http.Error(w, "Could not decode body", http.StatusInternalServerError)
				return
			}
			u.UserAgent = r.UserAgent()
			u.IP = ip
			u.SecureMark = utils.CleanText(u.SecureMark)
			if err := s.SetUser(u); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			encoder := json.NewEncoder(w)
			if err := encoder.Encode(u); err != nil {
				http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
				return
			}
		default:
			u, err := s.GetUser(ip)
			if err != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			encoder := json.NewEncoder(w)
			if err := encoder.Encode(u); err != nil {
				http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
				return
			}
		}
	})
}
