package seller

import (
	"time"
)

type Sell struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	UserId     uint    `json:"userid"`
	Coin       string  `gorm:"not null"  json:"coin"`
	Percentage float64 `json:"jumlah_jual"`
	Qtt        float64 `gorm:"not null" json:"qtt"`
	Intake     int     `json:"pendapatan"`
	//harga kripto ketika dijual
	PriceSell int       `json:"price"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
