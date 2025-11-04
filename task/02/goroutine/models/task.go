package task

import (
	"context"
	"time"
)

type WorkTask struct {
	Message string
}

func (t *WorkTask) Run(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		time.Sleep(100 * time.Millisecond)
		println("执行任务:", t.Message)
		return nil
	}
}

type BackHomeTask struct {
	Message string
}

func (t *BackHomeTask) Run(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		time.Sleep(400 * time.Millisecond)
		println("执行任务:", t.Message)
		return nil
	}
}
