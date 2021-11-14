package dialect

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

type postrgesql struct{}

func init() {
	RegisterDialect("postgres", &postrgesql{})
}

func (s *postrgesql) DataType(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "text"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Int32, reflect.Int:
		return "integer"
	case reflect.Int16:
		return "smallint"
	case reflect.Float32, reflect.Float64:
		return "double precision"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "timestamp with time zone"
		}
	}
	panic("invalid type")
}

func (s *postrgesql) GetPosSymbol(i int) string {
	return "$" + strconv.Itoa(i)
}

func (s *postrgesql) GetDialectSql(sql string) string {
	var sqlBuilder strings.Builder
	var symbolIndex int = 1
	for i := 0; i < len(sql); i++ {
		if sql[i] == '?' {
			sqlBuilder.WriteString(s.GetPosSymbol(symbolIndex))
			symbolIndex++
		} else {
			sqlBuilder.WriteByte(sql[i])
		}
	}
	return sqlBuilder.String()
}

func (s *postrgesql) TestExistSql(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "select count(*) from pg_class where relname = ?", args
}
