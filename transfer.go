package celcoin

import (
	"encoding/json"
	"fmt"
)

type BankTransfer struct {
	Document         string  `json:"document"`
	ExternalTerminal string  `json:"externalTerminal"`
	ExternalNSU      int     `json:"externalNSU"`
	AccountCode      int     `json:"accountCode"`
	DigitCode        string  `json:"digitCode"`
	BranchCode       int     `json:"branchCode"`
	InstitutionCode  int     `json:"institutionCode"`
	Name             string  `json:"name"`
	Value            float32 `json:"value"`
	BankAccountType  int     `json:"bankAccountType"`
}

type BankTransferResponse struct {
	CreateDate             string                              `json:"createDate"`
	TransactionID          int32                               `json:"transactionId"`
	Urlreceipt             string                              `json:"Urlreceipt"`
	ErrorCode              string                              `json:"errorCode"`
	Message                string                              `json:"message"`
	Status                 int32                               `json:"status"`
	NextSettle             bool                                `json:"nextSettle"`
	DateNextLiquidation    string                              `json:"DateNextLiquidation"`
	Value                  float32                             `json:"value"`
	DestinationAccountData *BankTransferDestinationAccountData `json:"destinationAccountData"`
	AuthenticationAPI      *BankTransferAuthenticationAPI      `json:"authenticationAPI"`
}

type BankTransferAuthenticationAPI struct {
	Bloco1        string `json:"Bloco1"`
	Bloco2        string `json:"Bloco2"`
	BlocoCompleto string `json:"BlocoCompleto"`
}

type BankTransferDestinationAccountData struct {
	Agency               int    `json:"agency"`
	InstitutionCode      int    `json:"institutionCode"`
	Account              int    `json:"account"`
	AccountVerifierDigit string `json:"accountVerifierDigit"`
	Document             string `json:"document"`
	InstitutionName      string `json:"institutionName"`
	FullName             string `json:"fullName"`
	BankAccountType      int    `json:"bankAccountType"`
	DocumentType         string `json:"documentType"`
}

//BankTransfer - Criar transferência bancária
func (celcoin *CelcoinClient) BankTransfer(req BankTransfer) (*BankTransferResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *BankTransferResponse
	err, errAPI := celcoin.Request("POST", "transactions/banktransfer", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//BankTransferStatus - status transaction
func (celcoin *CelcoinClient) BankTransferStatus(transactionId int) (*BankTransferResponse, *Error, error) {
	var response *BankTransferResponse
	err, errAPI := celcoin.Request("GET", fmt.Sprintf("transactions/banktransfer/status-transfer/%d", transactionId), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
