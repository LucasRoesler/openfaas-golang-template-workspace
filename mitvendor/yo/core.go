package yo

import "fmt"

// Yo is a function that takes a string and returns a string
func Yo[T string | []byte](s T) string {
	return fmt.Sprintf("Yo %s!", s)
}
