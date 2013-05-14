package test

import (
	"enum-tryout/gen-go/enum"
	"fmt"
)

// user-code
func main() {
	myFooOne := enum.Foo.One
	fmt.Printf("Foo One. type:%T, value:%#v\n", myFooOne, myFooOne)

	mySecondExampleOne := enum.SecondExample.One
	fmt.Printf("SecondExample One. type:%T, value:%#v\n", mySecondExampleOne, mySecondExampleOne)

	// compiler error sice types are different:
	// if myFooOne == mySecondExampleOne 

	switch myFooOne {
	case enum.Foo.One:
		fmt.Println("ofcourse, this is printed")
	case enum.Foo.Two:
		fmt.Println("ofcourse this is not printed..")
	case enum.Foo.FortyTwo:
		fmt.Println("And this isn't printed either")
	case enum.Foo.InvalidOrUnexpectedEnumValueXXX:
		fmt.Println("Whoops.. that's unexpected..")
	}

	myLowerCaseFirstChar := enum.Foo.LowerCaseFirstChar
	fmt.Println(myLowerCaseFirstChar) // prints: lowerCaseFirstChar

	// it's imposible to create an invalid enum option like this:
	// var imposible = enum.barEnumType(12345)
	// because the barEnumType is unexported.
}
