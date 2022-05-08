package models

type Widget struct {
	Model

	Name      string `json:"name" gorm:"size:100"`
	Component string `json:"component" gorm:"type:text"`

	Themes []*Theme `gorm:"many2many:theme_widgets"`
}
