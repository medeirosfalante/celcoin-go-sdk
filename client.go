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
	client     *http.Client
	Username   string
	Password   string
	Env        string
	Version    string
	ExpiteTime int64
	Token      string
}

type Error struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Body      string `json:"body"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	RokenType   string `json:"token_type"`
}

func NewCelcoinClient(username, password, env, version string) *CelcoinClient {
	return &CelcoinClient{
		client:   &http.Client{Timeout: 60 * time.Second},
		Username: username,
		Password: password,
		Env:      env,
		Version:  version,
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

	log.Printf("\n\n endpoint %s\n\n", string(endpoint))

	req.Header.Add("Content-Type", "application/json")

	switch celcoin.Version {
	case "v4":
		req.SetBasicAuth(celcoin.Username, celcoin.Password)
	case "v5":
		_, err := celcoin.RequestToken()
		if err != nil {
			return err, nil
		}

		log.Printf("\n\n token %s\n\n", fmt.Sprintf("Bearer %s", celcoin.Token))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", celcoin.Token))
	}

	res, err := celcoin.client.Do(req)
	if err != nil {
		return err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)

	log.Printf("\n\n bodyResponse %s\n\n", string(bodyResponse))
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
		return fmt.Sprintf("https://sandbox-apicorp.celcoin.com.br/%s", CelcoinClient.Version)
	}
	return fmt.Sprintf("https://sandbox-apicorp.celcoin.com.br/%s", CelcoinClient.Version)
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
	endpoint := fmt.Sprintf("%s/%s", celcoin.devProd(), "token")
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
