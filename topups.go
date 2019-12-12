package celcoin

import (
	"encoding/json"
	"fmt"
	"net/url"
)

//TopupsRequest - Modelo para solicitar recarga
type TopupsRequest struct {
	TopupData  TopupData  `json:"topupData"`
	CpfCnpj    string     `json:"cpfCnpj"`
	SignerCode string     `json:"signerCode"`
	ProviderID int32      `json:"providerId"`
	Phone      TopupPhone `json:"phone"`
}

//TopupData - Modelo para envio do valor
type TopupData struct {
	Value float32 `json:"value"`
}

//TopupPhone - Modelo telefone do cliente
type TopupPhone struct {
	StateCode   string `json:"stateCode"`
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}

//TopupsResponse - Modelo de resposta
type TopupsResponse struct {
	Authentication    int32                 `json:"authentication"`
	AuthenticationAPI int32                 `json:"authenticationAPI"`
	Receipt           TopupsReceiptResponse `json:"receipt"`
	SettleDate        string                `json:"settleDate"`
	CreateDate        string                `json:"createDate"`
	TransactionID     int32                 `json:"transactionId"`
	Urlreceipt        string                `json:"Urlreceipt"`
	ErrorCode         string                `json:"errorCode"`
	Message           string                `json:"message"`
	Status            int                   `json:"status"`
	Receiptformatted  string                `json:"receiptformatted"`
}

//TopupsReceiptResponse - Modelo Dados Recebedor
type TopupsReceiptResponse struct {
	ReceiptData      string `json:"receiptData"`
	Receiptformatted string `json:"receiptformatted"`
}

//TopupProvider - Modelo da entidade do provedor
type TopupProvider struct {
	Category   int32   `json:"category"`
	Name       string  `json:"name"`
	ProviderID int32   `json:"providerId"`
	MaxValue   float32 `json:"maxValue"`
	MinValue   float32 `json:"minValue"`
}

//TopupProviderValue - Modelo  da entidade de valor
type TopupProviderValue struct {
	Code        int     `json:"code"`
	Cost        int     `json:"cost"`
	Detail      string  `json:"detail"`
	ProductName string  `json:"productName"`
	CheckSum    int     `json:"checkSum"`
	DueProduct  int     `json:"dueProduct"`
	ValueBonus  float32 `json:"valueBonus"`
	MaxValue    float32 `json:"maxValue"`
	MinValue    float32 `json:"minValue"`
}

//TopupProviderResponse - Modelo de resposta da lista do Provider
type TopupProviderResponse struct {
	Providers []TopupProvider `json:"providers"`
	ErrorCode string          `json:"errorCode"`
	Message   string          `json:"message"`
	Status    int             `json:"status"`
}

//TopupProviderValueResponse - Modelo de resposta da lista do values do provider
type TopupProviderValueResponse struct {
	Providers []TopupProviderValue `json:"providers"`
	ErrorCode string               `json:"errorCode"`
	Message   string               `json:"message"`
	Status    int                  `json:"status"`
}

//CreateTopups - Criar um recarga
func (celcoin *CelcoinClient) CreateTopups(req TopupsRequest) (*TopupsResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *TopupsResponse
	err, errAPI := celcoin.Request("POST", "transactions/topups", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//ListTopupsProvider - listar provider
func (celcoin *CelcoinClient) ListTopupsProvider(StateCode, Type, Category string) (*TopupProviderResponse, *Error, error) {
	params := url.Values{}
	params.Add("stateCode", StateCode)
	params.Add("type", Type)
	params.Add("category", Category)
	var response *TopupProviderResponse
	err, errAPI := celcoin.Request("GET", fmt.Sprintf("transactions/topups/providers?%s", params.Encode()), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//ListTopupsProviderValues - listar provider
func (celcoin *CelcoinClient) ListTopupsProviderValues(StateCode, ProviderID string) (*TopupProviderValueResponse, *Error, error) {
	params := url.Values{}
	params.Add("stateCode", StateCode)
	params.Add("providerId", ProviderID)
	var response *TopupProviderValueResponse
	err, errAPI := celcoin.Request("GET", fmt.Sprintf("transactions/topups/provider-values?%s", params.Encode()), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
