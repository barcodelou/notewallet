package kripto

import (
	"myapp/Model/seller"
	configs "myapp/config"
)

func SellSumQtt(seller seller.Sell) float64 {
	var hasil float64
	configs.DB.Raw("SELECT SUM(remnant_qtt) FROM transactions WHERE coin = ? AND user_id = ?", seller.Coin, seller.UserId).Scan(&hasil)
	return hasil
}

func SellCountQtt(seller seller.Sell) float64 {
	var countqtt float64
	configs.DB.Raw("SELECT COUNT(qtt) FROM transactions WHERE coin = ? AND user_id = ?", seller.Coin, seller.UserId).Scan(&countqtt)
	return countqtt
}

func SellSummIntake(seller seller.Sell) float64 {
	var total float64
	configs.DB.Raw("SELECT SUM(intake) FROM sells WHERE position = ? AND user_id = ?", "aktif", seller.UserId).Scan(&total)
	return total
}

func ResultAllSell(id string) int {
	var allSell int
	configs.DB.Raw("SELECT SUM(intake) FROM sells WHERE user_id = ?", id).Scan(&allSell)
	return allSell
}

func ResultAllBuy(id string) int {
	var allBuy int
	configs.DB.Raw("SELECT SUM(outake) FROM transactions WHERE user_id = ?", id).Scan(&allBuy)
	return allBuy
}
func FindUserById(id string) string {
	var userName string
	configs.DB.Raw("SELECT name FROM users WHERE id = ?", id).Scan(&userName)
	return userName
}

func Findconclusion(allSell int, allBuy int) (string, int) {
	var adding int
	var consclusions string
	adding = allSell - allBuy
	if adding < 0 {
		consclusions = "loss"
	} else if adding == 0 {
		consclusions = "break event point"
	} else {
		consclusions = "profit"
	}
	return consclusions, adding
}

func TransactionExist(id string) (bool, bool) {
	var FoundBuy bool
	var FoundSell bool
	configs.DB.Raw("SELECT EXISTS(SELECT 1 FROM transactions WHERE user_id = ?) AS found", id).Scan(&FoundBuy)
	configs.DB.Raw("SELECT EXISTS(SELECT 1 FROM sells WHERE user_id = ?) AS found", id).Scan(&FoundSell)
	return FoundBuy, FoundSell
}
