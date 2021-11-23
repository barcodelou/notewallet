package transaction

import (
	"myapp/Model/result"

	// "myapp/Model/seller"
	"time"
)

type Transaction struct {
	ID     uint `gorm:"primaryKey;autoIncrement:true" json:"id"`
	UserId uint `json:"userId"`
	//modal yang dikeluarkan user
	Outake int `json:"jumlah_beli"`
	//harga kripto ketika dibeli
	PriceBuy int `json:"price"`
	//jumlah kripto ketika dibeli
	Qtt float64 `gorm:"not null" json:"qtt"`
	//kripto yang dibeli
	Coin string `gorm:"not null"  json:"coin"`
	//posisi aktif/inaktif
	Position string `gorm:"not null;default:aktif" json:"position"`
	//sisa qtt yang dari terjual
	RemnantQtt float64 `json:"remnant_qtt"`
	//sisa percen qtt yang dari terjual
	CreatedAt         time.Time       `gorm:"autoCreateTime"`
	TransactionResult []result.Result `gorm:"foreign key:AssetResult;references:ID" `
}
