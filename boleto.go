package celcoin

import (
	"encoding/json"
)

//BankslipRequest - Modelo para criação de um boleto
type BankslipRequest struct {
	Document         string                        `json:"document"`
	IPAddress        string                        `json:"ipAddress"`
	HashIntegration  string                        `json:"hashIntegration"`
	ExternalTerminal string                        `json:"externalTerminal"`
	Value            float32                       `json:"value"`
	BranchCode       int                           `json:"branchCode"`
	InstitutionCode  int                           `json:"institutionCode"`
	WalletCode       int                           `json:"walletCode"`
	Covenant         string                        `json:"covenant"`
	AccountCode      int                           `json:"accountCode"`
	DigitCode        string                        `json:"digitCode"`
	PayerAddress     BankslipPayerAddressRequest   `json:"payerAddress"`
	Instructions     []BankslipInstructionsRequest `json:"instructions"`
	AccountType      int                           `json:"accountType"`
	PayerName        string                        `json:"payerName"`
	ResponseType     int                           `json:"responseType"`
	DueDate          string                        `json:"dueDate"`
}

//BankslipPayerAddressRequest - Modelo de endereço do cliente
type BankslipPayerAddressRequest struct {
	ZipCode    string `json:"zipCode"`
	City       string `json:"city"`
	Additional string `json:"additional"`
	Street     string `json:"street"`
	State      string `json:"state"`
}

//BankslipInstructionsRequest - Modelo para instruições do boleto
type BankslipInstructionsRequest struct {
	Text string `json:"text"`
}

//Bankslip - Criar boleto
func (celcoin *CelcoinClient) Bankslip(req BankslipRequest) (*BankslipRequest, *Error, error) {
	data, _ := json.Marshal(req)
	var response *BankslipRequest
	err, errAPI := celcoin.Request("POST", "transactions/bankslip", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
