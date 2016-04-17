/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytservice

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/olekukonko/tablewriter"
)

//import "log"

////////////////////////////////////////////////////////////////////////////////

const (
	FIELD_STRING = iota
	FIELD_DATETIME
	FIELD_NUMBER
	FIELD_BOOLEAN
)

////////////////////////////////////////////////////////////////////////////////

var pathcache map[string][]string = make(map[string][]string)

////////////////////////////////////////////////////////////////////////////////

// Value object
type CellValue struct {
	StringValue string
}

// Row object
type Row struct {
	values map[string]CellValue
}

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
	rows   []*Row
}

////////////////////////////////////////////////////////////////////////////////

// Returns a service object given service account details
func NewTable() *Table {
	// Create a 'this' object
	this := &Table{}

	// Set columns names from column keys
	this.colmap = make(map[string]bool)
	this.fields = make(map[string]FieldSpec)
	this.parts = make(map[string]string)

	// Return table
	return this
}

func (this *Table) NewRow() *Row {
	row := new(Row)

	this.rows = append(this.rows, row)
	row.values = make(map[string]CellValue)

	return row
}

////////////////////////////////////////////////////////////////////////////////

// Set the output columns
func (this *Table) SetColumns(columns []string) {
	this.colkey = columns
	for _, key := range columns {
		this.colmap[key] = true
	}
}

// Add a column
func (this *Table) AddColumn(key string) {
	// Check for existing
	_, exists := this.colmap[key]
	// Append if doesn't exist yet
	if exists == false {
		this.colkey = append(this.colkey, key)
		this.colmap[key] = true
	}
}

// Register output formats
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

////////////////////////////////////////////////////////////////////////////////

func (this *Table) NumberOfColumns() int {
	return len(this.colkey)
}

func (this *Table) NumberOfRows() int {
	return len(this.rows)
}

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

func (this *Table) CSV(io io.Writer) error {
	w := csv.NewWriter(io)
	w.Write(this.colkey)
	for _, row := range this.rows {
		w.Write(row.asStringArray(this))
	}
	w.Flush()
	return w.Error()
}

func (this *Table) ASCII(io io.Writer) error {
	w := tablewriter.NewWriter(io)
	w.SetHeader(this.colkey)
	w.SetAutoFormatHeaders(false)
	for _, row := range this.rows {
		w.Append(row.asStringArray(this))
	}
	w.Render()
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// this function returns an array of items on the path, which are split by
// forward slashes. Uses a path cache so that it can quickly return the
// split path without having to split the string again
func splitPath(key string) []string {
	if split, ok := pathcache[key]; ok {
		return split
	}
	pathcache[key] = strings.Split(key, "/")
	return splitPath(key)
}

// Returns a CellValue given an struct
func (this *FieldSpec) value(item reflect.Value) CellValue {
	value := item
	for _, key := range splitPath(this.Path) {
		if value.Kind() != reflect.Struct {
			panic(fmt.Sprint("Non-struct for key '", key, "', kind is ", value.Kind()))
		}
		// Get value
		value = value.FieldByName(key)
		// Deal with pointers
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		// TODO
	}
	return CellValue{StringValue: fmt.Sprint(value)}
}

////////////////////////////////////////////////////////////////////////////////

func (this *Table) appendItem(item reflect.Value) error {
	// get a new row
	row := this.NewRow()
	for _, key := range this.colkey {
		field, ok := this.fields[key]
		if !ok {
			return errors.New(fmt.Sprint("Unable to determine spec for '", key, "'"))
		}

		// deal with pointers to structs as well as structs
		if item.Kind() == reflect.Ptr {
			row.SetCell(key, field.value(item.Elem()))
		} else {
			row.SetCell(key, field.value(item))
		}
	}
	// success
	return nil
}

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

func (this *Table) AddColumnsForPart(part string) error {
	var foundPart bool
	for key, value := range this.parts {
		if part == value {
			foundPart = true
			this.AddColumn(key)
		}
	}
	if foundPart == false {
		return errors.New(fmt.Sprint("No such part: '", part, "'"))
	}
	// success
	return nil
}

func (this *Table) RemoveColumnsForPart(part string) error {
	fmt.Print("Remove ", part)
	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (this *Row) asString(key string) string {
	return fmt.Sprint(this.values[key].StringValue)
}

func (this *Row) asStringArray(table *Table) []string {
	values := make([]string, table.NumberOfColumns())
	for i, key := range table.colkey {
		values[i] = this.asString(key)
	}
	return values
}

func (this *Row) SetString(key string, value string) {
	this.values[key] = CellValue{StringValue: value}
}

func (this *Row) SetCell(key string, value CellValue) {
	this.values[key] = value
}

////////////////////////////////////////////////////////////////////////////////
