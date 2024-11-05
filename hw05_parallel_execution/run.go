package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := &sync.WaitGroup{}
	channel := make(chan Task)

	var countErrors, maxError uint32 = 0, uint32(m)
	var returnErr error

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(ch <-chan Task) {
			defer wg.Done()
			for task := range ch {
				if err := task(); err != nil {
					atomic.AddUint32(&countErrors, 1)
				}
			}
		}(channel)
	}

	for _, task := range tasks {
		if countErrors >= maxError {
			break
		}

		channel <- task
	}

	close(channel)
	wg.Wait()

	if countErrors >= maxError {
		returnErr = ErrErrorsLimitExceeded
	}

	return returnErr
}
