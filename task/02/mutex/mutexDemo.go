package mutex

import (
	"sync"
	"sync/atomic"
)

/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
func Count(workers, countInc int) int {
	count := 0
	mutex := &sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			for i := 0; i < countInc; i++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	return count
}

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
func CountAtomic(workers, countInc int) int64 {
	var count int64
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < countInc; i++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	return count
}
