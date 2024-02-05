package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofish2020/tinygpool"
)

func main() {
	exec()
	execTimeOut()
}

func exec() {
	w := sync.WaitGroup{}
	demo := func() {
		w.Done()
		fmt.Println("Hello World")
	}

	gPool := tinygpool.NewPool(10, 1)

	for i := 0; i < 20; i++ {
		w.Add(1)
		gPool.Exec(demo)
	}

	w.Wait()
}

func execTimeOut() {
	w := sync.WaitGroup{}

	demo := func() {
		// 10s模拟任务处理时间很久
		time.Sleep(10 * time.Second)
		fmt.Println("Hello World")
		w.Done()

	}

	// 最多启动10个协程处理任务
	gPool := tinygpool.NewPool(10, 0)

	for i := 0; i < 20; i++ {
		w.Add(1)

		// 并行追加20个任务
		go func() {
			// 10个任务可以正常处理，10个任务因为没有多余协程可以处理而添加超时
			err := gPool.ExecTimeout(demo, 1*time.Second)
			if err != nil {
				fmt.Println(err.Error())
				w.Done()
			}
		}()
	}
	w.Wait()
}
