package request_input

type Shop struct {
	ID            uint   `json:"id"`
	DefaultDomain string `json:"defaultDomain"`
	Name          string `json:"name"`
}

type ShopSaveParam struct {
	UserID uint   `json:"user_id"`
	Shops  []Shop `json:"shops"`
}

// type TemplateSectionsRevertParams struct {
// 	ThemeTemplateID uint `json:"theme_template_id"`
// 	VersionID       uint `json:"version_id"`
// }
