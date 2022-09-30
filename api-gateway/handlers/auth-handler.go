package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Routing request to Auth Server")

		client := &http.Client{}

		resp, _ := client.Get("http://auth_server")

		body, _ := ioutil.ReadAll(resp.Body)

		io.WriteString(w, string(body))
}