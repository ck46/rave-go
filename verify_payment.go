package rave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (raveClient *RaveClient) VerifyTransaction(trn_id string) *RaveResponse {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf(VerificationUrl, trn_id),
		bytes.NewReader([]byte("")),
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
