package models

type BkTemplateSection struct {
	Model

	VersionID         uint `json:"versionId" gorm:"not null"`
	TemplateSectionID uint `json:"templateSectionId" gorm:"not null"`
	ThemeTemplateID   uint `json:"themeTemplateId" gorm:"not null"`
	ThemeID           uint `json:"themeId" gorm:"not null"`
	SectionID         uint `json:"sectionId" gorm:"not null"`
	ThemeLayoutID     uint `json:"themeLayoutId" gorm:"not null"`

	Data string `json:"data" gorm:"type:text"`

	Version TemplateSectionVersion `json:"version" gorm:"foreignKey:version_id;constraint:OnDelete:CASCADE"`
}
