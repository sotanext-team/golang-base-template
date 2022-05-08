package models

type TemplateSection struct {
	Model `mapstructure:",squash"`

	SectionID        uint `json:"sectionId" gorm:"not null"`
	ThemeTemplateID  uint `json:"themeTemplateId" gorm:"not null"`
	CurrentVersionID uint `json:"currentVersionId"`

	CID       string `json:"cid" gorm:"size:50"`
	Name      string `json:"name" gorm:"size:100"`
	Area      string `json:"area" gorm:"size:20"`
	Component string `json:"component" gorm:"type:text"`
	Position  int    `json:"position"`
	Display   bool   `json:"display" gorm:"default:true"`
	Status    string `json:"status"` // TODO: should have default value and enum value
	Deleted   bool   `json:"deleted" gorm:"default:false"`

	Revisions     []*BkTemplateSection `json:"revisions" gorm:"constraint:OnDelete:CASCADE"`
	ThemeTemplate *ThemeTemplate
}

func (t *TemplateSection) TransformResponseList() {
	switch t.Area {
	case "header":
		t.Area = "HEADER"
	case "main":
		t.Area = "MAIN"
	case "footer":
		t.Area = "FOOTER"
	}
	return
}
