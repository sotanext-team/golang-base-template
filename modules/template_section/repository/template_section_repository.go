package repository

import (
	"context"

	"app-api/ent"
)

type TemplateSectionRepository interface {
	// Create(client *ent.Client, ctx context.Context, templateSection ent.TemplateSection) (ent.TemplateSection, error)
	FindByID(client *ent.Client, ctx context.Context, id uint64) (*ent.TemplateSection, error)
	ListByThemeTemplateID(client *ent.Client, ctx context.Context, themeTemplateID uint64) ([]*ent.TemplateSection, error)
	ListByThemeTemplateIDWithoutDeleted(client *ent.Client, ctx context.Context, themeTemplateID uint64) ([]*ent.TemplateSection, error)
	// ListByThemeTemplateIDWithCursorPagination(client *ent.Client, ctx context.Context, themeTemplateID int, cursor string, limit int) ([]ent.TemplateSection, string, error)
	FindByIDs(client *ent.Client, ctx context.Context, ids []uint64) ([]*ent.TemplateSection, error)
	FindByIDsUnscoped(client *ent.Client, ctx context.Context, ids []uint64) ([]*ent.TemplateSection, error)
	BatchCreate(client *ent.Client, ctx context.Context, arr []*ent.TemplateSection) ([]*ent.TemplateSection, error)
	Update(client *ent.Client, ctx context.Context, templateSection *ent.TemplateSection) (*ent.TemplateSection, error)
	// Delete(client *ent.Client, ctx context.Context, templateSection *ent.TemplateSection) error
}
