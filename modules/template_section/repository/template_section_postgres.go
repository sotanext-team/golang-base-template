package repository

import (
	"context"

	"app-api/ent"
	"app-api/ent/templatesection"
)

type templateSectionImpl struct {
}

func NewTemplateSectionRepository() TemplateSectionRepository {

	return &templateSectionImpl{}
}

// func (instance *templateSectionImpl) Create(client *ent.Client, ctx context.Context, templateSection models.TemplateSection) (models.TemplateSection, error) {
// 	err := client.Create(&templateSection).Error
// 	if err != nil {
// 		return models.TemplateSection{}, err
// 	}

// 	return templateSection, nil
// }

func (instance *templateSectionImpl) FindByID(client *ent.Client, ctx context.Context, id uint64) (*ent.TemplateSection, error) {
	return client.TemplateSection.Query().
		Where(templatesection.IDEQ(id)).First(ctx)
}

func (instance *templateSectionImpl) ListByThemeTemplateID(client *ent.Client, ctx context.Context, themeTemplateID uint64) ([]*ent.TemplateSection, error) {
	return client.Debug().TemplateSection.Query().
		Where(templatesection.ThemeTemplateID(themeTemplateID)).
		All(ctx)
}

func (instance *templateSectionImpl) ListByThemeTemplateIDWithoutDeleted(client *ent.Client, ctx context.Context, themeTemplateID uint64) ([]*ent.TemplateSection, error) {
	return client.TemplateSection.Query().
		Where(templatesection.ThemeTemplateIDEQ(themeTemplateID), templatesection.DeletedAtIsNil()).
		All(ctx)
}

// func (instance *templateSectionImpl) ListByThemeTemplateIDWithCursorPagination(client *ent.Client, ctx context.Context, themeTemplateID int, cursor string, limit int) (list []models.TemplateSection, endCursor string, err error) {
// 	decodedCursor, err := esUtils.Base64DecodeStripped(cursor)
// 	if err != nil && cursor != "" {
// 		return nil, "", err
// 	}

// 	queryBuilder := client.Where("theme_template_id = ?", themeTemplateID)

// 	if limit > 0 {
// 		queryBuilder = queryBuilder.Limit(limit)
// 	}

// 	if decodedCursor != "" {
// 		i, err := strconv.ParseInt(decodedCursor, 10, 64)
// 		if err != nil {
// 			panic(err)
// 		}
// 		tm := time.UnixMicro(i)
// 		fmt.Println(tm)

// 		queryBuilder = queryBuilder.Where("created_at > ?", tm)
// 	}

// 	err = queryBuilder.Order("created_at ASC").Find(&list).Error

// 	// err = client.Where("theme_template_id = ? AND created_at > ?", themeTemplateID, decodedCursor).Limit(limit).Find(&list).Error

// 	// nextCursor := ""
// 	// if len(list) == limit {
// 	if len(list) > 0 {
// 		endCursor = esUtils.Base64EncodeStripped(esUtils.NumberToString(list[len(list)-1].CreatedAt.UnixMicro))
// 	}
// 	return
// }

func (instance *templateSectionImpl) FindByIDs(client *ent.Client, ctx context.Context, ids []uint64) ([]*ent.TemplateSection, error) {
	return client.TemplateSection.Query().
		Where(templatesection.IDIn(ids...), templatesection.DeletedAtIsNil()).All(ctx)
}

func (instance *templateSectionImpl) FindByIDsUnscoped(client *ent.Client, ctx context.Context, ids []uint64) ([]*ent.TemplateSection, error) {
	return client.TemplateSection.Query().
		Where(templatesection.IDIn(ids...)).All(ctx)
}

func (instance *templateSectionImpl) BatchCreate(client *ent.Client, ctx context.Context, templateSections []*ent.TemplateSection) ([]*ent.TemplateSection, error) {
	bulk := make([]*ent.TemplateSectionCreate, len(templateSections))
	for i, templateSection := range templateSections {
		bulk[i] = client.TemplateSection.Create().
			SetGlobalSectionID(templateSection.GlobalSectionID).
			SetThemeTemplateID(templateSection.ThemeTemplateID).
			SetCurrentVersionID(templateSection.CurrentVersionID).
			SetCid(templateSection.Cid).
			SetName(templateSection.Name).
			SetArea(templateSection.Area).
			SetComponent(templateSection.Component).
			SetPosition(templateSection.Position).
			SetDisplay(templateSection.Display)
	}
	return client.TemplateSection.CreateBulk(bulk...).Save(ctx)
}

func (instance *templateSectionImpl) Update(client *ent.Client, ctx context.Context, templateSection *ent.TemplateSection) (*ent.TemplateSection, error) {
	return client.TemplateSection.
		UpdateOneID(templateSection.ID).
		SetCurrentVersionID(templateSection.CurrentVersionID).
		SetCid(templateSection.Cid).
		SetName(templateSection.Name).
		SetArea(templateSection.Area).
		SetComponent(templateSection.Component).
		SetPosition(templateSection.Position).
		SetDisplay(templateSection.Display).
		Save(ctx)
}

// func (instance *templateSectionImpl) Delete(client *ent.Client, ctx context.Context, templateSection *models.TemplateSection) error {
// 	if err := client.Delete(&templateSection).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
