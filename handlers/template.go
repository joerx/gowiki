package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const templateDir = "templates"

var templates *template.Template

var validPath = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]+)$")

type layoutContent struct {
	Content template.HTML
}

// Init reads and caches templates, if that fails exits the program
func init() {
	files, err := listTemplates()
	if err != nil {
		log.Fatal(err)
		return
	}
	templates = template.Must(template.ParseFiles(files...))
}

func listTemplates() ([]string, error) {
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, len(files))

	for _, f := range files {
		name := f.Name()
		if f.IsDir() || strings.HasPrefix(name, ".") {
			continue
		}
		res = append(res, templateDir+"/"+name)
	}

	return res, nil
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	// just buffering the entire page into memory is likely not a great idea
	// might want to something dynamic called from within the layout that writes
	// the content into the existing io.writer
	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, fmt.Sprintf("%s.html", name), data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// content was sanitized before, so we can assume it is save now
	td := layoutContent{Content: template.HTML(buf.String())}
	if err := templates.ExecuteTemplate(w, "layout.html", td); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
