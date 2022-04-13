// Code generated by entc, DO NOT EDIT.

package messagewithenum

import (
	"entgo.io/ent/dialect/sql"
	"github.com/asfsadas/contrib/entproto/internal/entprototest/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
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
func IDGT(id int) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EnumTypeEQ applies the EQ predicate on the "enum_type" field.
func EnumTypeEQ(v EnumType) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnumType), v))
	})
}

// EnumTypeNEQ applies the NEQ predicate on the "enum_type" field.
func EnumTypeNEQ(v EnumType) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnumType), v))
	})
}

// EnumTypeIn applies the In predicate on the "enum_type" field.
func EnumTypeIn(vs ...EnumType) predicate.MessageWithEnum {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEnumType), v...))
	})
}

// EnumTypeNotIn applies the NotIn predicate on the "enum_type" field.
func EnumTypeNotIn(vs ...EnumType) predicate.MessageWithEnum {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEnumType), v...))
	})
}

// EnumWithoutDefaultEQ applies the EQ predicate on the "enum_without_default" field.
func EnumWithoutDefaultEQ(v EnumWithoutDefault) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnumWithoutDefault), v))
	})
}

// EnumWithoutDefaultNEQ applies the NEQ predicate on the "enum_without_default" field.
func EnumWithoutDefaultNEQ(v EnumWithoutDefault) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnumWithoutDefault), v))
	})
}

// EnumWithoutDefaultIn applies the In predicate on the "enum_without_default" field.
func EnumWithoutDefaultIn(vs ...EnumWithoutDefault) predicate.MessageWithEnum {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEnumWithoutDefault), v...))
	})
}

// EnumWithoutDefaultNotIn applies the NotIn predicate on the "enum_without_default" field.
func EnumWithoutDefaultNotIn(vs ...EnumWithoutDefault) predicate.MessageWithEnum {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEnumWithoutDefault), v...))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MessageWithEnum) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MessageWithEnum) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
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
func Not(p predicate.MessageWithEnum) predicate.MessageWithEnum {
	return predicate.MessageWithEnum(func(s *sql.Selector) {
		p(s.Not())
	})
}
