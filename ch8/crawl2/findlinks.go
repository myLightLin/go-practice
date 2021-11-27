package main

import (
	"fmt"
	"log"
	"os"

	"github.com/myLightLin/go-practice/ch5/links"
)

var tokes = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokes <- struct{}{}
	list, err := links.Extract(url)
	<-tokes

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
