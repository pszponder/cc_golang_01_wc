// internal/ccwc/helper.go

package ccwc

// Helper function to check for errors
//
//	If an error is found, the program will panic
//
// Parameters:
// e - The error to check
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
