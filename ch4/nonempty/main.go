package main

import "fmt"

func main() {
	s := []string{"a", "", "c", "", "d", "w"}
	fmt.Println(nonempty(s))
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
