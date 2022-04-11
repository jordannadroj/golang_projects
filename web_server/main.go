package main

import (
	"fmt"
	"log"
	"net/http"
)

// w = responseWriter, the response that is returned from the request
//r = request, the request to the server,
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" { // ensure the path is /hello
		http.Error(w, "404 not found", http.StatusNotFound) // send to responsewrite, the error
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound) // send error to w, the error
	}

	fmt.Fprintf(w, "hello!") // Fprintf writes to a responseWriter
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { //r.ParseForm will parse the information passed into the form in /form.html
		fmt.Fprintf(w, "ParseForm() error %v", err)
	}
	fmt.Fprintf(w, "Form submitted\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %v\n", name)
	fmt.Fprintf(w, "Address = %v", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("../web_server")) // .Dir specifies the directory tree of the files to be used

	http.Handle("/", fileServer)          // create a handle to see root page to server
	http.HandleFunc("/form", formHandler) // handle /form route
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
