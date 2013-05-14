package main

import (
	"list-tryout/gen-go/list"
	"math"
	"testing"
)

func TestListInt32(t *testing.T) {
	myListInt32 := list.NewListInt32()
	//++ test myListInt32 type with reflect
	i32TestValues := []int32{math.MinInt32, (math.MinInt32 + 1), -100, -42, -4, -2, -1, 0, 1, 2, 4, 42, 100, (math.MaxInt32 - 1), math.MaxInt32}
	for _, insertValue := range i32TestValues {
		myListInt32.Push(insertValue)
	}
	i32Iter := myListInt32.Iter()
	for index, expectedValue := range i32TestValues {
		readValue, chanOpen := <-i32Iter
		if !chanOpen {
			t.Errorf("Iter channel for ListInt32 is closed. It is expected to be open as more values are expected. At #%d", index)
			break
		}
		if expectedValue != readValue {
			t.Errorf("i32 push/iter failed on #%d. Expected: %d. Received: %d.", index, expectedValue, readValue)
		}
	}
	if _, chanOpen := <-i32Iter; chanOpen {
		t.Error("Iter channel for ListInt32 is open. It is expected to be close as expected values have been read already.")
	}
}

func TestListString(t *testing.T) {
	myListString := list.NewListString()
	//++ test myListString type with reflect
	stringTestValues := []string{
		"Time is an illusion. Lunchtime doubly so.",
		"The ships hung in the sky in much the same way that bricks don't.",
		"You've got to know where your towel is.",
		"DON'T PANIC!",
		"Special chars: \n ' \" \\ ",
	}
	for _, insertValue := range stringTestValues {
		myListString.Push(insertValue)
	}
	stringIter := myListString.Iter()
	for index, expectedValue := range stringTestValues {
		readValue, chanOpen := <-stringIter
		if !chanOpen {
			t.Errorf("Iter channel for ListString is closed. It is expected to be open as more values are expected. At #%d", index)
			break
		}
		if expectedValue != readValue {
			t.Errorf("string push/iter failed on #%d. Expected: '%s'. Received: '%s'.", index, expectedValue, readValue)
		}
	}
	if _, chanOpen := <-stringIter; chanOpen {
		t.Error("Iter channel for ListString is open. It is expected to be close as expected values have been read already.")
	}
}
