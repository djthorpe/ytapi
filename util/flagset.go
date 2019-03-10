package util

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type FlagSet struct {
	flagset *flag.FlagSet
}

type Flag struct {
	Name  string
	Usage string
}

// NewFlagSet creates a new flagset object
func NewFlagSet() *FlagSet {
	this := new(FlagSet)
	this.flagset = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	this.flagset.SetOutput(ioutil.Discard)
	return this
}

// AddFlag adds a flag to the flagset
func (this *FlagSet) AddFlag(flag *Flag) error {
	// check for flag name clash
	if this.flagset.Lookup(flag.Name) != nil {
		return fmt.Errorf("Duplicate flag: ", flag.Name)
	}

	// set flag
	this.flagset.Var(nil, flag.Name, flag.Usage)

	// success
	return nil
}
