package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type IO struct {
	Input  string
	Output string
}

// Server runs on http://localhost:8080/

func main() {

	http.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("templates"))))

	http.HandleFunc("/", asciiFormHandler) // Handles /ascii-art

	fmt.Printf("Starting server at port 8080, access the page with 'localhost:8080' in a browser\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // Listens on port 8080
		log.Fatal(err)
	}
}

// Handles POST from /form
func asciiFormHandler(w http.ResponseWriter, r *http.Request) {

	var io IO
	whtml, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "404 - Resource not found", http.StatusNotFound)
	}

	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		http.Error(w, "404 - Page not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "405 - Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		whtml.Execute(w, io)
	} else {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}
		font := r.FormValue("fonts")
		io.Input = r.PostFormValue("input")

		if !checkFalseCharacters(io.Input) {
			http.Error(w, "400 - Bad request", http.StatusBadRequest)
			return
		}

		if io.Input == "" {
			whtml.Execute(w, io)
			return
		}

		ascii := getAscii(io.Input, font)

		whtml, _ = template.ParseFiles("templates/index.html")
		w.WriteHeader(http.StatusOK)
		io.Output = strArrayToString(ascii)
		err = whtml.Execute(w, io)
	}
	// Handles security, checks if url path is /ascii-art, else returns error 404
	/*if r.URL.Path != "/ascii-art" {
		http.Error(w, "404 - Page not found.", http.StatusNotFound)
		return
	} */

	// Security, checks if method is POST, else returns error 405

	// Check for any internal errors

	// Check for non-ascii characters, if found returns error code 400

	// Splits input by newlines
	/*	splitNewlines := strings.Split(input, "\n")

		// Reads splitNewLines by row
		for _, row := range splitNewlines {

			// Transforms string input into ascii output
			asciiOutput := getAscii(row, font)

			// Prints ascii row by row, since asciiOutput is a string array
			for _, asciiRow := range asciiOutput {
				fmt.Fprintf(w, "%s", asciiRow)
				fmt.Fprint(w, "\n")
			}
		}*/

}

func checkFalseCharacters(input string) bool {
	for _, char := range input {
		if !(char >= 32 && char <= 126 || (char == 13 || char == 10)) {
			return false
		}
	}
	return true
}
