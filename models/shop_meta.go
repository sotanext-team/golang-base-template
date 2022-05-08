package models

type ShopMeta struct {
	Model `mapstructure:",squash"`

	ShopID uint `json:"shopId" gorm:"not null"`

	Key   string `json:"key" gorm:"size:100"`
	Value string `json:"value" gorm:"type:text"`
}
