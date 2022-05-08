package types

import (
	"app-api/ent"
	"app-api/models"
)

type ThemeGraphListInput struct {
	models.GraphPagination
	OrderBy *ent.ThemeOrder
	Where   *ent.ThemeWhereInput
}
