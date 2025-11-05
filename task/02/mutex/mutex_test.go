package mutex

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	fmt.Println("Mutex 计数结果：", Count(10, 100))
}

func TestAtomic(t *testing.T) {
	fmt.Println("Atomic 计数结果：", CountAtomic(10, 100))
}
