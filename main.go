package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No PORT env specified!")
	}
	log.Println("PORT: " + port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
