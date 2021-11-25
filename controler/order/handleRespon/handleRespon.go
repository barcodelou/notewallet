package handlerespon

import (
	"myapp/Model/api"
	"myapp/Model/seller"
	"myapp/Model/transaction"
	"net/http"
)

func SucsessBuy(respon api.BaseResponse, buy transaction.Transaction) interface{} {
	i := api.BaseResponse{
		Code:    http.StatusOK,
		Message: "transaction done",
		Data: map[string]interface{}{
			"coin": buy.Coin,
			"qtt":  buy.Qtt,
		},
	}
	return i
}

func SucsessSell(respon api.BaseResponse, seller seller.Sell) interface{} {
	i := api.BaseResponse{
		Code:    http.StatusOK,
		Message: "transaction done",
		Data: map[string]interface{}{
			"coin": seller.Coin,
			"qtt":  seller.Qtt,
		},
	}
	return i
}

func ErrorBuy(respon api.BaseResponse, response string) interface{} {
	i := api.BaseResponse{
		Code:    http.StatusBadRequest,
		Message: response,
	}
	return i
}

func HandlerSell(hasil float64, percentage float64, countqtt float64) []string {
	result := []string{}
	if hasil == 0 {
		result = append(result, "asset is empty")
	}
	if percentage > 1 {
		result = append(result, "to much percentage must below 1")
	}
	if countqtt == 0 {
		result = append(result, "there is no history for purchase that coin")
	}
	return result
}

func HandlerUser(input string, err error) []string {
	var result []string
	if input == "" {
		result = append(result, "username kosong")
	}
	if err != nil {
		result = append(result, "wrong email")
	}
	return result
}

func HandlerGetJwt(response api.BaseResponse, token string) interface{} {
	i := api.BaseResponse{
		Code:    http.StatusOK,
		Message: "transaction done",
		Data: map[string]interface{}{
			"token": token,
		},
	}
	return i
}
