package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func GetAlbyPayInfo(lnwallet string, sats int) *AlbyPayInfo {
	strs := strings.Split(lnwallet, "@")
	if len(strs) < 2 {
		return nil
	}

	lud06Link := fmt.Sprintf("https://%s/lnurlp/%s/callback?amount=%d000", strs[0], strs[1], sats)
	response := httpGet(lud06Link)
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Printf("GetAlbyPayInfo io.ReadAll error %v", err)
			return nil
		}

		log.Println(string(body))

		payInfo := &AlbyPayInfo{}
		err = json.Unmarshal(body, payInfo)
		if err != nil {
			log.Printf("GetAlbyPayInfo json.Unmarshal error %v", err)
			return nil
		}

		return payInfo
	}

	return nil
}

// The json will like:
// {"status":"OK","successAction":{"tag":"message","message":"Thanks, sats received!"},"verify":"https://getalby.com/lnurlp/haorendashu/verify/zPJFUSpqVXxJ4jw9j8HE3XW8","routes":[],"pr":"lnbc1110n1pjc8af0pp50pxpcuhxswpateh325dj7zx4w0zhtuy3mf5ll45n5w4wz3sqea9qhp5dmegc2zn26n7pnctgh5t9v2ky8ueulcqhvn48nealc8c2eye2frqcqzzsxqyz5vqsp54m437lyceqnpn0m9wfclhcdx0ldztavsj5d2w09mq37hq3j64nrs9qyyssqtt9frxmdce67qcghhhhu70lrxx8s0eh4zfzvss2mu4hyq8w53mqynxsdxgrcal0dqnwwlpxprfyv4kkv7m67rk3pwulnrpk7lzxmc3qpecnd82"}

type AlbyPayInfo struct {
	Status string `json:"status"`
	Verify string `json:"verify"`
	Pr     string `json:"pr"`
}
