package models

type GlobalStyle struct {
	Model

	Name      string `json:"name" gorm:"size:100"`
	Data      string `json:"data" gorm:"type:text"`
	SettingID uint   `json:"settingId"` // I don't know what is it...
}
