package celcoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type CelcoinClient struct {
	client      *http.Client
	Username    string
	Password    string
	Env         string
	ExpiteTime  int64
	Token       string
	Openfinance bool
}

type Error struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Body      string `json:"body"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	RokenType   string `json:"token_type"`
}

func NewCelcoinClient(username, password, env string, openfinance bool) *CelcoinClient {
	return &CelcoinClient{
		client:      &http.Client{Timeout: 60 * time.Second},
		Username:    username,
		Password:    password,
		Env:         env,
		Openfinance: openfinance,
	}
}

func (celcoin *CelcoinClient) Request(method, action string, body []byte, out interface{}) (error, *Error) {
	if celcoin.client == nil {
		celcoin.client = &http.Client{Timeout: 60 * time.Second}
	}
	url := celcoin.devProd()
	if celcoin.Openfinance {
		url = celcoin.openfinanceUrl()
	}

	endpoint := fmt.Sprintf("%s/%s", url, action)
	log.Printf("endpoint %s \n", endpoint)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err, nil
	}

	_, err = celcoin.RequestToken()
	if err != nil {
		return err, nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", celcoin.Token))
	res, err := celcoin.client.Do(req)
	if err != nil {
		return err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	log.Printf("bodyResponse %s \n", bodyResponse)
	log.Printf("res.StatusCode  %d \n", res.StatusCode)
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
		return "https://sandbox-apicorp.celcoin.com.br/v5"
	}
	return "https://sandbox-apicorp.celcoin.com.br/v5"
}

func (CelcoinClient *CelcoinClient) TokenUri() string {
	if CelcoinClient.Env == "develop" {
		return "https://sandbox-apicorp.celcoin.com.br/v5"
	}
	return "https://sandbox-apicorp.celcoin.com.br/v5"
}

func (CelcoinClient *CelcoinClient) openfinanceUrl() string {
	if CelcoinClient.Env == "develop" {
		return "https://sandbox.openfinance.celcoin.com.br"
	}
	return "https://sandbox.openfinance.celcoin.com.br"
}

func (celcoin *CelcoinClient) RequestToken() (*TokenResponse, error) {
	var tokenResponse TokenResponse
	if celcoin.client == nil {
		celcoin.client = &http.Client{Timeout: 60 * time.Second}
	}
	data := url.Values{
		"client_id":     {celcoin.Username},
		"grant_type":    {"client_credentials"},
		"client_secret": {celcoin.Password},
	}
	endpoint := fmt.Sprintf("%s/%s", celcoin.TokenUri(), "token")
	res, err := http.PostForm(endpoint, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyResponse, &tokenResponse)
	if err != nil {
		return nil, err
	}
	celcoin.Token = tokenResponse.AccessToken
	return &tokenResponse, nil
}
