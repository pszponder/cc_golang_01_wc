// internal/ccwc/helper.go

package ccwc

// Check is a helper function to check for errors
//
//	If an error is found, the program will panic
//
// Parameters:
// e - The error to check
func check(e error) {
	if e != nil {
		panic(e)
	}
}
