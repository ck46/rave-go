package rave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (raveClient *RaveClient) RequestStandardPayment() *RaveResponse {
	reqBytes, _ := json.Marshal(raveClient.Request)
	reqBody := bytes.NewReader(reqBytes)
	req, _ := http.NewRequest(
		"POST",
		raveClient.PaymentUrl,
		reqBody,
	)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", raveClient.SecretKey))

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()

	if err != nil {
		log.Fatal("Error:", err)
	}
	data, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(data, &raveClient.Response)

	return &raveClient.Response
}
