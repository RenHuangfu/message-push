// Code generated by ent, DO NOT EDIT.

package businessapp

import (
	"message-push/app/receiver/common/model/po/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldName, v))
}

// AppKey applies equality check predicate on the "app_key" field. It's identical to AppKeyEQ.
func AppKey(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldAppKey, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldAppID, v))
}

// AppType applies equality check predicate on the "app_type" field. It's identical to AppTypeEQ.
func AppType(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldAppType, v))
}

// IsDel applies equality check predicate on the "is_del" field. It's identical to IsDelEQ.
func IsDel(v int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldIsDel, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldContainsFold(FieldName, v))
}

// AppKeyEQ applies the EQ predicate on the "app_key" field.
func AppKeyEQ(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldAppKey, v))
}

// AppKeyNEQ applies the NEQ predicate on the "app_key" field.
func AppKeyNEQ(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNEQ(FieldAppKey, v))
}

// AppKeyIn applies the In predicate on the "app_key" field.
func AppKeyIn(vs ...string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldIn(FieldAppKey, vs...))
}

// AppKeyNotIn applies the NotIn predicate on the "app_key" field.
func AppKeyNotIn(vs ...string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNotIn(FieldAppKey, vs...))
}

// AppKeyGT applies the GT predicate on the "app_key" field.
func AppKeyGT(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGT(FieldAppKey, v))
}

// AppKeyGTE applies the GTE predicate on the "app_key" field.
func AppKeyGTE(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGTE(FieldAppKey, v))
}

// AppKeyLT applies the LT predicate on the "app_key" field.
func AppKeyLT(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLT(FieldAppKey, v))
}

// AppKeyLTE applies the LTE predicate on the "app_key" field.
func AppKeyLTE(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLTE(FieldAppKey, v))
}

// AppKeyContains applies the Contains predicate on the "app_key" field.
func AppKeyContains(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldContains(FieldAppKey, v))
}

// AppKeyHasPrefix applies the HasPrefix predicate on the "app_key" field.
func AppKeyHasPrefix(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldHasPrefix(FieldAppKey, v))
}

// AppKeyHasSuffix applies the HasSuffix predicate on the "app_key" field.
func AppKeyHasSuffix(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldHasSuffix(FieldAppKey, v))
}

// AppKeyEqualFold applies the EqualFold predicate on the "app_key" field.
func AppKeyEqualFold(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEqualFold(FieldAppKey, v))
}

// AppKeyContainsFold applies the ContainsFold predicate on the "app_key" field.
func AppKeyContainsFold(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldContainsFold(FieldAppKey, v))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLTE(FieldAppID, v))
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldContains(FieldAppID, v))
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldHasPrefix(FieldAppID, v))
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldHasSuffix(FieldAppID, v))
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEqualFold(FieldAppID, v))
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldContainsFold(FieldAppID, v))
}

// AppTypeEQ applies the EQ predicate on the "app_type" field.
func AppTypeEQ(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldAppType, v))
}

// AppTypeNEQ applies the NEQ predicate on the "app_type" field.
func AppTypeNEQ(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNEQ(FieldAppType, v))
}

// AppTypeIn applies the In predicate on the "app_type" field.
func AppTypeIn(vs ...string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldIn(FieldAppType, vs...))
}

// AppTypeNotIn applies the NotIn predicate on the "app_type" field.
func AppTypeNotIn(vs ...string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNotIn(FieldAppType, vs...))
}

// AppTypeGT applies the GT predicate on the "app_type" field.
func AppTypeGT(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGT(FieldAppType, v))
}

// AppTypeGTE applies the GTE predicate on the "app_type" field.
func AppTypeGTE(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGTE(FieldAppType, v))
}

// AppTypeLT applies the LT predicate on the "app_type" field.
func AppTypeLT(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLT(FieldAppType, v))
}

// AppTypeLTE applies the LTE predicate on the "app_type" field.
func AppTypeLTE(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLTE(FieldAppType, v))
}

// AppTypeContains applies the Contains predicate on the "app_type" field.
func AppTypeContains(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldContains(FieldAppType, v))
}

// AppTypeHasPrefix applies the HasPrefix predicate on the "app_type" field.
func AppTypeHasPrefix(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldHasPrefix(FieldAppType, v))
}

// AppTypeHasSuffix applies the HasSuffix predicate on the "app_type" field.
func AppTypeHasSuffix(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldHasSuffix(FieldAppType, v))
}

// AppTypeEqualFold applies the EqualFold predicate on the "app_type" field.
func AppTypeEqualFold(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEqualFold(FieldAppType, v))
}

// AppTypeContainsFold applies the ContainsFold predicate on the "app_type" field.
func AppTypeContainsFold(v string) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldContainsFold(FieldAppType, v))
}

// IsDelEQ applies the EQ predicate on the "is_del" field.
func IsDelEQ(v int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldIsDel, v))
}

// IsDelNEQ applies the NEQ predicate on the "is_del" field.
func IsDelNEQ(v int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNEQ(FieldIsDel, v))
}

// IsDelIn applies the In predicate on the "is_del" field.
func IsDelIn(vs ...int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldIn(FieldIsDel, vs...))
}

// IsDelNotIn applies the NotIn predicate on the "is_del" field.
func IsDelNotIn(vs ...int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNotIn(FieldIsDel, vs...))
}

// IsDelGT applies the GT predicate on the "is_del" field.
func IsDelGT(v int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGT(FieldIsDel, v))
}

// IsDelGTE applies the GTE predicate on the "is_del" field.
func IsDelGTE(v int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGTE(FieldIsDel, v))
}

// IsDelLT applies the LT predicate on the "is_del" field.
func IsDelLT(v int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLT(FieldIsDel, v))
}

// IsDelLTE applies the LTE predicate on the "is_del" field.
func IsDelLTE(v int) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLTE(FieldIsDel, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.BusinessApp {
	return predicate.BusinessApp(sql.FieldLTE(FieldUpdateTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BusinessApp) predicate.BusinessApp {
	return predicate.BusinessApp(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BusinessApp) predicate.BusinessApp {
	return predicate.BusinessApp(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BusinessApp) predicate.BusinessApp {
	return predicate.BusinessApp(sql.NotPredicates(p))
}
