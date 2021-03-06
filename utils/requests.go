package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(robots)
}
