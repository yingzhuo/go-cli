package cli

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
	"time"
)

// Flag represents the state of a flag
type Flag struct {
	Name        string // name as it appears on command line
	Usage       string // help message
	Placeholder string // placeholder in usage
	Hidden      bool   // allow flags to be hidden from help/usage text

	IsBool        bool   // if the flag is bool value
	DefValue      string // default value (as text); for usage message
	NoOptDefValue string // default value (as text); if the flag is on the command line without any options
	EnvVar        string // default value load from environ

	Value interface{} // returns final value

	wrapValue Value // returns final value, wrapped Flag.Value
	visited   bool  // If the user set the value
}

// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
type Value interface {
	String() string
	Set(string) error
}

func (f *Flag) initialize() {
	if f.Value != nil {
		switch val := f.Value.(type) {
		case *bool:
			f.IsBool = true
			f.wrapValue = &boolValue{val}
		case *string:
			f.wrapValue = &stringValue{val}
		case *[]string:
			f.wrapValue = &stringSliceValue{val}
		case *int:
			f.wrapValue = &intValue{val}
		case *[]int:
			f.wrapValue = &intSliceValue{val}
		case *int8:
			f.wrapValue = &int8Value{val}
		case *int16:
			f.wrapValue = &int16Value{val}
		case *int32:
			f.wrapValue = &int32Value{val}
		case *int64:
			f.wrapValue = &int64Value{val}
		case *uint:
			f.wrapValue = &uintValue{val}
		case *[]uint:
			f.wrapValue = &uintSliceValue{val}
		case *uint8:
			f.wrapValue = &uint8Value{val}
		case *uint16:
			f.wrapValue = &uint16Value{val}
		case *uint32:
			f.wrapValue = &uint32Value{val}
		case *uint64:
			f.wrapValue = &uint64Value{val}
		case *float32:
			f.wrapValue = &float32Value{val}
		case *float64:
			f.wrapValue = &float64Value{val}
		case *[]float64:
			f.wrapValue = &float64SliceValue{val}
		case *time.Time:
			f.wrapValue = &timeValue{val}
		case *time.Duration:
			f.wrapValue = &timeDurationValue{val}
		case *time.Location:
			f.wrapValue = &timeLocationValue{val}
		case *net.IP:
			f.wrapValue = &ipValue{val}
		case *[]net.IP:
			f.wrapValue = &ipSliceValue{val}
		case *net.IPMask:
			f.wrapValue = &ipMaskValue{val}
		case *net.IPNet:
			f.wrapValue = &ipNetValue{val}
		case *[]net.IPNet:
			f.wrapValue = &ipNetSliceValue{val}
		case *url.URL:
			f.wrapValue = &urlValue{val}
		case *[]url.URL:
			f.wrapValue = &urlSliceValue{val}
		case *os.FileMode:
			f.wrapValue = &fileModeValue{val}
		case Value:
			f.wrapValue = val
		default:
			panic(fmt.Sprintf("unknown type of flag.Value: %T", f.Value))
		}
	}

	if f.Value == nil {
		if f.IsBool {
			f.wrapValue = &boolValue{new(bool)}
		} else {
			f.wrapValue = &stringValue{new(string)}
		}
	}

	if f.Placeholder == "" {
		f.Placeholder = "value"
	}

	envSet := false
	for _, name := range strings.Split(f.EnvVar, ",") {
		name = strings.TrimSpace(name)
		if value, ok := os.LookupEnv(name); ok {
			f.wrapValue.Set(value)
			envSet = true
			break
		}
	}

	if !envSet && f.DefValue != "" {
		f.wrapValue.Set(f.DefValue)
	}

	f.visited = false // reset
}

// Names returns the names including short names and aliases
func (f *Flag) Names() []string {
	names := strings.Split(f.Name, ",")
	for i, name := range names {
		names[i] = strings.TrimSpace(name)
	}
	return names
}

// SetValue sets the value of the named flag
func (f *Flag) SetValue(value string) error {
	f.visited = true
	return f.wrapValue.Set(value)
}

// GetValue returns the string value of flag
func (f *Flag) GetValue() string {
	return f.wrapValue.String()
}

func lookupFlag(flags []*Flag, name string) *Flag {
	for _, f := range flags {
		for _, n := range f.Names() {
			if n == name {
				return f
			}
		}
	}
	return nil
}
