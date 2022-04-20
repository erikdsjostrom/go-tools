//go:build ignore

// This test is currently broken, because the presence of constants in the instruction stream causes types to be used.
// For example, when returning type T, all the types of its fields will also be used, because we generate constants for
// them.

package main

type myNoCopy1 struct{}  // used
type myNoCopy2 struct{}  // used
type locker struct{}     // unused
type someStruct struct { // unused
	x int
}

func (myNoCopy1) Lock()      {} // used
func (recv myNoCopy2) Lock() {} // used
func (locker) Lock()         {} // unused
func (locker) Unlock()       {} // unused
func (someStruct) Lock()     {} // unused

type T struct { // used
	noCopy1 myNoCopy1  // used
	noCopy2 myNoCopy2  // used
	field1  someStruct // unused
	field2  locker     // unused
	field3  int        // unused
}

func main() { // used
	_ = T{}
}
