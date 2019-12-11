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
		Document: "602.505.310-38",
		Value:    10,
		PayerAddress: celcoin.BankslipPayerAddressRequest{
			ZipCode:    "90680-530",
			City:       "Porto Alegre",
			Additional: "",
			Street:     "Rua Rita Barem",
			State:      "RS",
		},
		PayerName:    "Laura Rayssa",
		ResponseType: 3,
		DueDate:      "2019-12-11T02:06:53.374Z",
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
