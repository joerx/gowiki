package wiki

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

const storagePath = "pages"

// Page represents a page in the wiki
type Page struct {
	Title string
	Body  []byte
}

// Save saves a page to disk
func (p *Page) Save() error {
	filename := storagePath + "/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// LoadPage loads a page given by title from disk
func LoadPage(title string) (*Page, error) {
	filename := storagePath + "/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// ListPages returns a list of all wiki pages that currently exist
func ListPages() ([]string, error) {
	files, err := ioutil.ReadDir(storagePath)
	if err != nil {
		return nil, err
	}

	var res []string
	var name string

	for _, f := range files {
		name = f.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		name = strings.TrimSuffix(name, filepath.Ext(name))
		res = append(res, name)
	}

	return res, nil
}
