package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for k, v := range os.Args {
		fmt.Println(k, v)
	}
	end := time.Since(start)
	fmt.Println(end)
	fmt.Println(strings.Join(os.Args[1:], ","))
}
