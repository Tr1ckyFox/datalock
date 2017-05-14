package handlers

import (
	"html/template"
	"io/ioutil"
	"path"
	"strings"

	"github.com/leominov/datalock/server"
)

var (
	Templates *template.Template
)

func ParseTemplates(config *server.Config) error {
	var err error
	var allFiles []string
	files, err := ioutil.ReadDir(config.TemplatesDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".tmpl") {
			allFiles = append(allFiles, path.Join(config.TemplatesDir, filename))
		}
	}
	Templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		return err
	}
	return nil
}
