package models

type ShopEvent struct {
	Model `mapstructure:",squash"`

	ShopID uint `json:"shopId" gorm:"not null"`

	EventType string `json:"eventType" gorm:"size:100"`
	Data      string `json:"data" gorm:"type:text"`
}
