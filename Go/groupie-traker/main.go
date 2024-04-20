package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
TODO fix animations
TODO error handling
TODO Make a real template for pages
*/

func main() {
	//two handles give server access to folders from where to fetch relevant files
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.HandleFunc("/", index)

	fmt.Println("Starting server, please await further instructions")

	if parseData() {
		fmt.Printf("\nVisit in your browser http://localhost:8080\nPress CTRL + C in terminal to stop the server\n")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("./templates/testPage.html"))
	if r.URL.Path != "/" {
		http.Error(w, "Bad request - 404 resource not found.", http.StatusNotFound)
	}
	//parses artist data from API and saves info into a struct
	//executes template with struct containing artist data
	tmpl.Execute(w, CleanedUpArtists)
}
