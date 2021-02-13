package celcoin_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/celcoin-go-sdk"
)

func TestBankTransfer(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), false)
	bankTransfer := celcoin.BankTransfer{
		Document:        "80213959828",
		AccountCode:     100,
		DigitCode:       "3",
		BranchCode:      100,
		InstitutionCode: 341,
		Name:            "Ryan Cauã Felipe Araújo",
		Value:           10,
		BankAccountType: "cc",
	}
	payAuthorizeResponse, errAPI, err := client.BankTransfer(bankTransfer)
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

func TestBankTransferStatus(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), false)
	payAuthorizeResponse, errAPI, err := client.BankTransferStatus(3725218)
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
