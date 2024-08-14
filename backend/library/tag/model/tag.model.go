package model

import (
	"reflect"
)

type TagDefinition struct {
	TagName    string
	TagValue   string
	FieldRef   reflect.Value
	FieldName  string
	FieldType  reflect.Type
	FieldValue interface{}
	IsStruct   bool
	IsSlice    bool
}
