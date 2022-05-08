package repository

import (
	"context"

	"app-api/ent"
)

type BkTemplateSectionRepository interface {
	FindLastBackupWithVersion(client *ent.Client, ctx context.Context, sectionID uint64) (*ent.BkTemplateSection, error)
	BatchCreate(client *ent.Client, ctx context.Context, backups []ent.BkTemplateSection) ([]*ent.BkTemplateSection, error)
	ListRevisionsFromID(client *ent.Client, ctx context.Context, themeTemplateID uint64, versionID uint64) ([]*ent.BkTemplateSection, error)
}
