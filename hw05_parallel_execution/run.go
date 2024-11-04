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
	channel := make(chan Task)
	channelForErrors := make(chan int, n)
	countErrors := 0
	var returnErr error

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(ch <-chan Task) {
			defer wg.Done()
			for task := range ch {
				if err := task(); err != nil {
					channelForErrors <- 0
				}
			}
		}(channel)
	}

	for _, task := range tasks {
		select {
		case channel <- task:
		case <-channelForErrors:
			countErrors++
		}

		if countErrors >= m {
			break
		}
	}

	close(channel)
	wg.Wait()

	for i := 0; i < n; i++ {
		select {
		case <-channelForErrors:
			countErrors++
		default:
		}
	}

	close(channelForErrors)

	if countErrors >= m {
		returnErr = ErrErrorsLimitExceeded
	}

	return returnErr
}
