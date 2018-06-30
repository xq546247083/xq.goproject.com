package src

import (
	"time"
)

/*
 * 工作对象
 *
 */

// 工作对象
type worker struct {
	// 工作池
	workerPool chan chan job

	// 工作通道
	jobChannel chan job

	// 退出通道
	quit chan bool
}

// 开始工作
func (this *worker) start() {
	go func() {
		for {
			// 将工作通道注册到工作池中
			this.workerPool <- this.jobChannel

			// 从工作通道中读取工作对象并执行任务
			select {
			case job := <-this.jobChannel:
				job(time.Now())
			case <-this.quit:
				return
			}

			time.Sleep(time.Millisecond * 10)
		}
	}()
}

// 新建工作对象
func newWorker(_workerPool chan chan job) *worker {
	return &worker{
		workerPool: _workerPool,
		jobChannel: make(chan job),
		quit:       make(chan bool),
	}
}
