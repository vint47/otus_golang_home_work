package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	channel := make(chan Task)

	var countErrors, maxError uint32 = 0, uint32(m)
	var returnErr error

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range channel {
				if err := task(); err != nil {
					mu.Lock()
					countErrors++
					mu.Unlock()
				}
			}
		}()
	}

	for _, task := range tasks {
		mu.RLock()
		isFinish := countErrors >= maxError
		mu.RUnlock()

		if isFinish {
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
