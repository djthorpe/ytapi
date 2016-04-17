/*
  Copyright David Thorpe 2015-2016 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytservice

import (
	"io"
	"fmt"
	"encoding/csv"

	"github.com/olekukonko/tablewriter"
)

////////////////////////////////////////////////////////////////////////////////

const (
	FIELD_STRING = iota
)

////////////////////////////////////////////////////////////////////////////////

// Row object
type Row struct {
	values map[string]interface{}
}

// Table object
type Table struct {
	colkey  []string
	colname map[string]string
	rows    []*Row
}

// Field specification
type FieldSpec struct {
	Key  string
	Path string
	Type int
}

////////////////////////////////////////////////////////////////////////////////

// Returns a service object given service account details
func NewTable(columns []string) *Table {
	// Create a 'this' object
	this := &Table{
		colkey:  columns,
		colname: map[string]string{},
		rows:    []*Row{},
	}

	// Set columns names from column keys
	this.colname = make(map[string]string,len(columns))
	for _,key := range(columns) {
		this.colname[key] = key
	}

	// Return table
	return this
}

func (this *Table) NewRow() *Row {
	row := new(Row)

	this.rows = append(this.rows, row)
	row.values = make(map[string]interface{})

	return row
}

////////////////////////////////////////////////////////////////////////////////

// Set the output columns
func (this *Table) SetColumns(columns []string) {
	// TODO
}

// Register output formats
func (this *Table) RegisterPart(part string,fields []FieldSpec) {
	// TODO
}

////////////////////////////////////////////////////////////////////////////////

func (this *Table) NumberOfColumns() int {
	return len(this.colkey)
}

func (this *Table) NumberOfRows() int {
	return len(this.rows)
}

func (this *Table) AppendColumn(key string,name string) {
	// Check for existing
	_,exists := this.colname[key]
	// Set the name
	this.colname[key] = name
	// Append if doesn't exist yet
	if exists == false {
		this.colkey = append(this.colkey,key)
	}
}

func (this *Table) CSV(io io.Writer) error {
	w := csv.NewWriter(io)
	w.Write(this.colkey)
	for _,row := range(this.rows) {
		w.Write(row.asStringArray(this))
	}
	w.Flush()
	return w.Error()
}

func (this *Table) ASCII(io io.Writer) error {
	w := tablewriter.NewWriter(io)
	w.SetHeader(this.colkey)
	for _,row := range(this.rows) {
		w.Append(row.asStringArray(this))
	}
	w.Render()
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (this *Row) asString(key string) string {
	return fmt.Sprint(this.values[key])
}

func (this *Row) asStringArray(table *Table) []string {
	values := make([]string,table.NumberOfColumns())
	for i,key := range(table.colkey) {
		values[i] = this.asString(key)
	}
	return values
}

func (this *Row) SetString(key string, value string) {	
	this.values[key] = value
}
