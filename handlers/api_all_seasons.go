package handlers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"

	"gopkg.in/xmlpath.v2"

	"github.com/leominov/datalock/server"
	"github.com/leominov/datalock/utils"
)

var (
	xpathSeasons = xmlpath.MustCompile(`.//ul[contains(@class,'tabs-result')]/li[contains(@class,'act')]/h2`)
	xpathLink    = xmlpath.MustCompile(`.//@href`)
)

type Season struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func ApiAllSeasonsHandler(s *server.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if len(url) == 0 {
			http.Error(w, "Incorrect request", http.StatusBadRequest)
			return
		}
		b, err := utils.HttpGetRaw(s.AbsoluteLink(url), map[string]string{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		root, err := xmlpath.ParseHTML(bytes.NewReader(b))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		seasons := []Season{}
		iter := xpathSeasons.Iter(root)
		for iter.Next() {
			node := iter.Node()
			title := strings.TrimSpace(node.String())
			link, ok := xpathLink.String(iter.Node())
			if !ok {
				continue
			}
			seasons = append(seasons, Season{
				Title: utils.StandardizeSpaces(title),
				Link:  strings.TrimSpace(link),
			})
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		switch r.URL.Query().Get("_format") {
		case "xml":
			w.Header().Set("Contern-Type", "application/xml;charset=utf-8")
			encoder := xml.NewEncoder(w)
			encoder.Encode(seasons)
		default:
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			encoder := json.NewEncoder(w)
			encoder.Encode(seasons)
		}
	})
}
