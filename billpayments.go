package celcoin

import (
	"encoding/json"
	"fmt"
)

//BillPayment - Modelo para criação de um pagamento de conta
type BillPayment struct {
	ExternalNSU            int32    `json:"externalNSU"`
	ExternalTerminal       string   `json:"externalTerminal"`
	Cpfcnpj                string   `json:"cpfCnpj"`
	BillData               BillData `json:"billData"`
	BarCode                BarCode  `json:"barCode"`
	DueDate                string   `json:"dueDate"`
	TransactionIDAuthorize int32    `json:"transactionIdAuthorize"`
}

//BillPaymentAuthorize - Modelo para criação de autorização de pagamento
type BillPaymentAuthorize struct {
	ExternalTerminal string  `json:"externalTerminal"`
	ExternalNSU      int32   `json:"externalNSU"`
	BarCode          BarCode `json:"barCode"`
}

//BillData - Modelo dos valores do pagamento
type BillData struct {
	Value               float32 `json:"value"`
	OriginalValue       float32 `json:"originalValue"`
	ValueWithDiscount   float32 `json:"valueWithDiscount"`
	ValueWithAdditional float32 `json:"valueWithAdditional"`
}

//BarCode - Modelo referente ao boleto para pagamento
type BarCode struct {
	Type      int32  `json:"type"`
	Digitable string `json:"digitable"`
	BarCode   string `json:"barCode"`
}

// BarCodeRegisterData Retorna os dados do boleto a ser pago
type BarCodeRegisterData struct {
	DocumentRecipient       string  `json:"documentRecipient"`
	DocumentPayer           string  `json:"documentPayer"`
	PayDueDate              string  `json:"payDueDate"`
	NextBusinessDay         string  `json:"nextBusinessDay"`
	DueDateRegister         string  `json:"dueDateRegister"`
	AllowChangeValue        bool    `json:"allowChangeValue"`
	Recipient               string  `json:"recipient"`
	Payer                   string  `json:"payer"`
	DiscountValue           float32 `json:"discountValue"`
	InterestValueCalculated float32 `json:"interestValueCalculated"`
	MaxValue                float32 `json:"maxValue"`
	MinValue                float32 `json:"minValue"`
	FineValueCalculated     float32 `json:"fineValueCalculated"`
	OriginalValue           float32 `json:"originalValue"`
	TotalUpdated            float32 `json:"totalUpdated"`
	TotalWithDiscount       float32 `json:"totalWithDiscount"`
	TotalWithAdditional     float32 `json:"totalWithAdditional"`
}

//BillPaymentAuthorizeResponse - Retorno do BillPaymentsAuthorize
type BillPaymentAuthorizeResponse struct {
	Assignor      string              `json:"assignor"`
	RegisterData  BarCodeRegisterData `json:"registerData"`
	SettleDate    string              `json:"settleDate"`
	DueDate       string              `json:"dueDate"`
	EndHour       string              `json:"endHour"`
	IniteHour     string              `json:"initeHour"`
	NextSettle    string              `json:"nextSettle"`
	Digitable     string              `json:"digitable"`
	TransactionID int32               `json:"transactionId"`
	Type          int32               `json:"type"`
	Value         float32             `json:"value"`
	MaxValue      float32             `json:"maxValue"`
	MinValue      float32             `json:"minValue"`
	ErrorCode     string              `json:"errorCode"`
	Message       string              `json:"message"`
	Status        int                 `json:"status"`
	DocumentPayer string              `json:"documentPayer"`
}

//BillPayments - Criar um pagamento
func (celcoin *CelcoinClient) BillPayments(req BillPayment) (*BillPayment, *Error, error) {
	data, _ := json.Marshal(req)
	var response *BillPayment
	err, errAPI := celcoin.Request("POST", "transactions/billpayments", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//BillPaymentsAuthorize - Criar um autorização de pagamento
func (celcoin *CelcoinClient) BillPaymentsAuthorize(req BillPaymentAuthorize) (*BillPaymentAuthorizeResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *BillPaymentAuthorizeResponse
	err, errAPI := celcoin.Request("POST", "transactions/billpayments/authorize", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	fmt.Printf("BillPaymentAuthorizeResponse %#v", response)
	return response, nil, nil
}
