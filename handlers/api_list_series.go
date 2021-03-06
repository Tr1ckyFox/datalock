package handlers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/leominov/datalock/server"
	"github.com/leominov/datalock/utils"
	xmlpath "gopkg.in/xmlpath.v2"
)

var (
	xpathSeries         = xmlpath.MustCompile(`.//a`)
	xpathSeriesHref     = xmlpath.MustCompile(`.//@href`)
	xpathSeriesName     = xmlpath.MustCompile(`.//div[contains(@class, 'rside-t')]`)
	xpathSeriesNumber   = xmlpath.MustCompile(`.//div[contains(@class, 'rside-ss')]`)
	reInsideWhitespaces = regexp.MustCompile(`[\s\p{Zs}]{2,}`)
)

type Series struct {
	Name    string `json:"name"`
	Link    string `json:"link"`
	Comment string `json:"comment"`
}

func getSeriesListFromBody(body []byte) ([]Series, error) {
	series := []Series{}
	root, err := xmlpath.ParseHTML(bytes.NewReader(body))
	if err != nil {
		return series, err
	}
	iter := xpathSeries.Iter(root)
	for iter.Next() {
		node := iter.Node()
		link, ok := xpathSeriesHref.String(node)
		if !ok {
			continue
		}
		name, ok := xpathSeriesName.String(node)
		if !ok {
			continue
		}
		number, ok := xpathSeriesNumber.String(node)
		if !ok {
			continue
		}
		number = strings.TrimSpace(number)
		number = reInsideWhitespaces.ReplaceAllString(number, " ")
		series = append(series, Series{
			Link:    link,
			Name:    name,
			Comment: number,
		})
	}
	return series, nil
}

func ApiListSeriesHandler(s *server.Server, listType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := s.AbsoluteLink(fmt.Sprintf("/ajax.php?mode=%s", listType))
		b, err := utils.HttpGetRaw(url, map[string]string{
			"X-Requested-With": "XMLHttpRequest",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		series, err := getSeriesListFromBody(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		switch r.URL.Query().Get("_format") {
		case "xml":
			w.Header().Set("Contern-Type", "application/xml;charset=utf-8")
			encoder := xml.NewEncoder(w)
			encoder.Encode(series)
		default:
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			encoder := json.NewEncoder(w)
			encoder.Encode(series)
		}
	})
}
