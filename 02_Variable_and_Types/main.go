package main

import (
	f "fmt"
)

func main() {

	// Explicit
	var var1 int8 = 11

	// Implicit
	var var2 = 22
	var3 := 33

	// Multiple
	var4, var5 := 44, 55

	f.Println(var1, var2, var3, var4, var5)
}
