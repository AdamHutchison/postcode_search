package main

import (
	"io/ioutil"
	"net/http"
)

type PostcodeSearch struct {

}

func (p *PostcodeSearch) Search(postcode string) string {
	client := &http.Client{}

	resp, _ := client.Get("https://api.postcodes.io/postcodes/" + postcode)

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}