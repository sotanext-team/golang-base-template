package models

type Theme struct {
	Model `mapstructure:",squash"`

	ShopID uint `json:"shopId" gorm:"not null"`

	Name      string `json:"name" gorm:"size:100"`
	Thumbnail string `json:"thumbnail"`
	Publish   bool   `json:"publish"`

	ThemeGlobalStyles []*ThemeGlobalStyle `json:"themeGlobalStyles"`
	ThemePages        []*ThemePage        `json:"themePages"`
	ThemeMetas        []*ThemeMeta        `json:"themeMetas"`
	Widgets           []*Widget           `json:"widgets" gorm:"many2many:theme_widgets"`
	ThemeTemplates    []*ThemeTemplate    `json:"themeTemplates"`
}
