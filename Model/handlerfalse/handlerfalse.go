package handlerfalse

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
