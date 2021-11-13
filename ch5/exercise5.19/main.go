package main

import "fmt"

func main() {
	fmt.Println(returnNoZero())
}

func returnNoZero() (result int) {
	defer func() {
		result = 3
		_ = recover()
	}()
	panic("panic!")
}
