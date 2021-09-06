package main

import (
	"log"
	"net/http"
)

func handleStatic() {
	fs := http.FileServer(http.Dir("../client/static"))
	http.Handle("/", fs)
}

func main() {
	handleStatic()

	log.Println("listening on port :3000")
	
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}