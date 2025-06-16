package main

import (
	"fmt"
	"sync"
)

type Job struct {
	JobId int
	Value int
}
type Result struct {
	WorkerId    int
	ResultId    int
	ResultValue int
}

const (
	NumbJobs   = 10
	NumbWorker = 4
)

func Worker(workerId int, Jobs <-chan Job, result chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range Jobs {
		fmt.Printf("Woker %v is doing Job no %v\n", workerId, job.JobId)
		result <- Result{
			WorkerId:    workerId,
			ResultId:    job.JobId,
			ResultValue: job.Value * 10,
		}
	}
}
func FadeIn(channels ...chan Result) chan Result {
	var wg sync.WaitGroup
	merged := make(chan Result)

	output := func(c <-chan Result) {
		defer wg.Done()
		for val := range c {
			merged <- val
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	var wg sync.WaitGroup
	JobCh := make(chan Job, NumbJobs)
	WorkCh := make([]chan Result, NumbWorker)

	jobs := []Job{
		{JobId: 1, Value: 10},
		{JobId: 2, Value: 20},
		{JobId: 3, Value: 30},
		{JobId: 4, Value: 40},
		{JobId: 5, Value: 50},
		{JobId: 6, Value: 60},
		{JobId: 7, Value: 70},
		{JobId: 8, Value: 80},
		{JobId: 9, Value: 90},
		{JobId: 10, Value: 100},
	}
	for i := 0; i < NumbWorker; i++ {
		WorkCh[i] = make(chan Result, NumbJobs)
		wg.Add(1)
		go Worker(i+1, JobCh, WorkCh[i], &wg)
	}
	for _, job := range jobs {
		JobCh <- job
	}
	close(JobCh)
	go func() {
		wg.Wait()
		for _, ch := range WorkCh {
			close(ch)
		}
	}()
	ResultFanIn := FadeIn(WorkCh...)

	for result := range ResultFanIn {
		fmt.Printf("Result from Worker %v: Job %v -> Result %v\n",
			result.WorkerId, result.ResultId, result.ResultValue)
	}

}
