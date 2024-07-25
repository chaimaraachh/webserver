package main 

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		htttp.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.method != "GET" {
		htttp.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf("Hello")
}
func formHandler(w http.ResponseWriter, r "http.Request"){
	if err := r.Parseform(); err != nil {
	fmt.Fprintf(w, "ParseForm() err: %v", err)
	return
	
	fmt.Fprintf(w, "POST request successful")
	hame := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address -", address)}

func main() {
	fileServer := http.fileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)


	fmt.Printf("Starting server at port 8080")

	if err := http.ListenAndServer(":8080", nil); err != nil {
		log.Fatal(err)
	}


}