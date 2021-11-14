package clause

import "strings"

type SqlType int

const (
	INSERT SqlType = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
	UPDATE
	DELETE
	DROP
	CREATE
)

type Clause struct {
	sql     map[SqlType]string
	sqlVars map[SqlType][]interface{}
}

func (c *Clause) BuildSqlSlice(typ SqlType, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[SqlType]string)
		c.sqlVars = make(map[SqlType][]interface{})
	}
	sql, vars := Querybuilders[typ](vars...)
	c.sql[typ] = sql
	c.sqlVars[typ] = vars
}

func (c *Clause) BuildSql(sqlOrders ...SqlType) (string, []interface{}) {
	var sqls []string
	var vars []interface{}
	for _, order := range sqlOrders {
		if sqlSlice, ok := c.sql[order]; ok {
			sqls = append(sqls, sqlSlice)
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}
