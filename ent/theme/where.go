// Code generated by entc, DO NOT EDIT.

package theme

import (
	"app-api/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// ShopID applies equality check predicate on the "shop_id" field. It's identical to ShopIDEQ.
func ShopID(v uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShopID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Thumbnail applies equality check predicate on the "thumbnail" field. It's identical to ThumbnailEQ.
func Thumbnail(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThumbnail), v))
	})
}

// Publish applies equality check predicate on the "publish" field. It's identical to PublishEQ.
func Publish(v bool) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublish), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedAt)))
	})
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedAt)))
	})
}

// ShopIDEQ applies the EQ predicate on the "shop_id" field.
func ShopIDEQ(v uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShopID), v))
	})
}

// ShopIDNEQ applies the NEQ predicate on the "shop_id" field.
func ShopIDNEQ(v uint64) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShopID), v))
	})
}

// ShopIDIn applies the In predicate on the "shop_id" field.
func ShopIDIn(vs ...uint64) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldShopID), v...))
	})
}

// ShopIDNotIn applies the NotIn predicate on the "shop_id" field.
func ShopIDNotIn(vs ...uint64) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldShopID), v...))
	})
}

// ShopIDIsNil applies the IsNil predicate on the "shop_id" field.
func ShopIDIsNil() predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldShopID)))
	})
}

// ShopIDNotNil applies the NotNil predicate on the "shop_id" field.
func ShopIDNotNil() predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldShopID)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ThumbnailEQ applies the EQ predicate on the "thumbnail" field.
func ThumbnailEQ(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThumbnail), v))
	})
}

// ThumbnailNEQ applies the NEQ predicate on the "thumbnail" field.
func ThumbnailNEQ(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThumbnail), v))
	})
}

// ThumbnailIn applies the In predicate on the "thumbnail" field.
func ThumbnailIn(vs ...string) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldThumbnail), v...))
	})
}

// ThumbnailNotIn applies the NotIn predicate on the "thumbnail" field.
func ThumbnailNotIn(vs ...string) predicate.Theme {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Theme(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldThumbnail), v...))
	})
}

// ThumbnailGT applies the GT predicate on the "thumbnail" field.
func ThumbnailGT(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThumbnail), v))
	})
}

// ThumbnailGTE applies the GTE predicate on the "thumbnail" field.
func ThumbnailGTE(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThumbnail), v))
	})
}

// ThumbnailLT applies the LT predicate on the "thumbnail" field.
func ThumbnailLT(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThumbnail), v))
	})
}

// ThumbnailLTE applies the LTE predicate on the "thumbnail" field.
func ThumbnailLTE(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThumbnail), v))
	})
}

// ThumbnailContains applies the Contains predicate on the "thumbnail" field.
func ThumbnailContains(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThumbnail), v))
	})
}

// ThumbnailHasPrefix applies the HasPrefix predicate on the "thumbnail" field.
func ThumbnailHasPrefix(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThumbnail), v))
	})
}

// ThumbnailHasSuffix applies the HasSuffix predicate on the "thumbnail" field.
func ThumbnailHasSuffix(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThumbnail), v))
	})
}

// ThumbnailEqualFold applies the EqualFold predicate on the "thumbnail" field.
func ThumbnailEqualFold(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThumbnail), v))
	})
}

// ThumbnailContainsFold applies the ContainsFold predicate on the "thumbnail" field.
func ThumbnailContainsFold(v string) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThumbnail), v))
	})
}

// PublishEQ applies the EQ predicate on the "publish" field.
func PublishEQ(v bool) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublish), v))
	})
}

// PublishNEQ applies the NEQ predicate on the "publish" field.
func PublishNEQ(v bool) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPublish), v))
	})
}

// HasThemeTemplates applies the HasEdge predicate on the "themeTemplates" edge.
func HasThemeTemplates() predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ThemeTemplatesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ThemeTemplatesTable, ThemeTemplatesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasThemeTemplatesWith applies the HasEdge predicate on the "themeTemplates" edge with a given conditions (other predicates).
func HasThemeTemplatesWith(preds ...predicate.ThemeTemplate) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ThemeTemplatesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ThemeTemplatesTable, ThemeTemplatesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShop applies the HasEdge predicate on the "shop" edge.
func HasShop() predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShopTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShopTable, ShopColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShopWith applies the HasEdge predicate on the "shop" edge with a given conditions (other predicates).
func HasShopWith(preds ...predicate.Shop) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShopInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShopTable, ShopColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Theme) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Theme) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Theme) predicate.Theme {
	return predicate.Theme(func(s *sql.Selector) {
		p(s.Not())
	})
}