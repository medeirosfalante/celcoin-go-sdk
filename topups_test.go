package celcoin_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/celcoin-go-sdk"
)

func TestCreateTopupsDigitalProvider(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), false)

	TopupsRequest := celcoin.TopupsRequest{
		TopupData: celcoin.TopupData{
			Value: 20,
		},
		CpfCnpj:    "35029246002",
		SignerCode: "1111",
		ProviderID: 2132,
		Phone: celcoin.TopupPhone{
			StateCode:   "99",
			CountryCode: "55",
			Number:      "999999999",
		},
	}

	TopupsResponse, errAPI, err := client.CreateTopups(TopupsRequest)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if TopupsResponse == nil {
		t.Error("payResponse is null")
		return
	}

}
func TestListTopupsProvider(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), false)

	TopupsResponse, errAPI, err := client.ListTopupsProvider("21", "2", "0")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if TopupsResponse == nil {
		t.Error("payResponse is null")
		return
	}

}
func TestListTopupsProviderValues(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), false)
	TopupsResponse, errAPI, err := client.ListTopupsProviderValues("21", "2132")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if TopupsResponse == nil {
		t.Error("payResponse is null")
		return
	}

}
