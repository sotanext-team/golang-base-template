// Code generated by entc, DO NOT EDIT.

package ent

import (
	"app-api/ent/bktemplatesection"
	"app-api/ent/componentproperty"
	"app-api/ent/customcomponent"
	"app-api/ent/globaltemplate"
	"app-api/ent/schema"
	"app-api/ent/shop"
	"app-api/ent/templatesection"
	"app-api/ent/templatesectionversion"
	"app-api/ent/theme"
	"app-api/ent/themetemplate"
	"app-api/ent/todo"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bktemplatesectionMixin := schema.BkTemplateSection{}.Mixin()
	bktemplatesectionMixinFields0 := bktemplatesectionMixin[0].Fields()
	_ = bktemplatesectionMixinFields0
	bktemplatesectionMixinFields1 := bktemplatesectionMixin[1].Fields()
	_ = bktemplatesectionMixinFields1
	bktemplatesectionFields := schema.BkTemplateSection{}.Fields()
	_ = bktemplatesectionFields
	// bktemplatesectionDescCreatedAt is the schema descriptor for created_at field.
	bktemplatesectionDescCreatedAt := bktemplatesectionMixinFields1[0].Descriptor()
	// bktemplatesection.DefaultCreatedAt holds the default value on creation for the created_at field.
	bktemplatesection.DefaultCreatedAt = bktemplatesectionDescCreatedAt.Default.(func() time.Time)
	// bktemplatesectionDescUpdatedAt is the schema descriptor for updated_at field.
	bktemplatesectionDescUpdatedAt := bktemplatesectionMixinFields1[1].Descriptor()
	// bktemplatesection.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	bktemplatesection.DefaultUpdatedAt = bktemplatesectionDescUpdatedAt.Default.(func() time.Time)
	// bktemplatesection.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	bktemplatesection.UpdateDefaultUpdatedAt = bktemplatesectionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// bktemplatesectionDescID is the schema descriptor for id field.
	bktemplatesectionDescID := bktemplatesectionMixinFields0[0].Descriptor()
	// bktemplatesection.DefaultID holds the default value on creation for the id field.
	bktemplatesection.DefaultID = bktemplatesectionDescID.Default.(func() uint64)
	componentpropertyMixin := schema.ComponentProperty{}.Mixin()
	componentpropertyMixinFields0 := componentpropertyMixin[0].Fields()
	_ = componentpropertyMixinFields0
	componentpropertyMixinFields1 := componentpropertyMixin[1].Fields()
	_ = componentpropertyMixinFields1
	componentpropertyFields := schema.ComponentProperty{}.Fields()
	_ = componentpropertyFields
	// componentpropertyDescCreatedAt is the schema descriptor for created_at field.
	componentpropertyDescCreatedAt := componentpropertyMixinFields1[0].Descriptor()
	// componentproperty.DefaultCreatedAt holds the default value on creation for the created_at field.
	componentproperty.DefaultCreatedAt = componentpropertyDescCreatedAt.Default.(func() time.Time)
	// componentpropertyDescUpdatedAt is the schema descriptor for updated_at field.
	componentpropertyDescUpdatedAt := componentpropertyMixinFields1[1].Descriptor()
	// componentproperty.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	componentproperty.DefaultUpdatedAt = componentpropertyDescUpdatedAt.Default.(func() time.Time)
	// componentproperty.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	componentproperty.UpdateDefaultUpdatedAt = componentpropertyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// componentpropertyDescID is the schema descriptor for id field.
	componentpropertyDescID := componentpropertyMixinFields0[0].Descriptor()
	// componentproperty.DefaultID holds the default value on creation for the id field.
	componentproperty.DefaultID = componentpropertyDescID.Default.(func() uint64)
	customcomponentMixin := schema.CustomComponent{}.Mixin()
	customcomponentMixinFields0 := customcomponentMixin[0].Fields()
	_ = customcomponentMixinFields0
	customcomponentMixinFields1 := customcomponentMixin[1].Fields()
	_ = customcomponentMixinFields1
	customcomponentFields := schema.CustomComponent{}.Fields()
	_ = customcomponentFields
	// customcomponentDescCreatedAt is the schema descriptor for created_at field.
	customcomponentDescCreatedAt := customcomponentMixinFields1[0].Descriptor()
	// customcomponent.DefaultCreatedAt holds the default value on creation for the created_at field.
	customcomponent.DefaultCreatedAt = customcomponentDescCreatedAt.Default.(func() time.Time)
	// customcomponentDescUpdatedAt is the schema descriptor for updated_at field.
	customcomponentDescUpdatedAt := customcomponentMixinFields1[1].Descriptor()
	// customcomponent.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	customcomponent.DefaultUpdatedAt = customcomponentDescUpdatedAt.Default.(func() time.Time)
	// customcomponent.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	customcomponent.UpdateDefaultUpdatedAt = customcomponentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// customcomponentDescIsDraft is the schema descriptor for is_draft field.
	customcomponentDescIsDraft := customcomponentFields[6].Descriptor()
	// customcomponent.DefaultIsDraft holds the default value on creation for the is_draft field.
	customcomponent.DefaultIsDraft = customcomponentDescIsDraft.Default.(bool)
	// customcomponentDescID is the schema descriptor for id field.
	customcomponentDescID := customcomponentMixinFields0[0].Descriptor()
	// customcomponent.DefaultID holds the default value on creation for the id field.
	customcomponent.DefaultID = customcomponentDescID.Default.(func() uint64)
	globaltemplateMixin := schema.GlobalTemplate{}.Mixin()
	globaltemplateMixinFields0 := globaltemplateMixin[0].Fields()
	_ = globaltemplateMixinFields0
	globaltemplateMixinFields1 := globaltemplateMixin[1].Fields()
	_ = globaltemplateMixinFields1
	globaltemplateFields := schema.GlobalTemplate{}.Fields()
	_ = globaltemplateFields
	// globaltemplateDescCreatedAt is the schema descriptor for created_at field.
	globaltemplateDescCreatedAt := globaltemplateMixinFields1[0].Descriptor()
	// globaltemplate.DefaultCreatedAt holds the default value on creation for the created_at field.
	globaltemplate.DefaultCreatedAt = globaltemplateDescCreatedAt.Default.(func() time.Time)
	// globaltemplateDescUpdatedAt is the schema descriptor for updated_at field.
	globaltemplateDescUpdatedAt := globaltemplateMixinFields1[1].Descriptor()
	// globaltemplate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	globaltemplate.DefaultUpdatedAt = globaltemplateDescUpdatedAt.Default.(func() time.Time)
	// globaltemplate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	globaltemplate.UpdateDefaultUpdatedAt = globaltemplateDescUpdatedAt.UpdateDefault.(func() time.Time)
	// globaltemplateDescName is the schema descriptor for name field.
	globaltemplateDescName := globaltemplateFields[0].Descriptor()
	// globaltemplate.NameValidator is a validator for the "name" field. It is called by the builders before save.
	globaltemplate.NameValidator = func() func(string) error {
		validators := globaltemplateDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// globaltemplateDescID is the schema descriptor for id field.
	globaltemplateDescID := globaltemplateMixinFields0[0].Descriptor()
	// globaltemplate.DefaultID holds the default value on creation for the id field.
	globaltemplate.DefaultID = globaltemplateDescID.Default.(func() uint64)
	shopMixin := schema.Shop{}.Mixin()
	shopMixinFields0 := shopMixin[0].Fields()
	_ = shopMixinFields0
	shopMixinFields1 := shopMixin[1].Fields()
	_ = shopMixinFields1
	shopFields := schema.Shop{}.Fields()
	_ = shopFields
	// shopDescCreatedAt is the schema descriptor for created_at field.
	shopDescCreatedAt := shopMixinFields1[0].Descriptor()
	// shop.DefaultCreatedAt holds the default value on creation for the created_at field.
	shop.DefaultCreatedAt = shopDescCreatedAt.Default.(func() time.Time)
	// shopDescUpdatedAt is the schema descriptor for updated_at field.
	shopDescUpdatedAt := shopMixinFields1[1].Descriptor()
	// shop.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	shop.DefaultUpdatedAt = shopDescUpdatedAt.Default.(func() time.Time)
	// shop.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	shop.UpdateDefaultUpdatedAt = shopDescUpdatedAt.UpdateDefault.(func() time.Time)
	// shopDescShopName is the schema descriptor for shop_name field.
	shopDescShopName := shopFields[0].Descriptor()
	// shop.DefaultShopName holds the default value on creation for the shop_name field.
	shop.DefaultShopName = shopDescShopName.Default.(string)
	// shop.ShopNameValidator is a validator for the "shop_name" field. It is called by the builders before save.
	shop.ShopNameValidator = shopDescShopName.Validators[0].(func(string) error)
	// shopDescDefaultDomain is the schema descriptor for default_domain field.
	shopDescDefaultDomain := shopFields[1].Descriptor()
	// shop.DefaultDomainValidator is a validator for the "default_domain" field. It is called by the builders before save.
	shop.DefaultDomainValidator = func() func(string) error {
		validators := shopDescDefaultDomain.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(default_domain string) error {
			for _, fn := range fns {
				if err := fn(default_domain); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shopDescCustomDomain is the schema descriptor for custom_domain field.
	shopDescCustomDomain := shopFields[2].Descriptor()
	// shop.CustomDomainValidator is a validator for the "custom_domain" field. It is called by the builders before save.
	shop.CustomDomainValidator = shopDescCustomDomain.Validators[0].(func(string) error)
	// shopDescID is the schema descriptor for id field.
	shopDescID := shopMixinFields0[0].Descriptor()
	// shop.DefaultID holds the default value on creation for the id field.
	shop.DefaultID = shopDescID.Default.(func() uint64)
	templatesectionMixin := schema.TemplateSection{}.Mixin()
	templatesectionMixinFields0 := templatesectionMixin[0].Fields()
	_ = templatesectionMixinFields0
	templatesectionMixinFields1 := templatesectionMixin[1].Fields()
	_ = templatesectionMixinFields1
	templatesectionFields := schema.TemplateSection{}.Fields()
	_ = templatesectionFields
	// templatesectionDescCreatedAt is the schema descriptor for created_at field.
	templatesectionDescCreatedAt := templatesectionMixinFields1[0].Descriptor()
	// templatesection.DefaultCreatedAt holds the default value on creation for the created_at field.
	templatesection.DefaultCreatedAt = templatesectionDescCreatedAt.Default.(func() time.Time)
	// templatesectionDescUpdatedAt is the schema descriptor for updated_at field.
	templatesectionDescUpdatedAt := templatesectionMixinFields1[1].Descriptor()
	// templatesection.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	templatesection.DefaultUpdatedAt = templatesectionDescUpdatedAt.Default.(func() time.Time)
	// templatesection.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	templatesection.UpdateDefaultUpdatedAt = templatesectionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// templatesectionDescCid is the schema descriptor for cid field.
	templatesectionDescCid := templatesectionFields[3].Descriptor()
	// templatesection.CidValidator is a validator for the "cid" field. It is called by the builders before save.
	templatesection.CidValidator = func() func(string) error {
		validators := templatesectionDescCid.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(cid string) error {
			for _, fn := range fns {
				if err := fn(cid); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// templatesectionDescName is the schema descriptor for name field.
	templatesectionDescName := templatesectionFields[4].Descriptor()
	// templatesection.NameValidator is a validator for the "name" field. It is called by the builders before save.
	templatesection.NameValidator = func() func(string) error {
		validators := templatesectionDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// templatesectionDescPosition is the schema descriptor for position field.
	templatesectionDescPosition := templatesectionFields[7].Descriptor()
	// templatesection.DefaultPosition holds the default value on creation for the position field.
	templatesection.DefaultPosition = templatesectionDescPosition.Default.(int)
	// templatesectionDescDisplay is the schema descriptor for display field.
	templatesectionDescDisplay := templatesectionFields[8].Descriptor()
	// templatesection.DefaultDisplay holds the default value on creation for the display field.
	templatesection.DefaultDisplay = templatesectionDescDisplay.Default.(bool)
	// templatesectionDescID is the schema descriptor for id field.
	templatesectionDescID := templatesectionMixinFields0[0].Descriptor()
	// templatesection.DefaultID holds the default value on creation for the id field.
	templatesection.DefaultID = templatesectionDescID.Default.(func() uint64)
	templatesectionversionMixin := schema.TemplateSectionVersion{}.Mixin()
	templatesectionversionMixinFields0 := templatesectionversionMixin[0].Fields()
	_ = templatesectionversionMixinFields0
	templatesectionversionMixinFields1 := templatesectionversionMixin[1].Fields()
	_ = templatesectionversionMixinFields1
	templatesectionversionFields := schema.TemplateSectionVersion{}.Fields()
	_ = templatesectionversionFields
	// templatesectionversionDescCreatedAt is the schema descriptor for created_at field.
	templatesectionversionDescCreatedAt := templatesectionversionMixinFields1[0].Descriptor()
	// templatesectionversion.DefaultCreatedAt holds the default value on creation for the created_at field.
	templatesectionversion.DefaultCreatedAt = templatesectionversionDescCreatedAt.Default.(func() time.Time)
	// templatesectionversionDescUpdatedAt is the schema descriptor for updated_at field.
	templatesectionversionDescUpdatedAt := templatesectionversionMixinFields1[1].Descriptor()
	// templatesectionversion.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	templatesectionversion.DefaultUpdatedAt = templatesectionversionDescUpdatedAt.Default.(func() time.Time)
	// templatesectionversion.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	templatesectionversion.UpdateDefaultUpdatedAt = templatesectionversionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// templatesectionversionDescVersion is the schema descriptor for version field.
	templatesectionversionDescVersion := templatesectionversionFields[1].Descriptor()
	// templatesectionversion.VersionValidator is a validator for the "version" field. It is called by the builders before save.
	templatesectionversion.VersionValidator = func() func(string) error {
		validators := templatesectionversionDescVersion.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(version string) error {
			for _, fn := range fns {
				if err := fn(version); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// templatesectionversionDescName is the schema descriptor for name field.
	templatesectionversionDescName := templatesectionversionFields[2].Descriptor()
	// templatesectionversion.NameValidator is a validator for the "name" field. It is called by the builders before save.
	templatesectionversion.NameValidator = templatesectionversionDescName.Validators[0].(func(string) error)
	// templatesectionversionDescID is the schema descriptor for id field.
	templatesectionversionDescID := templatesectionversionMixinFields0[0].Descriptor()
	// templatesectionversion.DefaultID holds the default value on creation for the id field.
	templatesectionversion.DefaultID = templatesectionversionDescID.Default.(func() uint64)
	themeMixin := schema.Theme{}.Mixin()
	themeMixinFields0 := themeMixin[0].Fields()
	_ = themeMixinFields0
	themeMixinFields1 := themeMixin[1].Fields()
	_ = themeMixinFields1
	themeFields := schema.Theme{}.Fields()
	_ = themeFields
	// themeDescCreatedAt is the schema descriptor for created_at field.
	themeDescCreatedAt := themeMixinFields1[0].Descriptor()
	// theme.DefaultCreatedAt holds the default value on creation for the created_at field.
	theme.DefaultCreatedAt = themeDescCreatedAt.Default.(func() time.Time)
	// themeDescUpdatedAt is the schema descriptor for updated_at field.
	themeDescUpdatedAt := themeMixinFields1[1].Descriptor()
	// theme.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	theme.DefaultUpdatedAt = themeDescUpdatedAt.Default.(func() time.Time)
	// theme.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	theme.UpdateDefaultUpdatedAt = themeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// themeDescName is the schema descriptor for name field.
	themeDescName := themeFields[0].Descriptor()
	// theme.NameValidator is a validator for the "name" field. It is called by the builders before save.
	theme.NameValidator = func() func(string) error {
		validators := themeDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// themeDescThumbnail is the schema descriptor for thumbnail field.
	themeDescThumbnail := themeFields[1].Descriptor()
	// theme.ThumbnailValidator is a validator for the "thumbnail" field. It is called by the builders before save.
	theme.ThumbnailValidator = func() func(string) error {
		validators := themeDescThumbnail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(thumbnail string) error {
			for _, fn := range fns {
				if err := fn(thumbnail); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// themeDescPublish is the schema descriptor for publish field.
	themeDescPublish := themeFields[2].Descriptor()
	// theme.DefaultPublish holds the default value on creation for the publish field.
	theme.DefaultPublish = themeDescPublish.Default.(bool)
	// themeDescID is the schema descriptor for id field.
	themeDescID := themeMixinFields0[0].Descriptor()
	// theme.DefaultID holds the default value on creation for the id field.
	theme.DefaultID = themeDescID.Default.(func() uint64)
	themetemplateMixin := schema.ThemeTemplate{}.Mixin()
	themetemplateMixinFields0 := themetemplateMixin[0].Fields()
	_ = themetemplateMixinFields0
	themetemplateMixinFields1 := themetemplateMixin[1].Fields()
	_ = themetemplateMixinFields1
	themetemplateFields := schema.ThemeTemplate{}.Fields()
	_ = themetemplateFields
	// themetemplateDescCreatedAt is the schema descriptor for created_at field.
	themetemplateDescCreatedAt := themetemplateMixinFields1[0].Descriptor()
	// themetemplate.DefaultCreatedAt holds the default value on creation for the created_at field.
	themetemplate.DefaultCreatedAt = themetemplateDescCreatedAt.Default.(func() time.Time)
	// themetemplateDescUpdatedAt is the schema descriptor for updated_at field.
	themetemplateDescUpdatedAt := themetemplateMixinFields1[1].Descriptor()
	// themetemplate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	themetemplate.DefaultUpdatedAt = themetemplateDescUpdatedAt.Default.(func() time.Time)
	// themetemplate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	themetemplate.UpdateDefaultUpdatedAt = themetemplateDescUpdatedAt.UpdateDefault.(func() time.Time)
	// themetemplateDescName is the schema descriptor for name field.
	themetemplateDescName := themetemplateFields[1].Descriptor()
	// themetemplate.NameValidator is a validator for the "name" field. It is called by the builders before save.
	themetemplate.NameValidator = func() func(string) error {
		validators := themetemplateDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// themetemplateDescDefault is the schema descriptor for default field.
	themetemplateDescDefault := themetemplateFields[3].Descriptor()
	// themetemplate.DefaultDefault holds the default value on creation for the default field.
	themetemplate.DefaultDefault = themetemplateDescDefault.Default.(bool)
	// themetemplateDescID is the schema descriptor for id field.
	themetemplateDescID := themetemplateMixinFields0[0].Descriptor()
	// themetemplate.DefaultID holds the default value on creation for the id field.
	themetemplate.DefaultID = themetemplateDescID.Default.(func() uint64)
	todoMixin := schema.Todo{}.Mixin()
	todoMixinFields0 := todoMixin[0].Fields()
	_ = todoMixinFields0
	todoMixinFields1 := todoMixin[1].Fields()
	_ = todoMixinFields1
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoMixinFields1[0].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescUpdatedAt is the schema descriptor for updated_at field.
	todoDescUpdatedAt := todoMixinFields1[1].Descriptor()
	// todo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	todo.DefaultUpdatedAt = todoDescUpdatedAt.Default.(func() time.Time)
	// todo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	todo.UpdateDefaultUpdatedAt = todoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoFields[0].Descriptor()
	// todo.TextValidator is a validator for the "text" field. It is called by the builders before save.
	todo.TextValidator = todoDescText.Validators[0].(func(string) error)
	// todoDescPriority is the schema descriptor for priority field.
	todoDescPriority := todoFields[2].Descriptor()
	// todo.DefaultPriority holds the default value on creation for the priority field.
	todo.DefaultPriority = todoDescPriority.Default.(int)
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoMixinFields0[0].Descriptor()
	// todo.DefaultID holds the default value on creation for the id field.
	todo.DefaultID = todoDescID.Default.(func() uint64)
}
