package celcoin_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/celcoin-go-sdk"
)

func TestCreatePixBrCodeStatic(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), true)
	bankTransfer := celcoin.StaticBRCodeCreationRequest{
		Amount:                    10.12,
		Key:                       "03602763501",
		TransactionIdentification: "4444",
		Merchant: &celcoin.Merchant{
			PostalCode:           "88036280",
			City:                 "Florianopolis",
			MerchantCategoryCode: 0,
			Name:                 "Teste",
		},
	}
	brCoreResponse, errAPI, err := client.CreatePixBrCodeStatic(bankTransfer)
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
	brCoreResponse, errAPI, err := client.GetDic("86403637061")
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

func TestCreatePixPayment(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), true)

	payment := celcoin.PaymentFullRequest{
		Amount:                    10.55,
		ClientCode:                "BRAIN TST",
		TransactionIdentification: "",
		Endtoendid:                "E5009404420210203130601080362644",
		DebitParty: celcoin.DebitParty{
			Account:     "212696124",
			Branch:      "0001",
			TaxID:       "86403637061",
			AccountType: "CACC",
			Name:        "Carlos Eduardo Garcia de Carvalho",
		},
		CreditParty: celcoin.CreditParty{
			Key:         "a016bc8b-d89a-4f62-be99-5fa0ebb91d77",
			Bank:        "30306294",
			Branch:      1,
			Account:     "105156485",
			AccountType: "CACC",
			TaxID:       "01548564895",
			Name:        "Teste Creditor",
		},
	}
	brCoreResponse, errAPI, err := client.CreatePixPayment(&payment)
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
