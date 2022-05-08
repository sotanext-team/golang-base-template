// Code generated by entc, DO NOT EDIT.

package ent

import (
	"app-api/ent/templatesectionversion"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TemplateSectionVersion is the model entity for the TemplateSectionVersion schema.
type TemplateSectionVersion struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// ThemeTemplateID holds the value of the "theme_template_id" field.
	ThemeTemplateID uint64 `json:"themeTemplateId"`
	// Version holds the value of the "version" field.
	Version string `json:"version"`
	// Name holds the value of the "name" field.
	Name string `json:"customName"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TemplateSectionVersionQuery when eager-loading is set.
	Edges TemplateSectionVersionEdges `json:"edges"`
}

// TemplateSectionVersionEdges holds the relations/edges for other nodes in the graph.
type TemplateSectionVersionEdges struct {
	// BkTemplateSections holds the value of the bkTemplateSections edge.
	BkTemplateSections []*BkTemplateSection `json:"bkTemplateSections,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BkTemplateSectionsOrErr returns the BkTemplateSections value or an error if the edge
// was not loaded in eager-loading.
func (e TemplateSectionVersionEdges) BkTemplateSectionsOrErr() ([]*BkTemplateSection, error) {
	if e.loadedTypes[0] {
		return e.BkTemplateSections, nil
	}
	return nil, &NotLoadedError{edge: "bkTemplateSections"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TemplateSectionVersion) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case templatesectionversion.FieldID, templatesectionversion.FieldThemeTemplateID:
			values[i] = new(sql.NullInt64)
		case templatesectionversion.FieldVersion, templatesectionversion.FieldName:
			values[i] = new(sql.NullString)
		case templatesectionversion.FieldCreatedAt, templatesectionversion.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TemplateSectionVersion", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TemplateSectionVersion fields.
func (tsv *TemplateSectionVersion) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case templatesectionversion.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tsv.ID = uint64(value.Int64)
		case templatesectionversion.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tsv.CreatedAt = value.Time
			}
		case templatesectionversion.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tsv.UpdatedAt = value.Time
			}
		case templatesectionversion.FieldThemeTemplateID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field theme_template_id", values[i])
			} else if value.Valid {
				tsv.ThemeTemplateID = uint64(value.Int64)
			}
		case templatesectionversion.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				tsv.Version = value.String
			}
		case templatesectionversion.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tsv.Name = value.String
			}
		}
	}
	return nil
}

// QueryBkTemplateSections queries the "bkTemplateSections" edge of the TemplateSectionVersion entity.
func (tsv *TemplateSectionVersion) QueryBkTemplateSections() *BkTemplateSectionQuery {
	return (&TemplateSectionVersionClient{config: tsv.config}).QueryBkTemplateSections(tsv)
}

// Update returns a builder for updating this TemplateSectionVersion.
// Note that you need to call TemplateSectionVersion.Unwrap() before calling this method if this TemplateSectionVersion
// was returned from a transaction, and the transaction was committed or rolled back.
func (tsv *TemplateSectionVersion) Update() *TemplateSectionVersionUpdateOne {
	return (&TemplateSectionVersionClient{config: tsv.config}).UpdateOne(tsv)
}

// Unwrap unwraps the TemplateSectionVersion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tsv *TemplateSectionVersion) Unwrap() *TemplateSectionVersion {
	tx, ok := tsv.config.driver.(*txDriver)
	if !ok {
		panic("ent: TemplateSectionVersion is not a transactional entity")
	}
	tsv.config.driver = tx.drv
	return tsv
}

// String implements the fmt.Stringer.
func (tsv *TemplateSectionVersion) String() string {
	var builder strings.Builder
	builder.WriteString("TemplateSectionVersion(")
	builder.WriteString(fmt.Sprintf("id=%v", tsv.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(tsv.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(tsv.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", theme_template_id=")
	builder.WriteString(fmt.Sprintf("%v", tsv.ThemeTemplateID))
	builder.WriteString(", version=")
	builder.WriteString(tsv.Version)
	builder.WriteString(", name=")
	builder.WriteString(tsv.Name)
	builder.WriteByte(')')
	return builder.String()
}

// TemplateSectionVersions is a parsable slice of TemplateSectionVersion.
type TemplateSectionVersions []*TemplateSectionVersion

func (tsv TemplateSectionVersions) config(cfg config) {
	for _i := range tsv {
		tsv[_i].config = cfg
	}
}
