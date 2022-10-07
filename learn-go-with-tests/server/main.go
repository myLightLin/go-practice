package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5100", server); err != nil {
		log.Fatalf("could not listen on port 5100 %v", err)
	}
}
