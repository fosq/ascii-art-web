package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Server runs on http://localhost:8080/

func main() {

	//fileServer := http.FileServer(http.Dir("./templates")) // Creates file server object FileServer
	//http.Handle("/", fileServer)                           // Handles "/" path to static directory
	http.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("templates"))))

	http.HandleFunc("/", asciiFormHandler) // Handles /ascii-art

	fmt.Printf("Starting server at port 8080, access the page with 'localhost:8080' in a browser\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // Listens on port 8080
		log.Fatal(err)
	}
}

// Handles POST from /form
func asciiFormHandler(w http.ResponseWriter, r *http.Request) {

	var input string

	whtml, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "404 - Resource not found", http.StatusNotFound)
	}

	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		whtml.Execute(w, input)
	}
	// Handles security, checks if url path is /ascii-art, else returns error 404
	/*if r.URL.Path != "/ascii-art" {
		http.Error(w, "404 - Page not found.", http.StatusNotFound)
		return
	} */

	// Security, checks if method is POST, else returns error 405
	/*if r.Method != "POST" {
		http.Error(w, "405 - Method is not supported.", http.StatusMethodNotAllowed)
		return
	}*/

	// Check for any internal errors
	if err := r.ParseForm(); err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}

	input = r.PostFormValue("input")
	//font := r.FormValue("fonts")
	input = removeQuotations(input)

	if input == "" {
		w.WriteHeader(http.StatusOK)
		whtml.Execute(w, "Hello execute!")
		return
	}

	// Check for non-ascii characters, if found returns error code 400
	if !checkFalseCharacters(input) {
		http.Error(w, "400 - Bad request", http.StatusBadRequest)
		return
	}

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
	whtml.Execute(w, "hello")
}

// Removes " from input, if found at the beginning and end of string
func removeQuotations(input string) string {
	var newInput string
	lenInput := len(input) - 1
	if lenInput > 1 {
		if input[0] == '"' && input[lenInput] == '"' {
			newInput = input[1:lenInput]
			return newInput
		}
	}
	return input
}

func checkFalseCharacters(input string) bool {
	for _, char := range input {
		if !(char >= 32 && char <= 126 || (char == 13 || char == 10)) {
			return false
		}
	}
	return true
}
