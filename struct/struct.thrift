// Because struct is a keyword in Go, we require the thrift file to define a namespace which will then be used as package name.
namespace go structure

// This structure represents a window
struct Window {
	1: i32  width,
	2: i32  height,
	3: bool broken,
}

// This structure represents a simple building
struct Building {
	1:          string Name,         // The name of this building, note that the field name is uppercase.
	2:          i32    height,       // The height of this building
	3:          bool   doorOpen,     // Indicates if the door of this building is open
	4:          Window frontWindow,  // Front window
	5: optional bool   backDoorOpen, // There doesn't have to be a backdoor in each building
	6: optional Window backWindow,   // There doesn't have to be a back window in each building
}

