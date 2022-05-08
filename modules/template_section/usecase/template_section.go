package usecase

import (
	"context"
	"encoding/json"

	"app-api/constants"
	"app-api/db"
	"app-api/ent"
	graphModels "app-api/graph/models"
	"app-api/modules/template_section/custom_types"
	"app-api/modules/template_section/custom_types/request_input"
	"app-api/modules/template_section/repository"
	_themeTemplateRepo "app-api/modules/theme_template/repository"

	esUtils "github.com/es-hs/es-helper/utils"
	esVersionUtils "github.com/es-hs/es-helper/version"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type templateSectionUseCase struct {
	templateSectionRepo        repository.TemplateSectionRepository
	templateSectionVersionRepo repository.TemplateSectionVersionRepository
	bkTemplateSectionRepo      repository.BkTemplateSectionRepository
	themeTemplateRepo          _themeTemplateRepo.ThemeTemplateRepository
}

func NewTemplateSectionUseCase() TemplateSectionUseCase {
	templateSectionRepo := repository.NewTemplateSectionRepository()
	templateSectionVersionRepo := repository.NewTemplateSectionVersionRepository()
	bkTemplateSectionRepo := repository.NewBkTemplateSectionRepository()
	themeTemplateRepo := _themeTemplateRepo.NewThemeTemplateRepository()
	return &templateSectionUseCase{
		templateSectionRepo:        templateSectionRepo,
		templateSectionVersionRepo: templateSectionVersionRepo,
		bkTemplateSectionRepo:      bkTemplateSectionRepo,
		themeTemplateRepo:          themeTemplateRepo,
	}
}

type TemplateSectionUseCase interface {
	Save(ctx context.Context, params request_input.TemplateSectionsSaveParams) ([]*ent.TemplateSection, error)
	Revert(ctx context.Context, params request_input.TemplateSectionsRevertParams) ([]*ent.TemplateSection, error)
	ListByThemeTemplateID(ctx context.Context, themeTemplateID uint64) ([]*ent.TemplateSection, error)
	FindByID(ctx context.Context, id uint64) (*ent.TemplateSection, error)
}

func (instance *templateSectionUseCase) Save(ctx context.Context, params request_input.TemplateSectionsSaveParams) ([]*ent.TemplateSection, error) {
	themeTemplateID := params.ThemeTemplateID
	inputTemplateSections := params.Sections
	saveType := params.SaveType
	logrus.Info("xin day dkm")
	// find create and update then backup
	var (
		inputIDs                                                        []uint64
		result, existingTemplateSections, shouldCreateTemplateSections  []*ent.TemplateSection
		shouldUpdateTemplateSections                                    []graphModels.TemplateSectionInput
		bkTemplateSections                                              []ent.BkTemplateSection
		mapIndexInputTemplateSections, mapIndexExistingTemplateSections map[int]map[string]interface{}
		err                                                             error
		themeID                                                         uint64
	)

	shopValue := ctx.Value("currentShop")

	currentShop := shopValue.(*ent.Shop)

	// Find themeTemplate of currentShop, raise exception if not found
	themeTemplate, err := instance.themeTemplateRepo.FindByID(db.GetClient(), ctx, currentShop, uint64(params.ThemeTemplateID))
	if err != nil || themeTemplate.ID == 0 {
		return nil, err
	}

	themeID = themeTemplate.ThemeID

	for _, inputTemplateSection := range inputTemplateSections {
		inputIDs = append(inputIDs, inputTemplateSection.ID)
		templateSection := ent.TemplateSection{}
		copier.Copy(&templateSection, inputTemplateSection)
		if inputTemplateSection.ID > 0 {
			shouldUpdateTemplateSections = append(shouldUpdateTemplateSections, *inputTemplateSection)
		} else {
			templateSection.ThemeTemplateID = themeTemplateID
			shouldCreateTemplateSections = append(shouldCreateTemplateSections, &templateSection)
		}
	}

	existingTemplateSections, err = instance.templateSectionRepo.FindByIDs(db.GetClient(), ctx, inputIDs)
	if err != nil {
		return nil, err
	}

	// Find changed field then make a backup
	mapIndexInputTemplateSections, err = instance.convertArrayInputTemplateSectionToArrayMapIndex(inputTemplateSections)
	if err != nil {
		return nil, err
	}

	mapIndexExistingTemplateSections, err = instance.convertArrayTemplateSectionToArrayMapIndex(existingTemplateSections)
	if err != nil {
		return nil, err
	}

	for _, sectionTemplate := range inputTemplateSections {
		if sectionTemplate.ID > 0 {
			mapInputSectionTemplate := mapIndexInputTemplateSections[int(sectionTemplate.ID)]
			mapDBSectionTemplate := mapIndexExistingTemplateSections[int(sectionTemplate.ID)]
			changed := esUtils.DiffMap(mapInputSectionTemplate, mapDBSectionTemplate)
			if len(changed) > 0 {
				dataChanged, err := esUtils.MapToJson(changed)
				if err != nil {
					return nil, err
				}
				// TODO: Do not hard code
				bkTemplateSections = append(bkTemplateSections, ent.BkTemplateSection{
					// VersionID:         1, <- add later
					TemplateSectionID: sectionTemplate.ID,
					ThemeTemplateID:   themeTemplateID,
					ThemeID:           themeID,
					// SectionID:         sectionTemplate.SectionID,
					ThemeLayoutID: 1,
					Data:          dataChanged,
				})
			}
		}
	}

	err = db.WithTx(ctx, db.GetClient(), func(tx *ent.Tx) error {
		var newVersion ent.TemplateSectionVersion
		var versionCreated *ent.TemplateSectionVersion
		client := tx.Client()
		if len(bkTemplateSections) > 0 {
			lastVersion, err := instance.templateSectionVersionRepo.FindLastVersion(client, ctx, themeTemplateID)
			if err != nil && !ent.IsNotFound(err) {
				return err
			}
			if !ent.IsNotFound(err) {
				lastVersionValue := lastVersion.Version
				var location string
				switch saveType {
				case "AUTO_SAVE":
					location = "patch"
				case "NORMAL_SAVE":
					location = "minor"
				}
				nextVersion, err := esVersionUtils.IncreaseVersion(lastVersionValue, location)
				if err != nil {
					return err
				}

				newVersion = ent.TemplateSectionVersion{
					ThemeTemplateID: themeTemplateID,
					Version:         nextVersion,
				}

				if versionCreated, err = instance.templateSectionVersionRepo.Create(client, ctx, &newVersion); err != nil {
					return err
				}
			} else {
				// Create new version if there are any changes
				newVersion = ent.TemplateSectionVersion{
					ThemeTemplateID: themeTemplateID,
					Version:         constants.INIT_VERSION,
				}
				if versionCreated, err = instance.templateSectionVersionRepo.Create(client, ctx, &newVersion); err != nil {
					return err
				}
			}

			for index := range bkTemplateSections {
				bkTemplateSections[index].VersionID = versionCreated.ID
			}

			_, err = instance.bkTemplateSectionRepo.BatchCreate(client, ctx, bkTemplateSections)
			if err != nil {
				return err
			}
		}

		shouldCreateTemplateSections, err = instance.templateSectionRepo.BatchCreate(client, ctx, shouldCreateTemplateSections)
		if err != nil {
			return err
		}
		// Make backup before update
		for _, templateSection := range shouldUpdateTemplateSections {

			// Find templateSection in existingTemplateSections
			existTemplateSection := mapIndexExistingTemplateSections[int(templateSection.ID)]

			ts := ent.TemplateSection{}
			err = esUtils.MapToStruct(existTemplateSection, &ts)
			if err != nil {
				return err
			}

			err = copier.Copy(&ts, templateSection)
			if err != nil {
				return err
			}

			if newVersion.ID > 0 {
				ts.CurrentVersionID = versionCreated.ID
			}

			// if templateSection.DeletedAt.Valid != true {
			// templateSection.DeletedAt = gorm.DeletedAt{}
			var updatedTemplateSection *ent.TemplateSection
			updatedTemplateSection, err = instance.templateSectionRepo.Update(client, ctx, &ts)
			// ts.TransformResponseList()
			result = append(result, updatedTemplateSection)
			// } else {
			// 	err = instance.templateSectionRepo.Delete(tx, &templateSection)
			// }
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// for index := range shouldCreateTemplateSections {
	// 	shouldCreateTemplateSections[index].TransformResponseList()
	// }
	result = append(result, shouldCreateTemplateSections...)
	// result = append(result, shouldUpdateTemplateSections...)

	return result, nil
	// return instance.templateSectionRepo.Create(instance.database, inputTemplateSection)
}

func (instance *templateSectionUseCase) Revert(ctx context.Context, params request_input.TemplateSectionsRevertParams) ([]*ent.TemplateSection, error) {
	versionID := params.VersionID

	var (
		revisions []*ent.BkTemplateSection
		sections  []*ent.TemplateSection

		err error
	)

	if revisions, err = instance.bkTemplateSectionRepo.ListRevisionsFromID(db.GetClient(), ctx, params.ThemeTemplateID, versionID); err != nil {
		return nil, err
	}

	if sections, err = instance.templateSectionRepo.ListByThemeTemplateID(db.GetClient(), ctx, params.ThemeTemplateID); err != nil {
		return nil, err
	}

	mapSections := make(map[uint64]*ent.TemplateSection)
	for _, section := range sections {
		mapSections[section.ID] = section
	}

	for _, revision := range revisions {

		changedData := custom_types.TemplateSectionChangedFields{}
		err = json.Unmarshal([]byte(revision.Data), &changedData)
		if err != nil {
			return nil, err
		}
		mapSections[revision.TemplateSectionID] = instance.revertTemplateSectionField(*mapSections[revision.TemplateSectionID], changedData)
	}

	sections = nil
	for _, themeSection := range mapSections {
		themeSection.CurrentVersionID = versionID
		// themeSection.TransformResponseList()
		sections = append(sections, themeSection)
	}

	err = db.WithTx(ctx, db.GetClient(), func(tx *ent.Tx) error {
		client := tx.Client()
		for _, themeSection := range sections {
			_, err = instance.templateSectionRepo.Update(client, ctx, themeSection)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return sections, nil
}

func (instance *templateSectionUseCase) revertTemplateSectionField(currentData ent.TemplateSection, changedData custom_types.TemplateSectionChangedFields) *ent.TemplateSection {
	// v := reflect.ValueOf(changedData)
	// typeOfS := v.Type()

	// for i := 0; i < v.NumField(); i++ {
	// 	field := typeOfS.Field(i).Name
	// 	value := v.Field(i).Interface()
	// 	fmt.Printf("Field: %s\tValue: %+v\n", field, value)
	// 	if value != nil {
	// 	}
	// }
	// TODO: iterate through the fields of struct
	if changedData.Name != nil {
		currentData.Name = *changedData.Name
	}
	if changedData.Area != nil {
		currentData.Area = *changedData.Area
	}
	if changedData.Component != nil {
		currentData.Component = *changedData.Component
	}
	if changedData.Position != nil {
		currentData.Position = *changedData.Position
	}
	if changedData.Display != nil {
		currentData.Display = *changedData.Display
	}
	// if changedData.Deleted != nil {
	// 	currentData.Deleted = *changedData.Deleted
	// }
	return &currentData
}

func (instance *templateSectionUseCase) ListByThemeTemplateID(ctx context.Context, themeTemplateID uint64) ([]*ent.TemplateSection, error) {
	return instance.templateSectionRepo.ListByThemeTemplateIDWithoutDeleted(db.GetClient(), ctx, themeTemplateID)

}

func (instance *templateSectionUseCase) FindByID(ctx context.Context, id uint64) (*ent.TemplateSection, error) {
	return instance.templateSectionRepo.FindByID(db.GetClient(), ctx, id)
}

func (instance *templateSectionUseCase) convertArrayTemplateSectionToArrayMapIndex(templateSections []*ent.TemplateSection) (map[int]map[string]interface{}, error) {
	result := map[int]map[string]interface{}{}
	for _, templateSection := range templateSections {
		mapTemplateSection, err := esUtils.StructToMap(templateSection)
		if err != nil {
			return nil, err
		}
		result[int(templateSection.ID)] = mapTemplateSection
		// result = append(result, mapTemplateSection)

	}
	return result, nil
}

func (instance *templateSectionUseCase) convertArrayInputTemplateSectionToArrayMapIndex(templateSections []*graphModels.TemplateSectionInput) (map[int]map[string]interface{}, error) {
	result := map[int]map[string]interface{}{}
	for _, templateSection := range templateSections {
		mapTemplateSection, err := esUtils.StructToMap(templateSection)
		if err != nil {
			return nil, err
		}
		result[int(templateSection.ID)] = mapTemplateSection
		// result = append(result, mapTemplateSection)

	}
	return result, nil
}
