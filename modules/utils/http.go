package utils

import (
	"log"
	"net/http"
	"net/url"
)

func HttpGet(rawURL string) *http.Response {
	URL, err := url.Parse(rawURL)
	if err != nil {
		log.Printf("httpGet url.Parse error %v", err)
		return nil
	}

	req := &http.Request{
		Header: baseHeader(),
		Method: "GET",
		URL:    URL,
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
