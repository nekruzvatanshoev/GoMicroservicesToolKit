package web

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequest(url string) ([]byte, error) {

	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	return body, err
}


