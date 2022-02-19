package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 模拟一个网球赛，无缓冲通道例子
func main() {
	// create a unbuffered channel
	court := make(chan int)

	// add a count of two, one for each goroutine
	wg.Add(2)

	// launch two players
	go player("Nadal", court)
	go player("Djokovic", court)

	// start
	court <- 1

	// wait for the game to finish
	wg.Wait()
}

// player simulates a person playing the game of tennis
func player(name string, court chan int) {
	// defer the wait group
	defer wg.Done()

	for {
		// wait for the ball to be hit back to us
		ball, ok := <-court
		if !ok {
			// if the channel
			// was closed we're done
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// pick a random number
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// close the channel
			close(court)
			return
		}
		// hit the ball back
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// hit the ball back to the opposing player
		court <- ball
	}
}
