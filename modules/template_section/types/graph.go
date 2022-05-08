package types

import (
	"app-api/ent"
	"app-api/models"
)

type TemplateSectionVersionGraphListInput struct {
	models.GraphPagination
	OrderBy *ent.TemplateSectionVersionOrder
	Where   *ent.TemplateSectionVersionWhereInput
}
