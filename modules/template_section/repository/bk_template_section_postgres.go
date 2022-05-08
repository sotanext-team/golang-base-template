package repository

import (
	"context"

	"app-api/ent"
	"app-api/ent/bktemplatesection"

	"entgo.io/ent/dialect/sql"
)

type bkTemplateSectionImpl struct {
}

func NewBkTemplateSectionRepository() BkTemplateSectionRepository {
	return &bkTemplateSectionImpl{}
}

func (instance *bkTemplateSectionImpl) FindLastBackupWithVersion(client *ent.Client, ctx context.Context, sectionID uint64) (*ent.BkTemplateSection, error) {
	return client.BkTemplateSection.Query().
		Where(bktemplatesection.TemplateSectionIDEQ(sectionID)).
		Order(ent.Desc(bktemplatesection.FieldID)).
		WithVersion().
		First(ctx)
}

func (instance *bkTemplateSectionImpl) BatchCreate(client *ent.Client, ctx context.Context, backups []ent.BkTemplateSection) ([]*ent.BkTemplateSection, error) {
	bulk := make([]*ent.BkTemplateSectionCreate, len(backups))
	for i, backup := range backups {
		bulk[i] = client.BkTemplateSection.Create().
			SetVersionID(backup.VersionID).
			SetThemeTemplateID(backup.ThemeTemplateID).
			SetTemplateSectionID(backup.TemplateSectionID).
			SetThemeID(backup.ThemeID).
			SetThemeLayoutID(backup.ThemeLayoutID).
			SetData(backup.Data)
	}
	return client.Debug().BkTemplateSection.CreateBulk(bulk...).Save(ctx)
}

func (instance *bkTemplateSectionImpl) ListRevisionsFromID(client *ent.Client, ctx context.Context, themeTemplateID uint64, versionID uint64) ([]*ent.BkTemplateSection, error) {
	return client.BkTemplateSection.Query().
		Where(bktemplatesection.ThemeTemplateIDEQ(themeTemplateID), func(s *sql.Selector) {
			s.Where(sql.GTE(bktemplatesection.FieldVersionID, versionID))
		}).
		Order(ent.Desc(bktemplatesection.FieldVersionID)).
		All(ctx)
}
