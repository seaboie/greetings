// Basic print name with Go.   
// Just start learning go.
package greetings

import "fmt"

// Print hello name
func Hello(name string) string {
	message := fmt.Sprintf("Hi nice to meet you, %v. and welcome to thailand.", name)
	return message
}