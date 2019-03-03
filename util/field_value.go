/*
  Copyright David Thorpe 2019 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package util

import (
	"fmt"
	"reflect"
)

////////////////////////////////////////////////////////////////////////////////

type FieldType uint

type Value struct {
	v reflect.Value
	f *Field
}

type Field struct {
	Name string
	Path string
	Type FieldType
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

////////////////////////////////////////////////////////////////////////////////

// NewValue returns an empty value
func NewValue(v reflect.Value, f *Field) *Value {
	return &Value{v, f}
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
