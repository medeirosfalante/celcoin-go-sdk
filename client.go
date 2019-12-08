package celcoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CelcoinClient struct {
	client   *http.Client
	Username string
	Password string
	Env      string
}

type Error struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Body      string `json:"body"`
}

func NewCelcoinClient(username, password, env string) *CelcoinClient {
	return &CelcoinClient{
		client:   &http.Client{Timeout: 60 * time.Second},
		Username: username,
		Password: password,
		Env:      env,
	}
}

func (celcoin *CelcoinClient) Request(method, action string, body []byte, out interface{}) (error, *Error) {
	if celcoin.client == nil {
		celcoin.client = &http.Client{Timeout: 60 * time.Second}
	}
	endpoint := fmt.Sprintf("%s/%s", celcoin.devProd(), action)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err, nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(celcoin.Username, celcoin.Password)
	res, err := celcoin.client.Do(req)
	if err != nil {
		return err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 201 {

		var errAPI Error
		err = json.Unmarshal(bodyResponse, &errAPI)
		if err != nil {
			return err, nil
		}
		errAPI.Body = string(bodyResponse)
		return nil, &errAPI
	}
	err = json.Unmarshal(bodyResponse, out)
	if err != nil {
		return err, nil
	}
	return nil, nil
}

func (CelcoinClient *CelcoinClient) devProd() string {
	if CelcoinClient.Env == "develop" {
		return "https://sandbox-apicorp.celcoin.com.br/v4"
	}
	return "https://sandbox-apicorp.celcoin.com.br/v4"
}
