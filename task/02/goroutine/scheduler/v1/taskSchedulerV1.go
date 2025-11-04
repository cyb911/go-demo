package taskScheduler

import (
	"context"
	"errors"
	"sync"
	"time"
)

/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/

// 定义一个任务Task接口，任何实现 Run(ctx) 的类型都可以注册为任务
type Task interface {
	Run(ctx context.Context) error
}

type Result struct {
	ID        int
	StartedAt time.Time
	EndedAt   time.Time
	Duration  time.Duration
	Err       error
	Panic     any
}

// 调度器
type Scheduler struct {
	tasks    []Task
	worker   int
	taskChan chan int

	results []Result
	wg      sync.WaitGroup
	once    sync.Once
}

/*
构建任务调度器
tasks：任务列表
worker:执行的协程数（并发度）
*/
func BuildTasksScheduler(tasks []Task, worker int) *Scheduler {
	if worker <= 0 {
		worker = 1
	}
	return &Scheduler{
		tasks:    tasks,
		worker:   worker,
		taskChan: make(chan int),
		results:  make([]Result, len(tasks)),
	}
}

// 执行任务
func (s *Scheduler) Run(ctx context.Context) []Result {
	s.once.Do(func() {
		for i := 0; i < s.worker; i++ {
			s.wg.Add(1)
			go s.workerLoop(ctx)
		}
	})

	// 任务统一由调度器内部推送
	func() {
		for i := range s.tasks {
			select {
			case <-ctx.Done():
				return
			case s.taskChan <- i:
			}
		}
	}()

	close(s.taskChan)
	s.wg.Wait()
	return s.results
}

func (s *Scheduler) workerLoop(ctx context.Context) {
	defer s.wg.Done()
	for idx := range s.taskChan {
		s.executeTask(ctx, idx)
	}
}

func (s *Scheduler) executeTask(ctx context.Context, idx int) {
	start := time.Now()
	result := Result{ID: idx, StartedAt: start}

	func() {
		defer func() {
			if r := recover(); r != nil {
				result.Panic = r
			}
		}()
		task := s.tasks[idx]
		if task != nil {
			result.Err = task.Run(ctx)
		} else {
			result.Err = errors.New("nil task")
		}
	}()

	result.EndedAt = time.Now()
	result.Duration = result.EndedAt.Sub(start)
	s.results[idx] = result
}
