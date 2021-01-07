package celcoin

import (
	"encoding/json"
	"fmt"
)

type StaticBRCodeCreationRequest struct {
	Amount                    float32   `json:"amount"`
	Key                       string    `json:"key"`
	TransactionIdentification string    `json:"transactionIdentification"`
	Merchant                  *Merchant `json:"merchant"`
	AdditionalInformation     string    `json:"additionalInformation"`
}

type Merchant struct {
	PostalCode           string `json:"postalCode"`
	City                 string `json:"city"`
	MerchantCategoryCode int32  `json:"merchantCategoryCode"`
	Name                 string `json:"name"`
}

type StaticBRCodeCreationResponse struct {
	TransactionID             int32  `json:"transactionId"`
	Emvqrcps                  string `json:"emvqrcps"`
	TransactionIdentification string `json:"transactionIdentification"`
}

type DictResponse struct {
	Key        string      `json:"key"`
	KeyType    string      `json:"keyType"`
	Account    *PixAccount `json:"account"`
	Owner      *PixOwner   `json:"owner"`
	Endtoendid string      `json:"endtoendid"`
}

type PixAccount struct {
	OpeningDate   string `json:"openingDate"`
	Participant   string `json:"participant"`
	Branch        string `json:"branch"`
	AccountNumber string `json:"accountNumber"`
	AccountType   string `json:"accountType"`
}

type PixOwner struct {
	TaxIDNumber string `json:"taxIdNumber"`
	Type        string `json:"type"`
	Name        string `json:"name"`
}

//CreatePixBrCodeStatic - criar um brcode stático
func (celcoin *CelcoinClient) CreatePixBrCodeStatic(req StaticBRCodeCreationRequest) (*StaticBRCodeCreationResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *StaticBRCodeCreationResponse
	err, errAPI := celcoin.Request("POST", "pix/v1/brcode/static", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//GetDic - pegar infos do dic
func (celcoin *CelcoinClient) GetDic(key string) (*DictResponse, *Error, error) {
	var response *DictResponse
	err, errAPI := celcoin.Request("POST", fmt.Sprintf("pix/v1/dict/key/%s", key), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
