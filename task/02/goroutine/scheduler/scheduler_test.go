package taskscheduler_test

import (
	"context"
	"errors"
	"fmt"
	taskscheduler "go-demo/task/02/goroutine/scheduler"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	// 设置整体超时：所有任务最多 3 秒
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 定义一组任务(函数)
	tasks := []taskscheduler.Task{
		func(ctx context.Context) error {
			time.Sleep(400 * time.Millisecond)
			fmt.Println("任务1完成")
			return nil
		},
		func(ctx context.Context) error {
			time.Sleep(700 * time.Millisecond)
			fmt.Println("任务2完成")
			return errors.New("任务2出错：XX错误")
		},
		func(ctx context.Context) error {
			// 模拟可中断的任务
			select {
			case <-time.After(2 * time.Second):
				fmt.Println("任务3完成")
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		},
		func(ctx context.Context) error {
			// 故意触发 panic
			panic("任务4发生 panic")
		},
	}

	concurrency := 2
	results := taskscheduler.RunTasks(ctx, tasks, concurrency)

	fmt.Println("===== 执行结果 =====")
	for _, result := range results {
		// 结果可能因取消而有零值（未投递的任务 ID 会是默认 0）
		// 用 StartedAt.IsZero 判断是否执行过
		if result.StartedAt.IsZero() && result.EndedAt.IsZero() && result.Duration == 0 && result.Err == nil && result.Panic == nil {
			continue
		}
		fmt.Printf("任务ID=%d | 开始=%s | 结束=%s | 耗时=%v | 错误=%v | Panic=%v\n",
			result.ID, result.StartedAt.Format("15:04:05.000"), result.EndedAt.Format("15:04:05.000"), result.Duration, result.Err, result.Panic)
	}
}
