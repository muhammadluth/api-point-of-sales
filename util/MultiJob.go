package util

import "sync"

func MultiJob(worker ...func(waitGroup *sync.WaitGroup)) {
	var waitGroup sync.WaitGroup
	for _, work := range worker {
		waitGroup.Add(1)
		go work(&waitGroup)
	}
	waitGroup.Wait()
}
