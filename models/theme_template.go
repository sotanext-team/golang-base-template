package models

type ThemeTemplate struct {
	Model `mapstructure:",squash"`

	ThemeID            uint `json:"themeId" gorm:"not null"`
	ThemeGlobalStyleID uint `json:"themeGlobalStyleId"`
	ShopID             uint `json:"shopId" gorm:"not null"`

	Name       string `json:"name" gorm:"size:1000"`
	PageType   string `json:"pageType" gorm:"size:20"`
	IsDefault  bool   `json:"isDefault" gorm:"default:false"`
	IsExported bool   `json:"isExported" gorm:"default:false"`

	TemplateSections       []*TemplateSection        `gorm:"constraint:OnDelete:CASCADE"`
	TemplateSectionVersion []*TemplateSectionVersion `gorm:"constraint:OnDelete:CASCADE"`
}
