package celcoin_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/celcoin-go-sdk"
)

func TestBankslip(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), "v5")
	bankslipRequest := celcoin.BankslipRequest{
		Values: celcoin.BankslipValues{
			OriginalValue: 100,
		},
		Payer: celcoin.BankslipPayer{
			Name:         "Ryan Cauã Felipe Araújo",
			DocumentType: "CPF",
			Document:     "80213959828",
			Address:      "Rua Francisco Alves",
			District:     "Jardim Sônia Maria",
			City:         "Mauá",
			UF:           "SP",
			ZipCode:      "09380360",
			Email:        "test@gmail.com",
			Ddd:          "11",
			PhoneNumber:  "982245864",
		},
		DueDate:      time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
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

func TestGetTransactionStatus(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), "v5")
	bankslipRequest := celcoin.BankslipRequest{
		Values: celcoin.BankslipValues{
			OriginalValue: 100,
		},
		Payer: celcoin.BankslipPayer{
			Name:         "Ryan Cauã Felipe Araújo",
			DocumentType: "CPF",
			Document:     "80213959828",
			Address:      "Rua Francisco Alves",
			District:     "Jardim Sônia Maria",
			City:         "Mauá",
			UF:           "SP",
			ZipCode:      "09380360",
			Email:        "test@gmail.com",
			Ddd:          "11",
			PhoneNumber:  "982245864",
		},
		DueDate:      time.Now().AddDate(0, 0, 3).Format(time.RFC3339),
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
	transactionResponse, errAPI, err := client.GetTransactionStatus(fmt.Sprintf("%d", payAuthorizeResponse.TransactionID))
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if transactionResponse == nil {
		t.Error("payResponse is null")
		return
	}

}
