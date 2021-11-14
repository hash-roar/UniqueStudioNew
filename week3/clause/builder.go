package clause

import (
	"fmt"
	"strings"
)

// SELECT col1, col2, ...
//     FROM table_name
//     WHERE [ conditions ]
//     GROUP BY col1
//     HAVING [ conditions ]

type queryBuilder func(values ...interface{}) (string, []interface{})

var Querybuilders map[SqlType]queryBuilder

// var dial dialect.Dialect

func init() {
	Querybuilders = make(map[SqlType]queryBuilder)
	Querybuilders[INSERT] = _insert_
	Querybuilders[VALUES] = _values_
	Querybuilders[SELECT] = _select_
	Querybuilders[LIMIT] = _limit_
	Querybuilders[WHERE] = _where_
	Querybuilders[ORDERBY] = _ordeby_
	Querybuilders[UPDATE] = _update_
	Querybuilders[DELETE] = _delete_
	Querybuilders[DROP] = _drop_
	Querybuilders[CREATE] = _create_
}

// func InitDialect(d dialect.Dialect) {
// 	dial = d
// }

// insert into tablename ('name')

func _insert_(values ...interface{}) (sqlSlice string, vars []interface{}) {
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	sqlSlice = fmt.Sprintf("INSERT INTO %s (%v) ", tableName, fields)
	vars = []interface{}{}
	return
}

func _values_(values ...interface{}) (sqlSlice string, vars []interface{}) {
	var strTemp string
	var sqlTemp strings.Builder
	sqlTemp.WriteString("VALUES ")
	for i, value := range values {
		v := value.([]interface{})
		if strTemp == "" {
			var varsTemp []string
			for i := 0; i < len(v); i++ {
				varsTemp = append(varsTemp, "?")
			}
			strTemp = strings.Join(varsTemp, ",")
		}
		sqlTemp.WriteString(fmt.Sprintf("(%v)", strTemp))
		if i != len(values)-1 {
			sqlTemp.WriteString(", ")
		}
		vars = append(vars, v...)
	}
	sqlSlice = sqlTemp.String()
	return
}

func _select_(values ...interface{}) (sqlSlice string, vars []interface{}) {
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), []interface{}{}
}

func _limit_(values ...interface{}) (sqlSlice string, vars []interface{}) {
	return "LIMIT ?", values
}

func _where_(values ...interface{}) (sqlSlice string, vars []interface{}) {
	conditions, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", conditions), vars
}

func _ordeby_(values ...interface{}) (sqlSlice string, vars []interface{}) {
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}

func _update_(values ...interface{}) (sqlSlice string, vars []interface{}) {
	tableName := values[0]
	valueMap := values[1].(map[string]interface{})
	var (
		varsTemp []interface{}
		keys     []string
	)
	for k, v := range valueMap {
		vars = append(vars, v)
		keys = append(keys, k+" = ?")
	}
	return fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(keys, ", ")), varsTemp
}

func _delete_(values ...interface{}) (sqlSlice string, vars []interface{}) {
	return fmt.Sprintf("DELETE FROM %s", values[0]), []interface{}{}
}

func _drop_(values ...interface{}) (sqlSlice string, vars []interface{}) {
	tableName := values[0]
	return fmt.Sprintf("DROP TABLE %s", tableName), []interface{}{}
}

// tableName   dbFieldName []string  fieldType[]string
func _create_(values ...interface{}) (sqlSlice string, vars []interface{}) {
	var buf strings.Builder
	buf.WriteString("CREATE TABLE ?( ")
	vars = append(vars, values[0]) //tableName
	dbFieldNames := values[1].([]string)
	for i, typeName := range values[2].([]string) {
		buf.WriteString(" ?  ?,")
		vars = append(vars, dbFieldNames[i], typeName)
	}
	buf.WriteString(" )")
	return buf.String(), vars
}
