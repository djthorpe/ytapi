package ytapi

/*
  Copyright David Thorpe 2015-2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type Input struct {
	handle io.ReadCloser
	format InputFormat
	table  *Table
}

type InputFormat int

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	INPUT_CSV InputFormat = iota
)

////////////////////////////////////////////////////////////////////////////////
// VARIABLES

var (
	ErrFieldMismatch = errors.New("Number of fields does not match")
)

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// NewInput returns a new table object
func NewInput() *Input {
	this := &Input{}

	// Set defaults
	this.handle = os.Stdin
	this.format = INPUT_CSV
	this.table = nil

	// Success
	return this
}

// Close closes the input file handle
func (this *Input) Close() {
	if this.handle != os.Stdin {
		this.handle.Close()
		this.handle = os.Stdin
		this.table = nil
	}
}

// SetDataSource sets Incoming data source
func (this *Input) SetDataSource(handle io.ReadCloser, format InputFormat) {
	if this.handle != os.Stdin {
		this.handle.Close()
	}
	this.handle = handle
	this.format = format
}

// ReadAll will read everything
func (this *Input) ReadAll() error {
	reader := csv.NewReader(this.handle)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		} else if this.isComment(record) {
			continue
		}

		if this.table == nil {
			// Create the field names
			this.table = NewTable()
			this.table.SetColumns(this.setFieldNames(record))
		} else if err := this.append(record); err != nil {
			return err
		}
	}
	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *Input) setFieldNames(fields []string) []string {
	fieldNames := make([]string, len(fields))
	for i, field := range fields {
		field = strings.TrimSpace(strings.ToLower(field))
		field = strings.Replace(field, " ", "_", -1)
		fieldNames[i] = field
	}
	return fieldNames
}

func (this *Input) append(record []string) error {
	values := this.table.NewRow()
	fmt.Println(values, "=>", record)
	return nil
}

func (this *Input) isComment(record []string) bool {
	if len(record) == 0 {
		return true
	}
	if strings.HasPrefix(record[0], "#") {
		return true
	}
	if strings.HasPrefix(record[0], "//") {
		return true
	}
	return false
}
