package schema_test

import (
	"fmt"
	"lform/dialect"
	"lform/schema"
	"reflect"
	"time"
)

type User struct {
	Id        int    `lform:"PRIMARY KEY"`
	Name      string `lform:"UNIQUE"`
	CreteTime time.Time
}

func SchemaFlatenValue_test() {
	userTest := User{Id: 1, Name: "username", CreteTime: time.Now()}
	dial, _ := dialect.GetDialect("postgres")
	s := schema.Parse(userTest, dial)
	result := s.FlatenValues(userTest)
	if reflect.DeepEqual(result, []interface{}{1, "username", userTest.CreteTime}) {
		fmt.Println("parse time error")
	}
	fmt.Println("SchemaFlatenValue_test success")
}

func SchemaParse_test() {
	userTest := User{Id: 1, Name: "username", CreteTime: time.Now()}
	dial, _ := dialect.GetDialect("postgres")
	s := schema.Parse(userTest, dial)
	if s.Name != "User" {
		fmt.Println("parse table name error")
	}
	if !reflect.DeepEqual(s.FieldNames, []string{"Id", "Name", "CreateTime "}) {
		fmt.Println("parse table fieldname error")
	}
	if !reflect.DeepEqual(s.DbFieldName, []string{"id", "name", "crete_time"}) {
		fmt.Println("parse dbFieldName error")
	}
	fmt.Println("SchemaParse_test success")
}
