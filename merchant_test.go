package celcoin_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/celcoin-go-sdk"
)

func TestGetBalance(t *testing.T) {
	godotenv.Load(".env.test")
	client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"), false)
	balanceResponse, errAPI, err := client.Balance()
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if balanceResponse == nil {
		t.Error("payResponse is null")
		return
	}

}
