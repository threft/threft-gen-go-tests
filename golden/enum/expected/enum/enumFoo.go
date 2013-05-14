package enum

// This complete file is to be generated by the thrift4go generator.
// This is not implemented yet, (current enum structures consist of a typed int64 and constants in a file named ttypes.go)
// This file only contains code for the enum 'Foo', as defined in the thrift file enums.thrift
// The filename for this code is the combination of the word "enum" and the name of the generated enum, capitalized.
// For example: enum name: `Foo` becomes `enumFoo.go`
// Enum name is always capitalized. `secondExample` becomes `enumSecondExample.go`

import (
	"errors"
	"math"
)

// The Foo enum type
//=uncapitalized(enumname)+"EnumType"
type fooEnumType int64

// This is now tricky as InvalidOrUnexpectedEnumValueXXX == math.MinInt64
// Will return 0 for InvalidOrUnexpectedEnumValueXXX. http://play.golang.org/p/0KmE5tQoe-
func (fet fooEnumType) Int32() int32 {
	return int32(fet)
}

// This is now tricky as InvalidOrUnexpectedEnumValueXXX == math.MinInt64
// Will return 0 for InvalidOrUnexpectedEnumValueXXX when int is 32-bit. http://play.golang.org/p/0KmE5tQoe-
func (fet fooEnumType) Int() int {
	return int(fet)
}

func (fet fooEnumType) String() string {
	switch fet {
	case Foo.One:
		return "One"
	case Foo.Two:
		return "Two"
	case Foo.LowerCaseFirstChar:
		return "lowerCaseFirstChar" // note that "lowerCaseFirstChar" has the lowercase first char.
	case Foo.FortyTwo:
		return "FortyTwo"
	case Foo.InvalidOrUnexpectedEnumValueXXX:
		return "InvalidOrUnexpectedEnumValueXXX"
	}
	// This should NOT hapen. If this executes, code wasn't generated correctly.
	panic("invalid fooEnumType value.")
}

// Doubting if IsEnum() is required...
// Generator knows this so should generate code that "knows" this.
// This kind of metadata should not be required.
// func (fet fooEnumType) IsEnum() bool {
// 	return true
// }

// Unexported struct with each enum value as field
//=uncapitalized(enumname)+"EnumStruct"
type fooEnumStruct struct {
	One                             fooEnumType
	Two                             fooEnumType
	LowerCaseFirstChar              fooEnumType
	FortyTwo                        fooEnumType
	InvalidOrUnexpectedEnumValueXXX fooEnumType
}

// The Foo enum type instance
// The value for the enum value is either incremntal (if none is set in the .thrift file)
//   or it is the value that was defined in the .thrift file.
// The name of an enum is always capitalized. So `secondExample` is to become `SecondExample`
//=capitalized(enumname)
var Foo = &fooEnumStruct{
	One:                             fooEnumType(1),
	Two:                             fooEnumType(2),
	LowerCaseFirstChar:              fooEnumType(3), // upper case to make the field visible outside the package
	FortyTwo:                        fooEnumType(42),
	InvalidOrUnexpectedEnumValueXXX: fooEnumType(math.MinInt64), // Thrift defines all enums to be an i32 value. So math.MinInt64 is fine.
}

// ParseInt32 wrapper to use int
func (foo *fooEnumStruct) ParseInt(i int) (fooEnumType, error) {
	return foo.ParseInt32(int32(i))
}

// Takes an integer and returns the enum value + nil
// If there is no value for given number: returns InvalidOrUnexpectedEnumValueXXX + error.
func (foo *fooEnumStruct) ParseInt32(i int32) (fooEnumType, error) {
	switch i {
	case 1:
		return foo.One, nil
	case 2:
		return foo.Two, nil
	case 3:
		return foo.LowerCaseFirstChar, nil
	case 42:
		return foo.FortyTwo, nil
	}
	return foo.InvalidOrUnexpectedEnumValueXXX, errors.New("Given int/int32 does not match for a value on this enum.")
}

func (foo *fooEnumStruct) ParseString(s string) (fooEnumType, error) {
	switch s {
	case "One", "ONE", "one":
		return foo.One, nil
	case "Two", "TWO", "two":
		return foo.Two, nil
	case "lowerCaseFirstChar", "LowerCaseFirstChar", "LOWERCASEFIRSTCHAR", "lowercasefirstchar":
		// Note that LowerCaseFirstChar has a four cases in ParseString().
		// It has "lowerCaseFirstChar" as defined in the thrift file
		// AND "LowerCaseFirstChar" as capitalized to be exported in Go.
		return foo.LowerCaseFirstChar, nil
	case "FortyTwo", "FORTYTWO", "fortytwo":
		return foo.FortyTwo, nil
	}
	return foo.InvalidOrUnexpectedEnumValueXXX, errors.New("Given string does not match for a value on this enum.")
}