package models

type UserMeta struct {
	Model `mapstructure:",squash"`

	UserID uint `json:"userId" gorm:"not null"`

	Key   string `json:"key" gorm:"size:100"`
	Value string `json:"value" gorm:"type:text"`
}
