package main

import (
	"log"
	"net/http"

	"github.com/doneva593/wifisystem/doc"
)

func main() {
	s, err := doc.NewServer("doc")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("serving on 0.0.0.0:" + "9999")
	err = http.ListenAndServe(":9999", s)
	if err != nil {
		log.Fatal(err)
	}
}
