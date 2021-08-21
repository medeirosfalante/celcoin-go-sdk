package celcoin

type MerchantBalance struct {
	Anticipated        float32 `json:"anticipated"`
	ReconcileExecuting string  `json:"reconcileExecuting"`
	Consumed           float32 `json:"consumed"`
	Credit             float32 `json:"credit"`
	Balance            float32 `json:"balance"`
	ErroCode           string  `json:"erroCode"`
	Message            string  `json:"message"`
	Status             string  `json:"status"`
}

//Balance - balance da conta
func (celcoin *CelcoinClient) Balance() (*MerchantBalance, *Error, error) {
	var response *MerchantBalance
	err, errAPI := celcoin.Request("GET", "merchant/balance", nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
