package pulp

import "fmt"

// What is a function that takes a string and returns a string
func What[T string | []byte](s T) string {
	return fmt.Sprintf("Say %q one more time!", string(s))
}
