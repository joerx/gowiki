package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const storagePath = "pages"

// Page represents a page in the wiki
type Page struct {
	Title string
	Body  []byte
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("I love golang")}
	p1.save()
	p2, err := loadPage("TestPage")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(p2.Body))
}

func (p *Page) save() error {
	filename := storagePath + "/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := storagePath + "/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
