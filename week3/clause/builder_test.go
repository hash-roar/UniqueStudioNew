package clause_test

import (
	"fmt"
	"lform/clause"
	"testing"
)

type User struct {
	Id   int `lfrom:"PRIMARY KEY"`
	Name string
}

func TestBuilder(t *testing.T) {
	queryBuilders := clause.Querybuilders
	userStruct := &User{Id: 1, Name: "usename"}
	// insert
	if f, ok := queryBuilders[clause.INSERT]; ok {
		sql, _ := f("user", []string{"id", "name"})
		if sql != "INSERT INTO user (id,name) " {
			fmt.Println("build insert sql error")
		}
	} else {
		fmt.Println("get insert builder error")
	}

	//values
	if f, ok := queryBuilders[clause.VALUES]; ok {
		fmt.Println("values builder test")
		sql, vars := f([]interface{}{userStruct.Id, userStruct.Name}, []interface{}{2, "name2"})
		fmt.Sprintf("sql:%s  ,vars:%v", sql, vars)
		if sql != "VALUES (?,?), (?,?)" {
			fmt.Println("test values builder error")
		}
	} else {
		fmt.Println("get insert builder error")

	}
	//select
	if f, ok := queryBuilders[clause.SELECT]; ok {
		sql, _ := f("user", []string{"id", "name"})
		if sql != "SELECT id name FROM user" {
			fmt.Println("build select sql error")
		}
	} else {
		fmt.Println("get insert builder error")
	}

	// limit
	if f, ok := queryBuilders[clause.LIMIT]; ok {
		sql, _ := f(2)
		if sql != "LIMIT ?" {
			fmt.Println("build limit sql error")
		}
	} else {
		fmt.Println("get insert builder error")
	}
	//where
	if f, ok := queryBuilders[clause.WHERE]; ok {
		sql, _ := f("id = ?", userStruct.Id)
		if sql != "where id = ?" {
			fmt.Println("build where sql error")
		}
	} else {
		fmt.Println("get insert where error")
	}

	//orderby
	if f, ok := queryBuilders[clause.ORDERBY]; ok {
		sql, _ := f("id", userStruct.Id)
		if sql != "ORDER BY id" {
			fmt.Println("build orderby sql error")
		}
	} else {
		fmt.Println("get orderby builder error")
	}

	//delete
	if f, ok := queryBuilders[clause.DELETE]; ok {
		sql, _ := f("user")
		if sql != "DELETE FROM user" {
			fmt.Println("build delete sql error")
		}
	} else {
		fmt.Println("get insert delte error")
	}

	//drop

	if f, ok := queryBuilders[clause.DROP]; ok {
		sql, _ := f("user")
		if sql != "DROP TABLE user" {
			fmt.Println("build drop sql error")
		}
	} else {
		fmt.Println("get drop builder error")
	}

	//create
	if f, ok := queryBuilders[clause.CREATE]; ok {
		sql, _ := f("user", []string{"id", "name"}, []interface{}{"int", "varchar"})
		if sql != "CREATE TABLE user(  ?  ?, ?  ?, )" {
			fmt.Println("build create sql error")
		}
	} else {
		fmt.Println("get insert builder error")
	}
}
