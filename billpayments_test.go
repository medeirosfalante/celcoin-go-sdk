package celcoin_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/celcoin-go-sdk"
)

func TestBillPayments(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), false)
	barcode := celcoin.BarCode{
		Type:      0,
		Digitable: os.Getenv("CELCOIN_TEST_BOLETO_DIGITABLE"),
		BarCode:   os.Getenv("CELCOIN_TEST_BOLETO_BARCODE"),
	}
	billPaymentsAuthorize := celcoin.BillPaymentAuthorize{
		ExternalNSU:      0,
		ExternalTerminal: "1111",
		BarCode:          barcode,
	}
	payAuthorizeResponse, errAPI, err := client.BillPaymentsAuthorize(billPaymentsAuthorize)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if payAuthorizeResponse == nil {
		t.Error("payResponse is null")
		return
	}
	barcode.Type = payAuthorizeResponse.Type
	billpayments := celcoin.BillPayment{
		ExternalNSU:      0,
		ExternalTerminal: "1111",
		Cpfcnpj:          os.Getenv("CELCOIN_TEST_BOLETO_DOCUMENT"),
		BillData: celcoin.BillData{
			Value:               10,
			OriginalValue:       10,
			ValueWithDiscount:   0,
			ValueWithAdditional: 0,
		},
		BarCode:                barcode,
		DueDate:                payAuthorizeResponse.DueDate,
		TransactionIDAuthorize: payAuthorizeResponse.TransactionID,
	}
	payResponse, errAPI, err := client.BillPayments(billpayments)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if payResponse == nil {
		t.Error("payResponse is null")
		return
	}
}

func TestBillPaymentsAuthorize(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), false)
	barcode := celcoin.BarCode{
		Type:      0,
		Digitable: os.Getenv("CELCOIN_TEST_BOLETO_DIGITABLE"),
		BarCode:   os.Getenv("CELCOIN_TEST_BOLETO_BARCODE"),
	}
	billPaymentsAuthorize := celcoin.BillPaymentAuthorize{
		ExternalNSU:      0,
		ExternalTerminal: "1111",
		BarCode:          barcode,
	}
	payAuthorizeResponse, errAPI, err := client.BillPaymentsAuthorize(billPaymentsAuthorize)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if payAuthorizeResponse == nil {
		t.Error("payResponse is null")
		return
	}

}

func TestGetBillPayments(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), false)
	barcode := celcoin.BarCode{
		Type:      0,
		Digitable: os.Getenv("CELCOIN_TEST_BOLETO_DIGITABLE"),
		BarCode:   os.Getenv("CELCOIN_TEST_BOLETO_BARCODE"),
	}
	billPaymentsAuthorize := celcoin.BillPaymentAuthorize{
		ExternalNSU:      0,
		ExternalTerminal: "1111",
		BarCode:          barcode,
	}
	payAuthorizeResponse, errAPI, err := client.BillPaymentsAuthorize(billPaymentsAuthorize)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if payAuthorizeResponse == nil {
		t.Error("payResponse is null")
		return
	}
	barcode.Type = payAuthorizeResponse.Type
	billpayments := celcoin.BillPayment{
		ExternalNSU:      0,
		ExternalTerminal: "1111",
		Cpfcnpj:          os.Getenv("CELCOIN_TEST_BOLETO_DOCUMENT"),
		BillData: celcoin.BillData{
			Value:               10,
			OriginalValue:       10,
			ValueWithDiscount:   0,
			ValueWithAdditional: 0,
		},
		BarCode:                barcode,
		DueDate:                payAuthorizeResponse.DueDate,
		TransactionIDAuthorize: payAuthorizeResponse.TransactionID,
	}
	payResponse, errAPI, err := client.BillPayments(billpayments)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if payResponse == nil {
		t.Error("payResponse is null")
		return
	}
	paymentResponse, errAPI, err := client.GetBillPayments(fmt.Sprintf("%d", payResponse.TransactionID))
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if paymentResponse == nil {
		t.Error("payResponse is null")
		return
	}

}
