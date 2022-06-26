// Code generated by entc, DO NOT EDIT.

package groupmember

import (
	"github.com/AkiOuma/biz-groups/internal/interface/datasource/sqldb/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
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
func IDGT(id int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// MemberID applies equality check predicate on the "member_id" field. It's identical to MemberIDEQ.
func MemberID(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemberID), v))
	})
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroupID), v))
	})
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGroupID), v...))
	})
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGroupID), v...))
	})
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGroupID), v))
	})
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGroupID), v))
	})
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGroupID), v))
	})
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGroupID), v))
	})
}

// MemberIDEQ applies the EQ predicate on the "member_id" field.
func MemberIDEQ(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemberID), v))
	})
}

// MemberIDNEQ applies the NEQ predicate on the "member_id" field.
func MemberIDNEQ(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMemberID), v))
	})
}

// MemberIDIn applies the In predicate on the "member_id" field.
func MemberIDIn(vs ...int) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMemberID), v...))
	})
}

// MemberIDNotIn applies the NotIn predicate on the "member_id" field.
func MemberIDNotIn(vs ...int) predicate.GroupMember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMemberID), v...))
	})
}

// MemberIDGT applies the GT predicate on the "member_id" field.
func MemberIDGT(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMemberID), v))
	})
}

// MemberIDGTE applies the GTE predicate on the "member_id" field.
func MemberIDGTE(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMemberID), v))
	})
}

// MemberIDLT applies the LT predicate on the "member_id" field.
func MemberIDLT(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMemberID), v))
	})
}

// MemberIDLTE applies the LTE predicate on the "member_id" field.
func MemberIDLTE(v int) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMemberID), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
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
func Not(p predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		p(s.Not())
	})
}