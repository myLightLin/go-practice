package main

import (
	"fmt"
)

func main() {
	fmt.Println(string(65))
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
