package models

import "time"

type UserShop struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uint      `gorm:"primaryKey"`
	ShopID    uint      `gorm:"primaryKey"`
	IsMain    bool

	Shop *Shop `json:"shop" gorm:"foreignKey:shop_id"`
	User *User `json:"user" gorm:"foreignKey:user_id"`
}
