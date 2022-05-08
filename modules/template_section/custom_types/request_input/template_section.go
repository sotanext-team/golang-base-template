package request_input

import "app-api/graph/models"

type TemplateSection struct {
	ID        uint   `json:"id"`
	Area      string `json:"area"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Position  int    `json:"position"`
	Display   bool   `json:"display"`
	CID       string `json:"cid"`
	SectionID uint   `json:"sectionId"`
	Deleted   bool   `json:"deleted"`

	ThemeTemplateID uint `json:"-"`
}

type TemplateSectionsSaveParams struct {
	SaveType        models.SaveType                `json:"saveType"`
	ThemeTemplateID uint64                         `json:"themeTemplateId"`
	ThemeID         uint64                         `json:"themeId"`
	Sections        []*models.TemplateSectionInput `json:"sections"`
}

type TemplateSectionsRevertParams struct {
	ThemeTemplateID uint64 `json:"themeTemplateId"`
	VersionID       uint64 `json:"versionId"`
}
