package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 有缓冲通道例子
const (
	numberGoroutines = 4  // 要使用的goroutine数量
	taskLoad         = 10 // 每个goroutine要做的任务数量
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作都处理完成时关闭通道
	// 以便所有 goroutine 退出
	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 等待分派工作
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了，已被关闭
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// 随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们完成了这个工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
