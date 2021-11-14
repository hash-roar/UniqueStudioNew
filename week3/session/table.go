package session

import (
	"lform/schema"
	"reflect"
)

func (s *Session) Model(value interface{}) *Session {
	if s.table == nil || reflect.TypeOf(value) != reflect.TypeOf(s.table.Model) {
		s.table = schema.Parse(value, s.dialect)

	}
	return s
}

func (s *Session) GetTable() *schema.Schema {
	if s.table == nil {
		println("error model not set")
		return nil
	}
	return s.table
}
