/*
  Copyright David Thorpe 2019 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package util

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/olekukonko/tablewriter"
)

////////////////////////////////////////////////////////////////////////////////

type Table struct {
	colkey []string          // order of registered columns to display
	fields map[string]*Field // map a column name to a field (name, path, type)
	rows   []*Row
}

type Row struct {
	values []*Value
}

////////////////////////////////////////////////////////////////////////////////

// NewTable returns a new table object
func NewTable() *Table {
	this := &Table{}
	this.colkey = []string{}
	this.fields = make(map[string]*Field)
	this.rows = make([]*Row, 0)
	return this
}

func (this *Table) RenderText(dev *os.File) {
	table := tablewriter.NewWriter(dev)
	table.SetHeader(this.colkey)

	cells := make([]string, len(this.colkey))
	for _, row := range this.rows {
		// Convert to string
		for i := range cells {
			cells[i] = row.StringAtIndex(i)
		}
		table.Append(cells)
	}

	table.Render()
}

////////////////////////////////////////////////////////////////////////////////

// NewRow returns a new row object, appending it to the end of the table
func (this *Table) NewRow() *Row {
	values := new(Row)
	this.rows = append(this.rows, values)
	return values
}

// StringAtIndex returns string value at row index
func (this *Row) StringAtIndex(i int) string {
	if i >= len(this.values) {
		return "<nil>"
	} else {
		return fmt.Sprint(this.values[i])
	}
}

// Append appends a value onto the row
func (this *Row) Append(value *Value) {
	this.values = append(this.values, value)
}

////////////////////////////////////////////////////////////////////////////////

func (this *Table) AddColumn(field *Field) error {
	if _, exists := this.fields[field.Name]; exists {
		return fmt.Errorf("%v: Column already exists", field.Name)
	} else {
		this.colkey = append(this.colkey, field.Name)
		this.fields[field.Name] = field
		return nil
	}
}

func (this *Table) AddColumnsFrom(item reflect.Value) error {
	if item.Kind() == reflect.Ptr {
		item = item.Elem()
	}
	if item.Kind() != reflect.Struct {
		return fmt.Errorf("AddColumnsFrom requires a Struct value")
	}
	for i := 0; i < item.NumField(); i++ {
		field := item.Type().Field(i)
		name, path := field.Name, field.Name
		type_ := FIELD_NONE
		// Obtain the name of the field from json tag
		if json, exists := field.Tag.Lookup("json"); exists {
			name = json
		}
		// Obtain the field type from the ytapi tag or from the struct
		if field_type, exists := field.Tag.Lookup("ytapi"); exists {
			switch field_type {
			case "datetime":
				type_ = FIELD_DATETIME
			case "string_array":
				type_ = FIELD_STRING_ARRAY
			case "string_map":
				type_ = FIELD_STRING_MAP
			case "seconds":
				type_ = FIELD_SECONDS
			default:
				fmt.Printf("Ignoring field: %v\n", field)
				continue
			}
		} else {
			// Skip any fields which aren't string, uint, int or bool
			switch field.Type.Kind() {
			case reflect.String:
				type_ = FIELD_STRING
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				type_ = FIELD_UINT
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				type_ = FIELD_INT
			case reflect.Bool:
				type_ = FIELD_BOOL
			default:
				fmt.Printf("Ignoring field: %v\n", field)
				continue
			}
		}
		// Add Column
		if err := this.AddColumn(&Field{Name: name, Path: path, Type: type_}); err != nil {
			return err
		}
	}

	// Return success
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// Append item to the table
func (this *Table) Append(items interface{}) error {
	arrayType := reflect.ValueOf(items)
	if arrayType.Kind() != reflect.Array && arrayType.Kind() != reflect.Slice {
		return fmt.Errorf("Append expects array type, got: %v ", arrayType.Kind())
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
			return fmt.Errorf("Append expects array, slice or struct type, got: %v ", item.Kind())
		}
	}
	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (this *Table) appendStructItem(item reflect.Value) error {
	// get a new row
	row := this.NewRow()

	// set row elements
	for _, key := range this.colkey {
		if field, exists := this.fields[key]; exists == false {
			return fmt.Errorf("Missing column: %v", key)
		} else if value, err := valueForPath(item, field); err != nil {
			return err
		} else if value != nil {
			row.Append(value)
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
			return fmt.Errorf("Missing column: %v", key)
		} else if value, err := valueForIndex(item, field, i); err != nil {
			return err
		} else if value != nil {
			row.Append(value)
		}
	}

	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func valueForIndex(item reflect.Value, field *Field, index int) (*Value, error) {
	if index >= item.Len() {
		return nil, fmt.Errorf("Array index out of range")
	}
	value := item.Index(index)
	if value.Kind() == reflect.Interface {
		return NewValue(value.Elem(), field), nil
	} else {
		return NewValue(value, field), nil
	}
}

func valueForPath(item reflect.Value, field *Field) (*Value, error) {
	value := item
	for _, key := range strings.Split(field.Path, "/") {
		if value.IsValid() == false {
			return nil, fmt.Errorf("Invalid path: %v", field.Path)
		}
		if value.Kind() != reflect.Struct {
			panic(fmt.Sprint("Non-struct for key '", key, "', kind is ", value.Kind()))
		}
		value = value.FieldByName(key)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
	}
	if value.IsValid() == false {
		return nil, fmt.Errorf("Invalid path: %v", field.Path)
	} else {
		return NewValue(value, field), nil
	}
}
