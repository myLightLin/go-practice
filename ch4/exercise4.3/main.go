package main

import "fmt"

func main() {
	a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	reverse(&a)
	fmt.Println(a)
}

func reverse(a *[8]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
