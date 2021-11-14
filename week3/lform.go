package lform

import (
	"database/sql"
	"fmt"
	"lform/dialect"
	"lform/session"
)

type Connent struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewConnent(driver, dns string) (e *Connent, err error) {
	db, err := sql.Open(driver, dns)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = db.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	dial, ok := dialect.GetDialect(driver)
	// clause.InitDialect(dial)
	if !ok {
		fmt.Println("get dialect error")
		return
	}
	e = &Connent{db: db, dialect: dial}
	fmt.Println("connect success")
	return
}

func (Connent *Connent) Close() {
	if err := Connent.db.Close(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("close sucess")
}
func (Connent *Connent) NewSession() *session.Session {
	return session.New(Connent.db, Connent.dialect)
}
