package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	h "github.com/AdamHutchison/api-gateway/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "API Gateway Hit")
	})

	http.HandleFunc("/postcodes", h.PostcodeHandler)
	http.HandleFunc("/auth", h.AuthHandler)

	fmt.Println("Serving API Gateway on port 8000")

	log.Fatal(http.ListenAndServe(":8000", nil))
}