package custom_types

import "app-api/ent/templatesection"

type TemplateSectionChangedFields struct {
	Name      *string               `json:"name,omitempty"`
	Area      *templatesection.Area `json:"area,omitempty"`
	Component *string               `json:"componenet,omitempty"`
	Position  *int                  `json:"position,omitempty"`
	Display   *bool                 `json:"display,omitempty"`
	Status    *string               `json:"status,omitempty"`
	Deleted   *bool                 `json:"deleted,omitempty"`
}
