package channel

import (
	"fmt"
	"testing"
)

func TestPrintUnBuffer(t *testing.T) {
	fmt.Println("=== 1: 无缓冲通道，1..10 ===")
	PrintUnBuffer()
}

func TestPrintBuffer(t *testing.T) {
	fmt.Println("=== 2: 有冲通道 ===")
	PrintBuffer(100, 10)
}

func TestPrintBufferV1(t *testing.T) {
	fmt.Println("=== 3: 多消费者 ===")
	PrintBufferV1(100, 10, 2)
}
