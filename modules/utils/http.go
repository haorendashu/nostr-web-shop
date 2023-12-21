package utils

import (
	"log"
	"net/http"
)

func httpGet(url string) *http.Response {
	req := &http.Request{
		Header: baseHeader(),
		Method: "GET",
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("httpGet http.DefaultClient.Do error %v", err)
		return nil
	}

	return response
}

func baseHeader() http.Header {
	return map[string][]string{
		"User-Agent": {"Nostr-Web-Shop"},
	}
}
