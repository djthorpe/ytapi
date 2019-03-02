/*
  Copyright David Thorpe 2019 All Rights Reserved
  Please see file LICENSE for information on distribution, etc
*/
package util

import (
	"fmt"

	// Frameworks
	"github.com/djthorpe/ytapi/brightcoveapi"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type BrightcoveCall func(*brightcoveapi.Client, []string) error

type Command struct {
	Name        string
	Description string
	Brightcove  BrightcoveCall
}

////////////////////////////////////////////////////////////////////////////////
// EXEC

func (this *Command) ExecBrightcove(client *brightcoveapi.Client, args []string) error {
	if this.Brightcove != nil {
		return this.Brightcove(client, args)
	} else {
		return fmt.Errorf("%v: Undefined brightcove call", this.Name)
	}
}
