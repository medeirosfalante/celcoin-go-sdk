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



<h3>Pagar Conta</h3>

```go
client := celcoin.NewCelcoinClient(os.Getenv("CELCOIN_USER"), os.Getenv("CELCOIN_PASSWORD"), os.Getenv("ENV"))
barcode := celcoin.BarCode{
    Type:      0,
    Digitable: "boleto digitável",
    BarCode:   "codigo de barra",
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
    Cpfcnpj:          "80213959828",
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
    panic(err)
}
if errAPI != nil {
    panic(errAPI)
}
```


