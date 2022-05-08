package models

type ThemePage struct {
	Model

	ThemeID         uint   `json:"themeId" gorm:"not null"`
	PageID          uint   `json:"pageId" gorm:"not null"`          // id from online-store
	ThemeTemplateID string `json:"themeTemplateId" gorm:"not null"` // it is template suffix

	Name      string `json:"name" gorm:"size:100"`
	Component string `json:"component" gorm:"type:text"`
	PageType  string `json:"pageType" gorm:"size:20"`

	ThemeTemplate *ThemeTemplate `json:"themeTemplate"`
}
