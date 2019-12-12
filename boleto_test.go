package celcoin_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/celcoin-go-sdk"
)

func TestBankslip(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"))
	bankslipRequest := celcoin.BankslipRequest{
		Values: celcoin.BankslipValues{
			OriginalValue: 100,
		},
		Payer: celcoin.BankslipPayer{
			Name:         "Nicole Manuela Fátima Gonçalves",
			DocumentType: "pf",
			Document:     "69053291423",
			Address:      "Rua Florianópolis",
			District:     "Liberdade",
			City:         "Cacoal",
			State:        "RO",
			ZipCode:      "76967-412",
			Email:        "test@gmail.com",
			Ddd:          "69",
			PhoneNumber:  "26640502",
		},
		DueDate:      "2019-12-11T02:06:53.374Z",
		DaysToExpire: 3,
		Instructions: celcoin.BankslipInstructionsRequest{
			Instruction1: "Teste",
		},
	}
	payAuthorizeResponse, errAPI, err := client.Bankslip(bankslipRequest)
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
