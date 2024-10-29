package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func main() {
	bizCheckUrl := "https://accounts.binance.com/bapi/accounts/v1/public/account/security/bizCheck"

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	bizCheckRequest := BizCheckRequest{
		Mobile:      "3381020494",
		MobileCode:  "IT",
		CallingCode: "39",
		BizType:     "login",
		SessionId:   "08A201BA3E334EDEABBD144A93806BA28E622295",
	}

	marshal, _ := json.Marshal(bizCheckRequest)
	requestBody := bytes.NewBufferString(string(marshal))

	request, err := http.NewRequest("POST", bizCheckUrl, requestBody)
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	if resp.StatusCode != 200 {
		panic(resp)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var bizCheckResponse BizCheckResponse
	err = json.Unmarshal(respBody, &bizCheckResponse)
	if err != nil {
		panic(err)
	}

	if bizCheckResponse.Data.Valid {
		println("Valid")
	} else {
		println("Invalid")
	}
}

type BizCheckRequest struct {
	Mobile      string `json:"mobile"`
	MobileCode  string `json:"mobileCode"`
	CallingCode string `json:"callingCode"`
	BizType     string `json:"bizType"`
	SessionId   string `json:"sessionId"`
}

type BizCheckResponse struct {
	Code          string      `json:"code"`
	Message       interface{} `json:"message"`
	MessageDetail interface{} `json:"messageDetail"`
	Data          struct {
		Valid     bool `json:"valid"`
		Pless     bool `json:"pless"`
		ForceFido bool `json:"forceFido"`
	} `json:"data"`
	Success bool `json:"success"`
}
