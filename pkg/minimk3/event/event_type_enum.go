// Code generated by go-enum
// DO NOT EDIT!

package event

import (
	"fmt"
)

const (
	// Pressed is a Type of type Pressed
	Pressed Type = iota
	// Released is a Type of type Released
	Released
)

const _TypeName = "PressedReleased"

var _TypeMap = map[Type]string{
	0: _TypeName[0:7],
	1: _TypeName[7:15],
}

func (i Type) String() string {
	if str, ok := _TypeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Type(%d)", i)
}

var _TypeValue = map[string]Type{
	_TypeName[0:7]:  0,
	_TypeName[7:15]: 1,
}

// ParseType attempts to convert a string to a Type
func ParseType(name string) (Type, error) {
	if x, ok := _TypeValue[name]; ok {
		return Type(x), nil
	}
	return Type(0), fmt.Errorf("%s is not a valid Type", name)
}
