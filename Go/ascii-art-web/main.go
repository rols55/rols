package main

import (
	"ascii-art-web/asciiart"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// inits and runs server
func main() {

	mux := http.NewServeMux()

	//lines 23- 24 are for the server to be able to access files inside server's root directory
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./style"))))
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	mux.HandleFunc("/ascii-art", InputHandler)
	mux.HandleFunc("/download", DownloadHandler)
	mux.HandleFunc("/", firstPage)

	fmt.Println("Visit in your browser http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}

func firstPage(w http.ResponseWriter, r *http.Request) {
	// template to present index page
	tmpl := template.Must(template.ParseGlob("./static/index.html"))

	if r.URL.Path != "/" {
		http.Error(w, "Bad request - 404 resource not found.", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, r)

}

func InputHandler(w http.ResponseWriter, r *http.Request) {
	// template to present ASCII art
	tmpl := template.Must(template.ParseGlob("./static/output.html"))

	//Data from the web form located in the index.html
	inputText := r.FormValue("inputText")
	banner := r.FormValue("banner")

	//To controll input text characters
	if asciiart.ParseInput(inputText) != "" {
		http.Error(w, "Bad Request - Invalid character(s) detected", http.StatusBadRequest)
		return
	}

	// In case of correct action we execute template and display ascii art on it
	if r.URL.Path == "/ascii-art" {

		//To controll methods
		switch r.Method {
		case "GET":
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		case "POST":
			AsciiWeb := asciiart.AsciiFormat(inputText, banner)
			tmpl.Execute(w, AsciiWeb)
		}

		// Error handling with wrong path
	} else {
		http.Error(w, "Bad request - 404 resource not found.", http.StatusNotFound)
		return
	}

}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	// file donwloads are handled by POST method
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/download" {
		http.Error(w, "Bad request - 404 resource not found.", http.StatusNotFound)
		return
	}

	file, err := os.Open("./download/download.txt")
	if err != nil {
		fmt.Println("error reading file")
		return
	}
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("error reading file stat")
		return
	}
	defer file.Close()
	w.Header().Set("Content-Disposition", "attchment; filename="+"download.txt")
	w.Header().Set("Conetnt-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(int(fi.Size())))
	io.Copy(w, file)
}
