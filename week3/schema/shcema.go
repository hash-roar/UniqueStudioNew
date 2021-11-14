package schema

import (
	"lform/dialect"
	"reflect"
	"strings"
)

type Field struct {
	Name          string
	Type          string
	Tag           string
	PrimaryKey    bool
	AutoIncrement bool
	NotNull       bool
	Unique        bool
}

type Schema struct {
	Model       interface{}
	Name        string
	Fields      []*Field
	FieldNames  []string
	DbFieldName []string
	fieldMap    map[string]*Field
}

func parseName(fieldName string) (dbFieldName string) {
	if fieldName == "" {
		return ""
	}
	var (
		buf                            strings.Builder
		lastCase, nextCase, nextNumber bool // upper case == true
		curCase                        = fieldName[0] <= 'Z' && fieldName[0] >= 'A'
	)
	for i, v := range fieldName[:len(fieldName)-1] {
		nextCase = fieldName[i+1] <= 'Z' && fieldName[i+1] >= 'A'
		nextNumber = fieldName[i+1] >= '0' && fieldName[i+1] <= '9'
		if curCase {
			if lastCase && (nextCase || nextNumber) {
				buf.WriteRune(v + 32)
			} else {
				if i > 0 && fieldName[i-1] != '_' && fieldName[i+1] != '_' {
					buf.WriteByte('_')
				}
				buf.WriteRune(v + 32)
			}
		} else {
			buf.WriteRune(v)
		}
		lastCase = curCase
		curCase = nextCase
	}
	if curCase {
		if !lastCase && len(fieldName) > 1 {
			buf.WriteByte('_')
		}
		buf.WriteByte(fieldName[len(fieldName)-1] + 32)
	} else {
		buf.WriteByte(fieldName[len(fieldName)-1])
	}
	dbFieldName = buf.String()
	return
}

func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name]
}

//把结构体展开成对应平铺值
func (s *Schema) FlatenValues(values interface{}) []interface{} {
	valuesTemp := reflect.Indirect(reflect.ValueOf(values))
	var result []interface{}
	for _, field := range s.Fields {
		result = append(result, valuesTemp.FieldByName(field.Name).Interface())
	}
	return result
}

func Parse(des interface{}, dial dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(des)).Type()
	schema := &Schema{Model: des,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field)}
	for i := 0; i < modelType.NumField(); i++ {
		structField := modelType.Field(i)
		if !structField.Anonymous {
			fieldTemp := &Field{Name: structField.Name,
				Type: dial.DataType(reflect.Indirect(reflect.New(structField.Type)))}

			if tagValue, ok := structField.Tag.Lookup("lform"); ok {
				fieldTemp.Tag = tagValue
				if strings.Contains(tagValue, "PRIMARY KEY") {
					fieldTemp.PrimaryKey = true
				}
				if strings.Contains(tagValue, "NOT NULL") {
					fieldTemp.NotNull = true
				}
				if strings.Contains(tagValue, "UNIQUE") {
					fieldTemp.Unique = true
				}
				if strings.Contains(tagValue, "AUTOINCREMENT") {
					fieldTemp.AutoIncrement = true
				}
			}
			schema.Fields = append(schema.Fields, fieldTemp)
			schema.FieldNames = append(schema.FieldNames, structField.Name)
			schema.DbFieldName = append(schema.DbFieldName, parseName(structField.Name))
			schema.fieldMap[structField.Name] = fieldTemp
		}

	}
	return schema
}
