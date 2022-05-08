package models

type ThemeMeta struct {
	Model

	ThemeID uint `json:"themeId" gorm:"not null"`

	Key   string `json:"key" gorm:"size:100"`
	Value string `json:"value" gorm:"type:text"`
}
