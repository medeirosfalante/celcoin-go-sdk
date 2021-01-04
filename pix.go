package celcoin

import "encoding/json"

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
