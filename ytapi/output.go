/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"fmt"
	"reflect"
	"errors"
)

////////////////////////////////////////////////////////////////////////////////

// Field specification
type FieldSpec struct {
	Key  string
	Path string
	Type int
}

// Table object
type Table struct {
	colkey []string
	colmap map[string]bool
	fields map[string]FieldSpec
	parts  map[string]string
	rows   []*Values
}

////////////////////////////////////////////////////////////////////////////////

// Returns a new table object
func NewTable() *Table {
	this := &Table{}
	this.colmap = make(map[string]bool)
	this.fields = make(map[string]FieldSpec)
	this.parts = make(map[string]string)
	return this
}

// Creates a new row object and appends it to the table
func (this *Table) NewRow() *Values {
	row := NewValues()
	this.rows = append(this.rows,row)
	return row
}

// Register a part & fields
func (this *Table) RegisterPart(part string, fields []FieldSpec) {
	for _, field := range fields {
		_, ok := this.fields[field.Key]
		if ok {
			panic(fmt.Sprint("Duplicate key '", field.Key, "' for part '", part, "'"))
		}
		this.fields[field.Key] = field
		this.parts[field.Key] = part
	}
}

// Set the default output columns
func (this *Table) SetColumns(columns []string) {
	this.colkey = columns
	for _, key := range columns {
		this.colmap[key] = true
	}
}

// Return number of columns
func (this *Table) NumberOfColumns() int {
	return len(this.colkey)
}

// Return number of rows
func (this *Table) NumberOfRows() int {
	return len(this.rows)
}

// Return parts which are used in the column output
func (this *Table) Parts() []string {

	// from existing columns, determine the parts
	var partmap = make(map[string]bool, len(this.colkey))
	for _, key := range this.colkey {
		value, ok := this.parts[key]
		if ok == false {
			panic(fmt.Sprint("Missing FieldSpec '", key, "'"))
		}
		partmap[value] = true
	}

	// now output part values
	var partvalue = make([]string, 0)
	for key, _ := range partmap {
		partvalue = append(partvalue, key)
	}

	// return the parts
	return partvalue
}

// Append items to the table
func (this *Table) Append(items interface{}) error {
	arrayType := reflect.ValueOf(items)
	if arrayType.Kind() != reflect.Array && arrayType.Kind() != reflect.Slice {
		return errors.New(fmt.Sprint("Append expects array type, got ", arrayType.Kind()))
	}
	for i := 0; i < arrayType.Len(); i++ {
		err := this.appendItem(arrayType.Index(i))
		if err != nil {
			return err
		}
	}
	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Private implementation

func (this *Table) appendItem(item reflect.Value) error {
	// get a new row
//	row := this.NewRow()
	for _, key := range this.colkey {
//		field, ok := this.fields[key]
//		if !ok {
//			return errors.New(fmt.Sprint("Column not specified: '", key, "'"))
//		}

		// deal with pointers to structs as well as structs
		if item.Kind() == reflect.Ptr {
			//row.Set(key, field.value(item.Elem()))
			fmt.Println(key,"=>",item.Elem())
		} else {
			//row.Set(key, field.value(item))
			fmt.Println(key,"=>",item)
		}
	}
	// success
	return nil
}


