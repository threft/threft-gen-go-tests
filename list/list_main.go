package main

// // This file shows how the generated lists are to be used in user-code.
// // For tests, see: list_test.go

// import (
// 	"list-tryout/gen-go/list"
// 	"log"
// )

// func main() {
// 	log.Println("Showing ListInt32.")
// 	{
// 		myListInt32 := list.NewListInt32()
// 		myListInt32.Push(42)
// 		myListInt32.Push(1337)

// 		// At some points, type conversion is inevitable.
// 		myInt := 12345
// 		myListInt32.Push(int32(myInt))

// 		// Types match, no conversion required.
// 		myInt32 := int32(123456)
// 		myListInt32.Push(myInt32)

// 		iterchan := myListInt32.Iter()
// 		for {
// 			value, chanOpen := <-iterchan
// 			if !chanOpen {
// 				break
// 			}
// 			log.Printf("got: %d\n", value)
// 		}
// 	}

// 	log.Println("Showing ListString.")
// 	{
// 		myListString := list.NewListString()
// 		myListString.Push("hi test")
// 		myListString.Push("Another string")
// 		myString := "hello there"
// 		myListString.Push(myString)
// 		myString2 := "0123456789"
// 		myListString.Push(myString2[2:6])

// 		iterchan := myListString.Iter()
// 		for {
// 			value, chanOpen := <-iterchan
// 			if !chanOpen {
// 				break
// 			}
// 			log.Printf("got: %s\n", value)
// 		}
// 	}
// }
