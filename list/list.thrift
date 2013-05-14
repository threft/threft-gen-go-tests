
// Thift file lists.thrift
// For each unique list type in this file: a defintion is to be generated into gen-go/lists.
// The struct in this thrft file is not provied as generated go code yet.
// list<string> is used twice in this file (recipes and notes), however, it is generated only once (listString.go).

struct StructWithLists {
	1: list<i32>      phonenumbers,
	2: list<string>   recipes,
	2: list<string>   notes,
}
