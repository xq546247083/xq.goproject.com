package src

import (
	"fmt"
	"time"
)

/*
 * 任务分配器
 *
 */

type Dispather struct {
	// 工作池
	workerPool chan chan job

	// 工作池最大数量
	maxWorkerCount int32

	// 任务通道
	jobQueue chan job
}

// Run 启动任务分配器
func (this *Dispather) Run() {
	for i := 0; i < int(this.maxWorkerCount); i++ {
		worker := newWorker(this.workerPool)
		worker.start()
	}

	go this.dispather()
}

// AddJob 添加任务
// 参数
// _job:任务对象
func (this *Dispather) AddJob(_job job) {
	this.jobQueue <- _job
	fmt.Println("debug")
}

// 分配任务
func (this *Dispather) dispather() {
	for {
		select {
		// 从工作队列读取一个工作对象
		case jobObj := <-this.jobQueue:
			go func(_jobObj job) {
				// 从工作池取出一条空闲工作通道,将工作对象发送到空闲的工作通道中
				workChannel := <-this.workerPool
				workChannel <- _jobObj
			}(jobObj)

		default:
			//fmt.Println("no job...")
		}

		time.Sleep(time.Millisecond * 10)
	}
}

// NewDispather 新建任务分配器
func NewDispather(_maxWorkerCount int32) *Dispather {
	return &Dispather{
		workerPool:     make(chan chan job, _maxWorkerCount),
		maxWorkerCount: _maxWorkerCount,
		jobQueue:       make(chan job),
	}
}
