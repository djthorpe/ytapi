/*
  Copyright David Thorpe 2019 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package util

import (
	"errors"
	"fmt"
	"reflect"
)

////////////////////////////////////////////////////////////////////////////////

type FieldType uint
type ScopeType uint

type Value struct {
	v reflect.Value
	f *Field
}

type Field struct {
	Name  string
	Path  string
	Type  FieldType
	Scope ScopeType
}

////////////////////////////////////////////////////////////////////////////////

const (
	FIELD_NONE FieldType = iota
	FIELD_UINT
	FIELD_INT
	FIELD_BOOL
	FIELD_STRING
	FIELD_STRING_ARRAY
	FIELD_STRING_MAP
	FIELD_DATETIME
	FIELD_SECONDS
)

const (
	SCOPE_NONE ScopeType = iota
	SCOPE_GLOBAL
	SCOPE_OPTIONAL
	SCOPE_REQUIRED
)

////////////////////////////////////////////////////////////////////////////////

// NewValue returns an empty value
func NewValue(v reflect.Value, f *Field) *Value {
	return &Value{v, f}
}

func (this *Value) Name() string {
	if this.f != nil {
		return this.f.Name
	} else {
		return ""
	}
}

func (this *Value) Type() FieldType {
	if this.f != nil {
		return this.f.Type
	} else {
		return FIELD_NONE
	}
}

func (this *Value) String() string {
	switch this.v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprint(this.Uint())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprint(this.Int())
	case reflect.Bool:
		return fmt.Sprint(this.Bool())
	default:
		return this.v.String()
	}
}

func (this *Value) Bool() bool {
	return this.v.Bool()
}

func (this *Value) Uint() uint64 {
	return this.v.Uint()
}

func (this *Value) Int() int64 {
	return this.v.Int()
}

func (this *Value) Set(value string) error {
	return errors.New("Not implemented")
}

////////////////////////////////////////////////////////////////////////////////

func (t FieldType) String() string {
	switch t {
	case FIELD_UINT:
		return "<uint>"
	case FIELD_INT:
		return "<int>"
	case FIELD_BOOL:
		return "<bool>"
	case FIELD_STRING:
		return "<string>"
	case FIELD_STRING_ARRAY:
		return "<string>,<string>,..."
	case FIELD_STRING_MAP:
		return "<stringmap>"
	case FIELD_DATETIME:
		return "<datetime>"
	case FIELD_SECONDS:
		return "<seconds>"
	default:
		return "<value>"
	}
}
