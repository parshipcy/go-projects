package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST req success\n")

	name := r.FormValue("name") //pulls the value of an HTML input whose name="key" from the HTTP request.
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name) // %s is replaced by the value of name
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) { //Declares an HTTP handler function. w is used to write the HTTP response back to the client. r contains the incoming HTTP request and all its data.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound) //Sets the HTTP status code to 404 //Writes the error message to the response body //Sets the Content-Type to text/plain; charset=utf-8
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "hello!") //Writes formatted text to a specific writer
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at 8080")
	// http.ListenAndServe(":8080", nil) starts the server on port 8080.
	// Passing nil means Go uses the default HTTP multiplexer where all routes were registered.
	// If the server fails to start, the program exits with a fatal log.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}


/*

Example 1: WITHOUT ParseForm() (problem case)
func formHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name=%s Address=%s", name, address)
}

What happens:
The POST body is not parsed
r.Form is empty
FormValue() may return empty strings

Output in browser
Name= Address=
This is unreliable behavior.


Example 2: WITH ParseForm() (correct)
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name=%s Address=%s", name, address)
}

What happens internally
ParseForm() decodes name=Alice&address=Paris

Go builds:
map[name:[Alice] address:[Paris]]

Output in browser
Name=Alice Address=Paris

*/