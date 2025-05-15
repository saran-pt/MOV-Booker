package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	fmt.Println("Server is running...");
	
	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe(":8080", nil));
}
