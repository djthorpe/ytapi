/*
  Copyright David Thorpe 2015-2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"io"
	"os"
	"fmt"
	"strings"
	"errors"
	"encoding/csv"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type Input struct {
	handle io.ReadCloser
	format InputFormat
	fields []string
	records []map[string]string
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

// Returns a new table object
func NewInput() *Input {
	this := &Input{}

	// Set defaults
	this.handle = os.Stdin
	this.format = INPUT_CSV
	this.fields = nil
	this.records = make([]map[string]string,0)

	// Success
	return this
}

// Close
func (this *Input) Close() {
	if this.handle != os.Stdin {
		this.handle.Close()
		this.handle = os.Stdin
	}
}

// Set Incoming data source
func (this *Input) SetDataSource(handle io.ReadCloser, format InputFormat) {
	if this.handle != os.Stdin {
		this.handle.Close()
	}
	this.handle = handle
	this.format = format
}

// Read everything
func (this *Input) ReadAll() error {
	var fields []string
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

		if fields == nil {
			// If fields is empty, then we create the field names
			fields = this.setFieldNames(record)
		} else if err := this.append(fields,record); err != nil {
			return err
		}
	}
	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *Input) setFieldNames(fields []string) []string {
	fieldNames := make([]string,len(fields))
	for i, field := range fields {
		field = strings.TrimSpace(strings.ToLower(field))
		field = strings.Replace(field," ","_",-1)
		fieldNames[i] = field
	}
	return fieldNames
}

func (this *Input) append(fields []string,record []string) error {
	if len(fields) != len(record) {
		return ErrFieldMismatch
	}
	fmt.Println(fields,"=>",record)
	return nil
}

func (this *Input) isComment(record []string) bool {
	if len(record)==0 {
		return true
	}
	if strings.HasPrefix(record[0],"#") {
		return true
	}
	if strings.HasPrefix(record[0],"//") {
		return true
	}
	return false
}

