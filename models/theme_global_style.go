package models

type ThemeGlobalStyle struct {
	Model

	ThemeID       uint `json:"themeId" gorm:"not null"`
	GlobalStyleID uint `json:"globalStyleId" gorm:"not null"`

	Name    string `json:"name" gorm:"size:100"`
	Data    string `json:"data" gorm:"type:text"`
	Default bool   `json:"default" gorm:"default:false"`

	GlobalStyle *GlobalStyle `json:"globalStyle"`
}
