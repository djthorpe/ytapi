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
	colkey     []string
	partorder  []string
	colmap     map[string]bool
	fields     map[string]*Flag
	parts      map[string]string
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
	this.partorder = make([]string, 0)
	this.colmap = make(map[string]bool)
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
func (this *Table) SetColumns(columns []string) {
	this.colkey = columns
	for _, key := range columns {
		this.colmap[key] = true
	}
}

// Add a field or part to the output columns
func (this *Table) AddColumn(name string) error {
	_, exists := this.colmap[name]
	if exists == true {
		// column already exists
		return nil
	}

	// check for column name
	if _, exists := this.fields[name]; exists {
		this.colkey = append(this.colkey, name)
		this.colmap[name] = true
		return nil
	}

	// check for snippet name
	fields := this.FieldsForPart(name)
	if len(fields) == 0 {
		return errors.New(fmt.Sprint("Unknown field or part name: ", name))
	}

	// add snippet columns
	for _, field := range fields {
		if err := this.RemoveColumn(field.Name); err != nil {
			return err
		}
		if err := this.AddColumn(field.Name); err != nil {
			return err
		}
	}
	return nil
}

// Remove a field or part to the output columns
func (this *Table) RemoveColumn(name string) error {
	_, exists := this.colmap[name]
	if exists == false {
		// column does not exist
		return nil
	}

	fmt.Printf("REMOVE ", name)
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
		err := this.appendItem(arrayType.Index(i))
		if err != nil {
			return err
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
		values[i] = row.GetString(this.fields[key])
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
    switch(this.format) {
        case OUTPUT_ASCII:
            return this.dataOutputASCII(this.dataOutput)
        case OUTPUT_CSV:
            return this.dataOutputCSV(this.dataOutput)
        default:
            panic("Unknown output format")
    }
}

// Set output format
func (this *Table) SetDataFormat(handle io.Writer,format int) {
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

func (this *Table) appendItem(item reflect.Value) error {
	// get a new row
	row := this.NewRow()

	// set row elements
	for _, key := range this.colkey {
		var value *Value
		var err error

		// get the flag for the row
		field, exists := this.fields[key]
		if !exists {
			return errors.New(fmt.Sprint("Missing column: '", key, "'"))
		}
		path, exists := this.paths[key]
		if !exists {
			return errors.New(fmt.Sprint("Missing column: '", key, "'"))
		}

		// deal with pointers to items as well as items
		if item.Kind() == reflect.Ptr {
			value, err = valueForPath(item.Elem(), field, path)
		} else {
			value, err = valueForPath(item, field, path)
		}
		if err != nil {
			return err
		}
		if value != nil {
			row.Set(value)
		}
	}

	// success
	return nil
}
