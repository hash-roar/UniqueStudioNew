package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

type Dialect interface {
	DataType(typ reflect.Value) string
	TestExistSql(tabName string) (string, []interface{})
	GetPosSymbol(i int) string
	GetDialectSql(sql string) string
}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
