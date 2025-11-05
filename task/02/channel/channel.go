package channel

import (
	"fmt"
	"sync"
)

/*
编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
*/

func PrintUnBuffer() {
	ch := make(chan int)
	// 负责发送
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		// 发送结束后，关闭通道（生产者负责通道的关闭）
		close(ch)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	// 负责打印
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("ch Print:", v)
		}
	}()

	wg.Wait()
}

/*
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
*/

func PrintBuffer(total int, size int) {
	if size < 0 {
		size = 1
	}
	ch := make(chan int, size)
	// 负责发送
	go func() {
		for i := 1; i <= total; i++ {
			ch <- i
		}
		// 发送结束后，关闭通道（生产者负责通道的关闭）
		close(ch)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	// 负责打印
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("ch Print:", v)
		}
	}()

	wg.Wait()
}

/*
多消费者模式下
*/
func PrintBufferV1(total int, size int, workers int) {
	if size < 0 {
		size = 1
	}
	ch := make(chan int, size)
	// 负责发送
	go func() {
		for i := 1; i <= total; i++ {
			ch <- i
		}
		// 发送结束后，关闭通道（生产者负责通道的关闭）
		close(ch)
	}()

	var wg sync.WaitGroup
	wg.Add(workers)
	// 负责打印
	for i := 1; i <= workers; i++ {
		go func(id int) {
			defer wg.Done()
			for v := range ch {
				fmt.Printf("worker-%d Print:%d\n", id, v)
			}
		}(i)
	}
	wg.Wait()
}
