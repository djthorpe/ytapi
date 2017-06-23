/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"encoding/csv"

	"github.com/djthorpe/ytapi/util"
	"github.com/olekukonko/tablewriter"
)

////////////////////////////////////////////////////////////////////////////////

const (
	OUTPUT_ASCII = iota
	OUTPUT_CSV
)

////////////////////////////////////////////////////////////////////////////////

type Table struct {
	colkey     []string          // order of registered columns to display
	colmap     map[string]bool   // whether a column exists in the display
	partorder  []string          // order of registered parts
	fields     map[string]*Flag  // field name -> field
	parts      map[string]string // field name -> part
	paths      map[string][]string
	rows       []*Values
	format     int
	infoOutput io.Writer
	dataOutput io.Writer
}

////////////////////////////////////////////////////////////////////////////////

// Returns a new table object
func NewTable() *Table {
	this := &Table{}
	this.colkey = []string{}
	this.colmap = make(map[string]bool)

	this.partorder = make([]string, 0)
	this.fields = make(map[string]*Flag)
	this.parts = make(map[string]string)
	this.paths = make(map[string][]string)

	this.format = OUTPUT_ASCII
	this.infoOutput = os.Stderr
	this.dataOutput = os.Stdout

	return this
}

// Creates a new row object and appends it to the table
func (this *Table) NewRow() *Values {
	row := NewValues()
	this.rows = append(this.rows, row)
	return row
}

// Register a part & fields
func (this *Table) RegisterPart(part string, fields []*Flag) {
	for _, field := range fields {
		_, exists := this.fields[field.Name]
		if exists {
			panic(fmt.Sprint("Duplicate key '", field.Name, "' for part '", part, "'"))
		}

		// generate path if it's not defined
		if field.Path == "" {
			this.paths[field.Name] = strings.Split(generatePath(part, field.Name), "/")
		} else {
			this.paths[field.Name] = strings.Split(field.Path, "/")
		}

		// save parts
		this.fields[field.Name] = field
		this.parts[field.Name] = part
	}
	// append part order
	this.partorder = append(this.partorder, part)
}

// Set the default output columns
func (this *Table) SetColumns(columns []string) error {
	this.colkey = []string{}
	this.colmap = make(map[string]bool)
	for _, key := range columns {
		if err := this.AddFieldOrPart(key); err != nil {
			return err
		}
	}

	// success
	return nil
}

// Add a column or part
func (this *Table) AddFieldOrPart(key string) error {
	// if field, then add the field
	if _, exists := this.fields[key]; exists {
		this.addField(key)
		return nil
	}

	// if snippet, then expand into fields
	fields := this.FieldsForPart(key)
	if len(fields) == 0 {
		return errors.New(fmt.Sprint("Unknown field or part name to add: ", key))
	}

	// add fields
	for _, field := range fields {
		this.addField(field.Name)
	}

	// success
	return nil
}

// Add a field to the output columns
func (this *Table) addField(key string) {
	// remove existing column
	if _, exists := this.colmap[key]; exists {
		this.removeField(key)
	}
	// append column
	this.colkey = append(this.colkey, key)
	this.colmap[key] = true
}

// Remove field from the output columns
func (this *Table) removeField(key string) {
	if _, exists := this.colmap[key]; exists == false {
		return
	}
	// regenerate columns
	j := -1
	for i, field := range this.colkey {
		if field == key {
			j = i
		}
	}
	if j >= 0 {
		this.colkey = append(this.colkey[:j], this.colkey[j+1:]...)
	}
	// remove from column map
	delete(this.colmap, key)
}

// Remove a field or part from the output columns
func (this *Table) RemoveFieldOrPart(key string) error {
	// if field, then add the field
	if _, exists := this.fields[key]; exists {
		this.removeField(key)
		return nil
	}
	// if snippet, then expand into fields
	fields := this.FieldsForPart(key)
	if len(fields) == 0 {
		return errors.New(fmt.Sprint("Unknown field or part name to remove: ", key))
	}
	// remove fields
	for _, field := range fields {
		this.removeField(field.Name)
	}
	// success
	return nil
}

// Return number of columns
func (this *Table) NumberOfColumns() int {
	return len(this.colkey)
}

// Return number of rows
func (this *Table) NumberOfRows() int {
	return len(this.rows)
}

// Return parts which are used in the column output,
// or if 'all' is set to true, return all parts registered in order
func (this *Table) Parts(all bool) []string {
	// if all parts should be returned...
	if all {
		return this.partorder
	}

	// from existing columns, determine the parts
	var partmap = make(map[string]bool, len(this.colkey))
	for _, key := range this.colkey {
		value, ok := this.parts[key]
		if ok == false {
			panic(fmt.Sprint("Missing Flag '", key, "'"))
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

// Return fields for a particular part
func (this *Table) FieldsForPart(part string) []*Flag {
	fields := make([]*Flag, 0)
	for key, value := range this.parts {
		if part != value {
			continue
		}
		fields = append(fields, this.fields[key])
	}
	return fields
}

// Append items to the table
func (this *Table) Append(items interface{}) error {
	arrayType := reflect.ValueOf(items)
	if arrayType.Kind() != reflect.Array && arrayType.Kind() != reflect.Slice {
		return errors.New(fmt.Sprint("Append expects array type, got ", arrayType.Kind()))
	}
	for i := 0; i < arrayType.Len(); i++ {
		item := arrayType.Index(i)
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}
		if item.Kind() == reflect.Struct {
			if err := this.appendStructItem(item); err != nil {
				return err
			}
		} else if item.Kind() == reflect.Array || item.Kind() == reflect.Slice {
			if err := this.appendArrayItem(item); err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprint("ytapi.Append expects array, slice or struct type, got ", item.Kind()))
		}
	}
	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Output methods

func (this *Table) asStringArray(row *Values) []string {
	values := make([]string, this.NumberOfColumns())
	for i, key := range this.colkey {
		if row.IsSet(this.fields[key]) {
			values[i] = row.GetString(this.fields[key])
		} else {
			values[i] = "<nil>"
		}
	}
	return values
}

func (this *Table) dataOutputASCII(io io.Writer) error {
	w := tablewriter.NewWriter(io)
	w.SetHeader(this.colkey)
	w.SetAutoFormatHeaders(false)
	for _, row := range this.rows {
		w.Append(this.asStringArray(row))
	}
	w.Render()
	return nil
}

func (this *Table) dataOutputCSV(io io.Writer) error {
	w := csv.NewWriter(io)
	w.Write(this.colkey)
	for _, row := range this.rows {
		w.Write(this.asStringArray(row))
	}
	w.Flush()
	return nil
}

func (this *Table) DataOutput() error {
	switch this.format {
	case OUTPUT_ASCII:
		return this.dataOutputASCII(this.dataOutput)
	case OUTPUT_CSV:
		return this.dataOutputCSV(this.dataOutput)
	default:
		panic("Unknown output format")
	}
}

// Set output format
func (this *Table) SetDataFormat(handle io.Writer, format int) {
	this.dataOutput = handle
	this.format = format
}

////////////////////////////////////////////////////////////////////////////////
// Logging output

func (this *Table) Info(message string) {
	fmt.Fprintln(this.infoOutput, message)
}

////////////////////////////////////////////////////////////////////////////////
// Private implementation

func generatePath(parts ...string) string {
	for i, part := range parts {
		parts[i] = util.UppercaseFirstLetter(part)
	}
	return strings.Join(parts, "/")
}

func valueForPath(item reflect.Value, field *Flag, path []string) (*Value, error) {
	value := item
	for _, key := range path {
		// check for invalid value
		if value.IsValid() == false {
			return nil, nil
		}
		if value.Kind() != reflect.Struct {
			panic(fmt.Sprint("Non-struct for key '", key, "', kind is ", value.Kind()))
		}
		// Get value
		value = value.FieldByName(key)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
	}
	return NewValue(field, value)
}

func valueForIndex(item reflect.Value, field *Flag, index int) (*Value, error) {
	value := item.Index(index)
	if value.Kind() == reflect.Interface {
		return NewValue(field, value.Elem())
	} else {
		return NewValue(field, value)
	}
}

func (this *Table) appendStructItem(item reflect.Value) error {
	// get a new row
	row := this.NewRow()

	// set row elements
	for _, key := range this.colkey {
		if field, exists := this.fields[key]; exists == false {
			return errors.New(fmt.Sprint("Missing column: '", key, "'"))
		} else if path, exists := this.paths[key]; exists == false {
			return errors.New(fmt.Sprint("Missing column: '", key, "'"))
		} else if value, err := valueForPath(item, field, path); err != nil {
			return err
		} else if value != nil {
			row.Set(value)
		}
	}

	// success
	return nil
}

func (this *Table) appendArrayItem(item reflect.Value) error {
	// get a new row
	row := this.NewRow()

	// set row elements
	for i, key := range this.colkey {
		if field, exists := this.fields[key]; exists == false {
			return errors.New(fmt.Sprint("Missing column: '", key, "'"))
		} else if value, err := valueForIndex(item, field, i); err != nil {
			return err
		} else if value != nil {
			row.Set(value)
		}
	}

	// success
	return nil
}
