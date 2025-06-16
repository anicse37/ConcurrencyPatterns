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
	ResultId    int
	ResultValue int
}

const (
	NumbJobs    = 10
	NumbWorkers = 4
)

func worker(id int, jobs <-chan Job, result chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %v is doing, job no %v\n", id, job.JobId)
		result <- Result{ResultId: job.JobId, ResultValue: job.Value * 10}
	}
}

func main() {
	var wg sync.WaitGroup
	JobsCh := make(chan Job, NumbJobs)
	ResultCh := make(chan Result, NumbJobs)
	jobs := []Job{
		{JobId: 1, Value: 1},
		{JobId: 2, Value: 10},
		{JobId: 3, Value: 20},
		{JobId: 4, Value: 30},
		{JobId: 5, Value: 40},
		{JobId: 6, Value: 50},
		{JobId: 7, Value: 60},
		{JobId: 8, Value: 70},
		{JobId: 9, Value: 80},
		{JobId: 10, Value: 90},
	}
	for i := 1; i <= NumbWorkers; i++ {
		wg.Add(1)
		go worker(i, JobsCh, ResultCh, &wg)
	}
	for _, job := range jobs {
		JobsCh <- job
	}
	close(JobsCh)

	wg.Wait()
	close(ResultCh)
	fmt.Println("--------------------------")
	resMap := make(map[int]int)
	for i := 1; i <= len(jobs); i++ {
		value := <-ResultCh
		resMap[value.ResultId] = value.ResultValue
	}
	for i := 1; i <= len(jobs); i++ {
		fmt.Printf("Result Id: %v, Result Value:%v\n", i, resMap[i])
	}

}
