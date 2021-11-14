package session

import (
	"database/sql"
	"fmt"
	"lform/clause"
	"lform/dialect"
	"lform/schema"
	"strings"
)

type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
	dialect dialect.Dialect
	table   *schema.Schema
	clause  clause.Clause
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db,
		dialect: dialect}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	sql := s.dialect.GetDialectSql(s.sql.String())
	fmt.Println(sql, "    ", s.sqlVars)
	if result, err = s.DB().Exec(sql, s.sqlVars...); err != nil {
		fmt.Println(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	fmt.Println(s.sql.String(), "    ", s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	fmt.Println(s.sql.String(), "    ", s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		fmt.Println(err)
	}
	return
}
