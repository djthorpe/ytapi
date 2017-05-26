/*
  Copyright David Thorpe 2015-2017 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package ytapi

import (
	"io"
	"os"
	"errors"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type Input struct {
	handle io.ReadCloser
	format InputFormat
}

type InputFormat int

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	INPUT_CSV InputFormat = iota
)

////////////////////////////////////////////////////////////////////////////////

// Returns a new table object
func NewInput() *Input {
	this := &Input{}

	// Set defaults
	this.handle = os.Stdin
	this.format = INPUT_CSV

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
	return errors.New("NOT IMPLEENTED")
}


