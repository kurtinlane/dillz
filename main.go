package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// Set initial variables
	accountSid := "ACcd11ee8452c1eacb8a16599b9a3662da"
	authToken := "bb6b4d096ed678f6263cbdcabd535094"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/SMS/Messages.json"

	phoneNumbers := []string{"+18019716820", "+14155318437", "+15412078581"}

	for index, each := range phoneNumbers {
		fmt.Printf("Phone number [%d] is [%s]\n", index, each)
		// Build out the data for our message
		v := url.Values{}
		// v.Set("To", "+18019716820")
		// v.Set("To", "+14155318437")
		v.Set("To", each)
		v.Set("From", "+18012141029")
		v.Set("Body", "Sup Birch, come get your daily dilllzzz")
		rb := *strings.NewReader(v.Encode())

		// Create client
		client := &http.Client{}

		req, _ := http.NewRequest("POST", urlStr, &rb)
		req.SetBasicAuth(accountSid, authToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		// Make request
		resp, _ := client.Do(req)
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(bodyBytes, &data)
		fmt.Println(data)
		fmt.Println(resp.Status)
	}
}
