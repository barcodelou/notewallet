package result

type Result struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	UserId        uint   `json:"userID"`
	TransactionId uint   `json:"transactionID"`
	UserName      string `gorm:"not null" json:"userName"`
	Coin          string `gorm:"not null" json:"coin"`
	//Conclusion profit or no
	Conclusion  string `json:"Conclusion" gorm:"not null"`
	AssetResult int    `gorm:"not null" json:"assetResult"`
}
