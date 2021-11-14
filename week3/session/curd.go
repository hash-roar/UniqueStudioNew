package session

import (
	"errors"
	"fmt"
	"lform/clause"
	"reflect"
)

// &struct1,&struct2

func (s *Session) Insert(values ...interface{}) (int64, error) {
	flatValues := make([]interface{}, 0)
	for _, value := range values {
		table := s.Model(value).GetTable()
		s.clause.BuildSqlSlice(clause.INSERT, table.Name, table.FieldNames)
		flatValues = append(flatValues, table.FlatenValues(value))
	}
	s.clause.BuildSqlSlice(clause.VALUES, flatValues...)
	sql, vars := s.clause.BuildSql(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()

}

// value &[]struct
func (s *Session) FindAll(value interface{}) error {
	structValue := reflect.Indirect(reflect.ValueOf(value))
	structType := structValue.Type().Elem()
	table := s.Model(reflect.New(structType).Elem().Interface()).GetTable()

	s.clause.BuildSqlSlice(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.BuildSql(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	results, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for results.Next() {
		resultValueSlice := reflect.New(structType).Elem()
		var vars []interface{}
		for _, name := range table.FieldNames {
			vars = append(vars, resultValueSlice.FieldByName(name).Addr().Interface())
		}
		if err := results.Scan(vars...); err != nil {
			return err
		}
		structValue.Set(reflect.Append(structValue, resultValueSlice))
	}
	return results.Close()
}

// & struct
func (s *Session) FindOne(value interface{}) error {
	var valueTemp []interface{}
	valueTemp = append(valueTemp, value)
	err := s.Limit(1).FindAll(&valueTemp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	reflect.Indirect(reflect.ValueOf(value)).Set(reflect.ValueOf(valueTemp[0]))
	return nil
}

func (s *Session) Update(values ...interface{}) (int64, error) {
	m, ok := values[0].(map[string]interface{})
	if !ok {
		m = make(map[string]interface{})
		for i := 0; i < len(values); i += 2 {
			m[values[i].(string)] = values[i+1]
		}
	}
	s.clause.BuildSqlSlice(clause.UPDATE, s.GetTable().Name, m)
	sql, vars := s.clause.BuildSql(clause.UPDATE, clause.WHERE)
	results, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return results.RowsAffected()
}

func (s *Session) Delete() (int64, error) {
	s.clause.BuildSqlSlice(clause.DELETE, s.GetTable().Name)
	sql, vars := s.clause.BuildSql(clause.DELETE, clause.WHERE)
	results, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return results.RowsAffected()
}

func (s *Session) Limit(num int) *Session {
	s.clause.BuildSqlSlice(clause.LIMIT, num)
	return s
}

func (s *Session) where(conditions string, args ...interface{}) *Session {
	var vars []interface{}
	vars = append(vars, conditions)
	vars = append(vars, args...)
	s.clause.BuildSqlSlice(clause.WHERE, vars...)
	return s

}

func (s *Session) OrderBy(cond string) *Session {
	s.clause.BuildSqlSlice(clause.ORDERBY, cond)
	return s
}

func (s *Session) First(value interface{}) error {
	structValue := reflect.Indirect(reflect.ValueOf(value))
	structTempSlice := reflect.New(reflect.SliceOf(structValue.Type())).Elem()
	if err := s.Limit(1).FindAll(structTempSlice.Addr().Interface()); err != nil {
		fmt.Println(err)
		return err
	}
	if structTempSlice.Len() == 0 {
		return errors.New("404")
	}
	structValue.Set(structTempSlice.Index(0))
	return nil
}
