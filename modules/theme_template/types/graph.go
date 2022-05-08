package types

import (
	"app-api/ent"
	"app-api/models"
)

type ThemeTemplateGraphListInput struct {
	models.GraphPagination
	OrderBy *ent.ThemeTemplateOrder
	Where   *ent.ThemeTemplateWhereInput
}
