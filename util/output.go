/*
  Copyright David Thorpe 2019 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package util

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

////////////////////////////////////////////////////////////////////////////////

type Table struct {
	colkey []string        // order of registered columns to display
	colmap map[string]bool // whether a column exists in the display
}

////////////////////////////////////////////////////////////////////////////////

// Returns a new table object
func NewTable() *Table {
	this := &Table{}
	this.colkey = []string{}
	this.colmap = make(map[string]bool)
	return this
}

func (this *Table) RenderText(dev *os.File) {
	table := tablewriter.NewWriter(dev)
	table.SetHeader(this.colkey)
	table.Render()
}

////////////////////////////////////////////////////////////////////////////////

func (this *Table) AddColumn(name string) error {
	if _, exists := this.colmap[name]; exists {
		return fmt.Errorf("%v: Column already exists", name)
	} else {
		this.colkey = append(this.colkey, name)
		this.colmap[name] = true
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////

func (this *Table) Append(obj interface{}) error {
	// TODO
	return nil
}
