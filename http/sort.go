package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"log"
	"time"
)

func sort() ([]byte, error) {
	var method string = "GET"
	var url string = "http://www.codespub.com/"

	// create a request
	req, err := http.NewRequest(method, url, nil)
	//req.Close = true
	if err != nil {
		return nil, err
	}

	// send JSON to firebase
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad HTTP Response: %v", resp.Status)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	for {
		go func() {
			s, err := sort()
			log.Print(string(s), err)

		}()
		time.Sleep(1000 * 1000 * time.Nanosecond)
	}
}
