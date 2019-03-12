package util

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"time"
)

type FlagSet struct {
	flagset *flag.FlagSet
}

var (
	ErrHelpRequested = flag.ErrHelp
)

// NewFlagSet creates a new flagset object
func NewFlagSet(name string) *FlagSet {
	this := new(FlagSet)
	this.flagset = flag.NewFlagSet(name, flag.ContinueOnError)
	this.flagset.SetOutput(ioutil.Discard)
	this.flagset.Name()
	return this
}

func (this *FlagSet) Name() string {
	return this.flagset.Name()
}

func (this *FlagSet) Parse() error {
	return this.flagset.Parse(os.Args[1:])
}

func (this *FlagSet) Args() []string {
	return this.flagset.Args()
}

func (this *FlagSet) NArg() int {
	return this.flagset.NArg()
}

func (this *FlagSet) FlagsForScope(scope ScopeType) []*Value {
	fields := make([]*Value, 0)
	this.flagset.VisitAll(func(f *flag.Flag) {
		if value := f.Value.(*Value); value.f.Scope == scope {
			fields = append(fields, value)
		}
	})
	return fields
}

////////////////////////////////////////////////////////////////////////////////

func (this *FlagSet) String(name, value, usage string, scope ScopeType) error {
	if this.flagset.Lookup(name) != nil {
		return fmt.Errorf("Duplicate flag: %s", name)
	} else if value := NewValue(reflect.ValueOf(value), &Field{Name: name, Description: usage, Type: FIELD_STRING, Scope: scope}); value == nil {
		return fmt.Errorf("Invalid flag: %s", name)
	} else {
		this.flagset.Var(value, name, usage)
	}
	// success
	return nil
}

func (this *FlagSet) Bool(name string, value bool, usage string, scope ScopeType) error {
	if this.flagset.Lookup(name) != nil {
		return fmt.Errorf("Duplicate flag: %s", name)
	} else if value := NewValue(reflect.ValueOf(value), &Field{Name: name, Description: usage, Type: FIELD_BOOL, Scope: scope}); value == nil {
		return fmt.Errorf("Invalid flag: %s", name)
	} else {
		this.flagset.Var(value, name, usage)
	}
	// success
	return nil
}

func (this *FlagSet) Uint(name string, value uint, usage string, scope ScopeType) error {
	if this.flagset.Lookup(name) != nil {
		return fmt.Errorf("Duplicate flag: %s", name)
	} else if value := NewValue(reflect.ValueOf(value), &Field{Name: name, Description: usage, Type: FIELD_UINT, Scope: scope}); value == nil {
		return fmt.Errorf("Invalid flag: %s", name)
	} else {
		this.flagset.Var(value, name, usage)
	}
	// success
	return nil
}

func (this *FlagSet) Duration(name string, value time.Duration, usage string, scope ScopeType) error {
	if this.flagset.Lookup(name) != nil {
		return fmt.Errorf("Duplicate flag: %s", name)
	} else if value := NewValue(reflect.ValueOf(value), &Field{Name: name, Description: usage, Type: FIELD_DURATION, Scope: scope}); value == nil {
		return fmt.Errorf("Invalid flag: %s", name)
	} else {
		this.flagset.Var(value, name, usage)
	}
	// success
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (this *FlagSet) GetString(name string) string {
	if flag := this.flagset.Lookup(name); flag == nil {
		return ""
	} else {
		return flag.Value.(*Value).String()
	}
}

func (this *FlagSet) GetBool(name string) bool {
	if flag := this.flagset.Lookup(name); flag == nil {
		return false
	} else {
		return flag.Value.(*Value).Bool()
	}
}

func (this *FlagSet) GetDuration(name string) time.Duration {
	if flag := this.flagset.Lookup(name); flag == nil {
		return 0
	} else {
		return flag.Value.(*Value).Duration()
	}
}
