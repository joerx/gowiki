package main

import (
	"log"
	"net/http"

	"github.com/joerx/gowiki/handlers"
)

func main() {
	http.HandleFunc("/view/", handlers.View)
	http.HandleFunc("/edit/", handlers.Edit)
	http.HandleFunc("/save/", handlers.Save)
	http.HandleFunc("/index/", handlers.Index)
	http.HandleFunc("/", handlers.FrontPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
