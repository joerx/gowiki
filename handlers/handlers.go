package handlers

import (
	"net/http"

	"github.com/joerx/gowiki/wiki"
)

// View handler renders the current page
var View = mkHandler(viewPage)

// Edit handler renders a form to edit/create the page given by title
var Edit = mkHandler(editPage)

// Save handler saves the content sent via form to the wiki
var Save = mkHandler(savePage)

func viewPage(w http.ResponseWriter, r *http.Request, title string) {
	p, err := wiki.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editPage(w http.ResponseWriter, r *http.Request, title string) {
	p, err := wiki.LoadPage(title)
	if err != nil {
		p = &wiki.Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func savePage(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &wiki.Page{Title: title, Body: []byte(body)}
	if err := p.Save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// Index handler lists all pages in the wiki
func Index(w http.ResponseWriter, r *http.Request) {
	pages, err := wiki.ListPages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderTemplate(w, "index", pages)
}

// FrontPage handler redirects the user to the page called "FrontPage"
func FrontPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

// mkHandler wraps app specific handler function into a http.HandlerFunc and adds some validation logic
func mkHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
		}
		fn(w, r, m[2])
	}
}
