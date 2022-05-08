package usecase

import (
	"context"

	"app-api/models"
	"app-api/modules/todo/repository"
)

type todoUseCase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUseCase() TodoUseCase {
	todoRepo := repository.NewTodoRepository()
	return &todoUseCase{
		todoRepo: todoRepo,
	}
}

type TodoUseCase interface {
	Save(context context.Context) ([]models.Todo, error)
	ListByThemeTemplateID(ctx context.Context, cursor string, limit int, themeTemplateID int) ([]models.Todo, error)
}

func (instance *todoUseCase) Save(context context.Context) ([]models.Todo, error) {
	return nil, nil
}

func (instance *todoUseCase) ListByThemeTemplateID(ctx context.Context, cursor string, limit int, themeTemplateID int) ([]models.Todo, error) {
	return nil, nil
}
