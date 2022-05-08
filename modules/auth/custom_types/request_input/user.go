package request_input

import "app-api/models"

type UserInput struct {
	Username string `json:"username" gorm:"size:100"`
	Email    string `json:"email" gorm:"size:100"`
	Password string `json:"password"  gorm:"size:100"`

	Shops []*models.Shop `json:"shop" gorm:"many2many:user_shops"`
}
