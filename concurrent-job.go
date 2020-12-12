package cjob

import (
	"sync"
)

type ConcurrentJob struct {
	ch chan struct{}
	wg sync.WaitGroup
}

func New(num int) (ret *ConcurrentJob) {
	ret = &ConcurrentJob{
		ch: make(chan struct{}, num),
	}

	return
}

func (this *ConcurrentJob) Do(fn func()) {
	this.wg.Add(1)
	go func() {
		this.ch <- struct{}{}
		fn()
		<-this.ch
		this.wg.Done()
	}()
}

func (this *ConcurrentJob) Wait() {
	this.wg.Wait()
	close(this.ch)
}
