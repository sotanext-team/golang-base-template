package repository

import (
	"context"

	"app-api/ent"
	"app-api/models"
)

type TodoRepository interface {
	ListByThemeTemplateIDWithCursorPagination(client *ent.Client, ctx context.Context) ([]models.Todo, error)
}
