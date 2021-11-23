package user

import (
	"myapp/Model/result"
	"myapp/Model/seller"

	// "myapp/Model/seller"
	"myapp/Model/transaction"
	"time"
)

type User struct {
	ID          uint                      `gorm:"primaryKey" json:"id"`
	Name        string                    `gorm:"unique;not null"  json:"nama" `
	Email       string                    `gorm:"unique;not null" json:"email"`
	Asset       int                       `gorm:"not null" json:"asset"`
	CreatedAt   time.Time                 `gorm:"autoCreateTime"`
	Transaction []transaction.Transaction `gorm:"foreign key:UserId;references:ID"`
	UserResult  []result.Result           `gorm:"foreign key:UserId;references:ID"`
	Seller      []seller.Sell             `gorm:"foreign key:UserId;references:ID"`
}
