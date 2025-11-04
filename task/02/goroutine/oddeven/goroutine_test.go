package oddeven_test

import (
	"fmt"
	"go-demo/task/02/goroutine/oddeven"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2) //设置等待协程数2

	go oddeven.PrintOdd(&wg, 10)
	go oddeven.PrintEven(&wg, 10)

	wg.Wait()

	fmt.Println("主协程结束")
}
