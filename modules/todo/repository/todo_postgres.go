package repository

import (
	"context"

	"app-api/ent"
	"app-api/models"
)

type todoImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &todoImpl{}
}

func (instance *todoImpl) ListByThemeTemplateIDWithCursorPagination(client *ent.Client, context context.Context) ([]models.Todo, error) {
	return nil, nil
}
