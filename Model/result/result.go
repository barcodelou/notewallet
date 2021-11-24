package result

import "time"

type Result struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"userID"`
	UserName  string    `gorm:"not null" json:"userName"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	//Conclusion profit or no
	Conclusion string `json:"Conclusion" gorm:"not null"`
	//transaction summary between buy and sell
	AssetResult int `gorm:"not null" json:"assetResult"`
}
