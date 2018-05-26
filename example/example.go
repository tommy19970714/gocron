package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am runnning task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	// Do jobs with params
	gocron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// Do jobs without params
	gocron.Every(1).Second().Do(task)
	gocron.Every(2).Seconds().Do(task)
	gocron.EveryWithId(1, 1).Minute().Do(task) //first parameter is job ID
	gocron.EveryWithId(2, 2).Minutes().Do(task)
	gocron.EveryOnlyId(3).Hour().Do(task) //first parameter is job ID and interval is 1
	gocron.Every(2).Hours().Do(task)
	gocron.Every(1).Day().Do(task)
	gocron.Every(2).Days().Do(task)

	// Do jobs on specific weekday
	gocron.Every(1).Monday().Do(task)
	gocron.Every(1).Thursday().Do(task)

	// function At() take a string like 'hour:min'
	gocron.Every(1).Day().At("10:30").Do(task)
	gocron.Every(1).Monday().At("18:30").Do(task)

	// remove, clear and next_run
	_, time := gocron.NextRun()
	fmt.Println(time)

	// get all jobs list
	jobs := gocron.AllJobs()
	for i, job := range jobs {
		fmt.Println(i, job.RunTime())
	}

	// gocron.RemoveFromFunc(task)
	// gocron.RemoveFromId(3)
	// gocron.Clear()

	// function Start start all the pending jobs
	<-gocron.Start()

	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(task)
	<-s.Start()
}
