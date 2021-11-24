package seller

import (
	"time"
)

type Sell struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserId uint   `json:"userid"`
	Coin   string `gorm:"not null"  json:"coin"`
	//user input that choose limitation for their selling
	Percentage float64 `json:"jumlah_jual"`
	//quantity of an asset
	Qtt float64 `gorm:"not null" json:"qtt"`
	//asset that user got from selling
	Intake int `json:"pendapatan"`
	//posisi aktif/inaktif
	Position string `gorm:"not null;default:aktif" json:"position"`
	//harga kripto ketika dijual
	PriceSell int       `json:"price"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
