/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"fmt"
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

// Register a part
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

