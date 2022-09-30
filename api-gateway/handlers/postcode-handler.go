package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func PostcodeHandler(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Routing request to postcode_service")

		postcode := r.URL.Query().Get("postcode")

		client := &http.Client{}

		resp, _ := client.Get("http://postcode_service/postcodes?postcode=" + postcode)

		body, _ := ioutil.ReadAll(resp.Body)

		io.WriteString(w, string(body))
}