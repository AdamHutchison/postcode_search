package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/postcodes", func(w http.ResponseWriter, r *http.Request) {

		postcode := r.URL.Query().Get("postcode")

		search := PostcodeSearch{}

		results := search.Search(postcode)

		io.WriteString(w, results)
	})	

	fmt.Println("Serving app on internal port 80")

	log.Fatal(http.ListenAndServe(":80", nil))
}
