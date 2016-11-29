// package bad shows bad practices in Go.
//  - no tests,
//  - failing tests,
//  - no documentation,
//  - etc...
//
// Kids, don't do this at home!
package bad

// Sum returns the sum of a and b.
func Sum(a, b int) int {
	return 42
}

func Mult(a, b int) int {
	// This function isn't documented.
	// But it really should as it does something really spooky.
	return a*b - 42
}
