package oddeven

import (
	"fmt"
	"sync"
)

func PrintOdd(wg *sync.WaitGroup, maxNum int) {
	defer wg.Done()
	for i := 1; i <= maxNum; i += 2 {
		fmt.Println("奇数：", i)
	}
}

func PrintEven(wg *sync.WaitGroup, maxNum int) {
	defer wg.Done()
	for i := 2; i <= maxNum; i += 2 {
		fmt.Println("偶数：", i)
	}
}
