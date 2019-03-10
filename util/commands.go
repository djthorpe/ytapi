/*
  Copyright David Thorpe 2019 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package util

import (
	"fmt"
	"os"

	// Frameworks
	"github.com/djthorpe/ytapi/brightcoveapi"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type BrightcoveCall func(*brightcoveapi.Client, *Table, []string) error
type FormatCall func(*Table) error

type Command struct {
	Name          string
	Description   string
	Usage         string
	OptionalFlags []string
	RequiredFlags []string
	Format        FormatCall
	Brightcove    BrightcoveCall
}

////////////////////////////////////////////////////////////////////////////////
// EXEC

func (this *Command) ExecBrightcove(client *brightcoveapi.Client, args []string) error {
	output := NewTable()
	if this.Brightcove == nil {
		return fmt.Errorf("%v: Undefined brightcove call", this.Name)
	} else if this.Format == nil {
		return fmt.Errorf("%v: Undefined format call", this.Name)
	} else if err := this.Format(output); err != nil {
		return fmt.Errorf("%v: %v", this.Name, err)
	} else if err := this.Brightcove(client, output, args); err != nil {
		return fmt.Errorf("%v: %v", this.Name, err)
	} else {
		output.RenderText(os.Stdout)
	}

	// Success
	return nil
}
