// Code generated by ent, DO NOT EDIT.

package user

import (
	"viseh-api/database/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhoneNumber, v))
}

// EncryptionKey applies equality check predicate on the "encryption_key" field. It's identical to EncryptionKeyEQ.
func EncryptionKey(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEncryptionKey, v))
}

// MedID applies equality check predicate on the "med_id" field. It's identical to MedIDEQ.
func MedID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMedID, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldData, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// EncryptionKeyEQ applies the EQ predicate on the "encryption_key" field.
func EncryptionKeyEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEncryptionKey, v))
}

// EncryptionKeyNEQ applies the NEQ predicate on the "encryption_key" field.
func EncryptionKeyNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEncryptionKey, v))
}

// EncryptionKeyIn applies the In predicate on the "encryption_key" field.
func EncryptionKeyIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEncryptionKey, vs...))
}

// EncryptionKeyNotIn applies the NotIn predicate on the "encryption_key" field.
func EncryptionKeyNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEncryptionKey, vs...))
}

// EncryptionKeyGT applies the GT predicate on the "encryption_key" field.
func EncryptionKeyGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEncryptionKey, v))
}

// EncryptionKeyGTE applies the GTE predicate on the "encryption_key" field.
func EncryptionKeyGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEncryptionKey, v))
}

// EncryptionKeyLT applies the LT predicate on the "encryption_key" field.
func EncryptionKeyLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEncryptionKey, v))
}

// EncryptionKeyLTE applies the LTE predicate on the "encryption_key" field.
func EncryptionKeyLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEncryptionKey, v))
}

// EncryptionKeyContains applies the Contains predicate on the "encryption_key" field.
func EncryptionKeyContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEncryptionKey, v))
}

// EncryptionKeyHasPrefix applies the HasPrefix predicate on the "encryption_key" field.
func EncryptionKeyHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEncryptionKey, v))
}

// EncryptionKeyHasSuffix applies the HasSuffix predicate on the "encryption_key" field.
func EncryptionKeyHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEncryptionKey, v))
}

// EncryptionKeyEqualFold applies the EqualFold predicate on the "encryption_key" field.
func EncryptionKeyEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEncryptionKey, v))
}

// EncryptionKeyContainsFold applies the ContainsFold predicate on the "encryption_key" field.
func EncryptionKeyContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEncryptionKey, v))
}

// MedIDEQ applies the EQ predicate on the "med_id" field.
func MedIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMedID, v))
}

// MedIDNEQ applies the NEQ predicate on the "med_id" field.
func MedIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldMedID, v))
}

// MedIDIn applies the In predicate on the "med_id" field.
func MedIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldMedID, vs...))
}

// MedIDNotIn applies the NotIn predicate on the "med_id" field.
func MedIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldMedID, vs...))
}

// MedIDGT applies the GT predicate on the "med_id" field.
func MedIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldMedID, v))
}

// MedIDGTE applies the GTE predicate on the "med_id" field.
func MedIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldMedID, v))
}

// MedIDLT applies the LT predicate on the "med_id" field.
func MedIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldMedID, v))
}

// MedIDLTE applies the LTE predicate on the "med_id" field.
func MedIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldMedID, v))
}

// MedIDContains applies the Contains predicate on the "med_id" field.
func MedIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldMedID, v))
}

// MedIDHasPrefix applies the HasPrefix predicate on the "med_id" field.
func MedIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldMedID, v))
}

// MedIDHasSuffix applies the HasSuffix predicate on the "med_id" field.
func MedIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldMedID, v))
}

// MedIDEqualFold applies the EqualFold predicate on the "med_id" field.
func MedIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldMedID, v))
}

// MedIDContainsFold applies the ContainsFold predicate on the "med_id" field.
func MedIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldMedID, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.User {
	return predicate.User(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRole, vs...))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v []byte) predicate.User {
	return predicate.User(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v []byte) predicate.User {
	return predicate.User(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v []byte) predicate.User {
	return predicate.User(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v []byte) predicate.User {
	return predicate.User(sql.FieldLTE(FieldData, v))
}

// DataIsNil applies the IsNil predicate on the "data" field.
func DataIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldData))
}

// DataNotNil applies the NotNil predicate on the "data" field.
func DataNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldData))
}

// HasChips applies the HasEdge predicate on the "chips" edge.
func HasChips() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChipsTable, ChipsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChipsWith applies the HasEdge predicate on the "chips" edge with a given conditions (other predicates).
func HasChipsWith(preds ...predicate.Chip) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newChipsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
