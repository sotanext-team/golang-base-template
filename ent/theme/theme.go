// Code generated by entc, DO NOT EDIT.

package theme

import (
	"time"
)

const (
	// Label holds the string label denoting the theme type in the database.
	Label = "theme"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldShopID holds the string denoting the shop_id field in the database.
	FieldShopID = "shop_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldThumbnail holds the string denoting the thumbnail field in the database.
	FieldThumbnail = "thumbnail"
	// FieldPublish holds the string denoting the publish field in the database.
	FieldPublish = "publish"
	// EdgeThemeTemplates holds the string denoting the themetemplates edge name in mutations.
	EdgeThemeTemplates = "themeTemplates"
	// EdgeShop holds the string denoting the shop edge name in mutations.
	EdgeShop = "shop"
	// Table holds the table name of the theme in the database.
	Table = "themes"
	// ThemeTemplatesTable is the table that holds the themeTemplates relation/edge.
	ThemeTemplatesTable = "theme_templates"
	// ThemeTemplatesInverseTable is the table name for the ThemeTemplate entity.
	// It exists in this package in order to avoid circular dependency with the "themetemplate" package.
	ThemeTemplatesInverseTable = "theme_templates"
	// ThemeTemplatesColumn is the table column denoting the themeTemplates relation/edge.
	ThemeTemplatesColumn = "theme_id"
	// ShopTable is the table that holds the shop relation/edge.
	ShopTable = "themes"
	// ShopInverseTable is the table name for the Shop entity.
	// It exists in this package in order to avoid circular dependency with the "shop" package.
	ShopInverseTable = "shops"
	// ShopColumn is the table column denoting the shop relation/edge.
	ShopColumn = "shop_id"
)

// Columns holds all SQL columns for theme fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldShopID,
	FieldName,
	FieldThumbnail,
	FieldPublish,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// ThumbnailValidator is a validator for the "thumbnail" field. It is called by the builders before save.
	ThumbnailValidator func(string) error
	// DefaultPublish holds the default value on creation for the "publish" field.
	DefaultPublish bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uint64
)