# celcoin-go-sdk

<h3>Initial client</h3>

```go
client := celcoin.NewCelcoinClient("user", "password")
```





<h3>Criar boleto</h3>

```go
client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"))
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
    panic(err)
}
if errAPI != nil {
    panic(errAPI)
}
```



