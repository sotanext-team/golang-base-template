package models

type GlobalTemplate struct {
	Model

	ShopID          uint `json:"shopId" gorm:"not null"`
	ThemeTemplateID uint `json:"themeTemplateId" gorm:"not null"`

	Name         string `json:"name" gorm:"size:100"`
	ViewCount    uint   `json:"viewCount" gorm:"default:0"`
	InstallCount uint   `json:"installCount" gorm:"default:0"`

	ThemeTemplate *ThemeTemplate `json:"themeTemplate"`
	Shop          *Shop          `json:"shop"`
}
