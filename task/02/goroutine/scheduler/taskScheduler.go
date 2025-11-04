package taskscheduler

import (
	"context"
	"errors"
	"sync"
	"time"
)

/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/

type Task func(ctx context.Context) error

type Result struct {
	ID        int
	StartedAt time.Time
	EndedAt   time.Time
	Duration  time.Duration
	Err       error
	Panic     any
}

/*
执行任务
ctx context ：上下文
tasks：任务列表
worker:执行的协程数（并发度）
*/
func RunTasks(ctx context.Context, tasks []Task, worker int) []Result {

	count := len(tasks)
	var taskWg sync.WaitGroup

	// 启动 worker

	// 默认至少有一个工作协程
	if worker <= 0 {
		worker = 1
	}
	//设置等待的协程的数量
	taskWg.Add(worker)
	// 定义任务通道
	taskChan := make(chan int)

	results := make([]Result, count)

	for i := 0; i < worker; i++ {
		go executeTask(ctx, &taskWg, taskChan, tasks, results)
	}

	// 向chan中投递信息，使用int值作为协作的信号量
loop:
	for i := 0; i < count; i++ {
		select {
		// 取消后，不在投递后续任务
		case <-ctx.Done():
			break loop
		case taskChan <- i:
		}
	}
	close(taskChan)
	taskWg.Wait()
	return results
}

func executeTask(ctx context.Context, wg *sync.WaitGroup, taskChan <-chan int, tasks []Task, results []Result) {
	defer wg.Done()

	for idx := range taskChan {
		start := time.Now()
		result := Result{ID: idx, StartedAt: start}

		// 单个任务可继承上层 ctx（也可按需给每个任务单独设置超时）
		func() {
			defer func() {
				if r := recover(); r != nil {
					result.Panic = r
				}
			}()
			if tasks[idx] != nil {
				result.Err = tasks[idx](ctx)
			} else {
				result.Err = errors.New("nil task")
			}
		}()

		result.EndedAt = time.Now()
		result.Duration = result.EndedAt.Sub(start)
		results[idx] = result
	}
}
