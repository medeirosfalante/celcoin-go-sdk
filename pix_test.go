package celcoin_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/celcoin-go-sdk"
)

func TestCreatePixBrCodeStatic(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), true)
	bankTransfer := celcoin.StaticBRCodeCreationRequest{
		Amount: 10.12,
		Key:    "03602763501",
		Merchant: &celcoin.Merchant{
			PostalCode:           "88036280",
			City:                 "Florianopolis",
			MerchantCategoryCode: 0,
			Name:                 "Teste",
		},
	}
	brCoreResponse, errAPI, err := client.CreatePixBrCodeStatic(bankTransfer)
	t.Errorf("brCoreResponse : %#v", brCoreResponse)
	t.Errorf("errAPI : %#v", errAPI)
	t.Errorf("err : %#v", err)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if brCoreResponse == nil {
		t.Error("payResponse is null")
		return
	}
}

func TestGetDic(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), true)
	brCoreResponse, errAPI, err := client.GetDic("62252168048")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if brCoreResponse == nil {
		t.Error("payResponse is null")
		return
	}

	if brCoreResponse.Key == "" {
		t.Error("key is invalid")
		return
	}
}
