package main

import (
	"fmt"
	"runtime"
	"sync"
)

var runCount int

func run(w *sync.WaitGroup) {
	runCount++
	fmt.Println(runCount)
	// 标记任务完成，减少WaitGroup
	w.Done()
}

func goRun() {
	// 调用CPU核数
	runtime.GOMAXPROCS(runtime.NumCPU())
	//c := make(chan bool, 10)
	// 创建WaitGroup
	wg := sync.WaitGroup{}
	// 添加10个WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go run(&wg)
	}
	//等待完成
	wg.Wait()
}

func useSelect() {
	c1, c2 := make(chan int), make(chan string)
	// 用于判断c1、c2是否关闭
	o := make(chan bool)
	go func() {
		// 使用无限循环
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1 ", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2 ", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hello"
	c1 <- 2
	c2 <- "world"

	close(c1)
	close(c2)

	<- o
}

func main() {
	useSelect()
}