// Problem 18
// writing a simple http application

package main

import(
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Server is running")
}

func about_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "My name is Arjun Tiwari and I am making a very basic http request and reponse language using golang")
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}