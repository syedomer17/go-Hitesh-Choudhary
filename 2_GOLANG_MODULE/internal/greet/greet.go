package greet 

import "strings"

func Hello(name string) string {
	clean :=  normalizedName(name)

	return "Hello, " + clean + "!"
}

// if u want to use this function outside of this package, u need to export it by capitalizing the first letter of the function name
func normalizedName(name string) string {
	n := strings.TrimSpace(name)

	if n == "" {
		return "Guest"
	}
	return strings.ToUpper(n)
}