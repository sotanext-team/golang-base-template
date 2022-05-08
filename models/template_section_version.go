package models

type TemplateSectionVersion struct {
	Model `mapstructure:",squash"`

	ThemeTemplateID uint   `json:"themeTemplateId" gorm:"not null"`
	Version         string `json:"version"`
	CustomName      string `json:"name"`

	BkTemplateSections []*BkTemplateSection `json:"bkTemplateSections" gorm:"foreignKey:version_id;constraint:OnDelete:CASCADE"`
}
