package enum

// This file only contains code for the enum 'secondExample', as defined in the thrift file enums.thrift
// !!!!!!!!!!!!! For information and explaining comments, please see foo.go !!!!!!

import (
	"errors"
	"math"
)

// The SecondExample enum type
type secondExampleEnumType int64

func (bet secondExampleEnumType) Int32() int32 {
	return int32(bet)
}

func (bet secondExampleEnumType) Int() int {
	return int(bet)
}

func (bet secondExampleEnumType) String() string {
	switch bet {
	case SecondExample.One:
		return "One"
	case SecondExample.Elite:
		return "Elite"
	}
	panic("invalid secondExampleEnumType value.")
}

type secondExampleEnumStruct struct {
	One                             secondExampleEnumType
	Elite                           secondExampleEnumType
	InvalidOrUnexpectedEnumValueXXX secondExampleEnumType
}

var SecondExample = &secondExampleEnumStruct{
	One:   secondExampleEnumType(1),
	Elite: secondExampleEnumType(1337),
	InvalidOrUnexpectedEnumValueXXX: secondExampleEnumType(math.MinInt64),
}

func (secondExample *secondExampleEnumStruct) ParseInt(i int) (secondExampleEnumType, error) {
	return secondExample.ParseInt32(int32(i))
}

func (secondExample *secondExampleEnumStruct) ParseInt32(i int32) (secondExampleEnumType, error) {
	switch i {
	case 1:
		return secondExample.One, nil
	case 1337:
		return secondExample.Elite, nil
	}
	return secondExample.InvalidOrUnexpectedEnumValueXXX, errors.New("Given int/int32 does not match for a value on this enum.")
}

func (secondExample *secondExampleEnumStruct) ParseString(s string) (secondExampleEnumType, error) {
	switch s {
	case "One", "ONE", "one":
		return secondExample.One, nil
	case "Elite", "ELITE", "elite":
		return secondExample.Elite, nil
	}
	return secondExample.InvalidOrUnexpectedEnumValueXXX, errors.New("Given string does not match for a value on this enum.")
}
