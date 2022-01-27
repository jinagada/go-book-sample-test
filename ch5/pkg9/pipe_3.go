package main

import (
	"fmt"
	"strconv"
)

const (
	JobCount = 10
	BufSize  = 5
)

type Job struct {
	name, log string
}
type jobHandler func(Job) Job

func (j Job) String() string {
	return "job name: " + j.name + "\n[log]\n" + j.log
}

func pipe(handler jobHandler, in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			out <- handler(job)
		}
		close(out)
	}()
	return out
}

func prepare() <-chan Job {
	out := make(chan Job, BufSize)
	go func() {
		for i := 0; i < JobCount; i++ {
			out <- Job{name: strconv.Itoa(i)}
		}
		close(out)
	}()
	return out
}

func first(job Job) Job {
	job.log += "first stage\n"
	return job
}

func second(job Job) Job {
	job.log += "second stage\n"
	return job
}

func third(job Job) Job {
	job.log += "third stage\n"
	return job
}

func last(job Job) Job {
	job.log += "last stage\n"
	return job
}

func main() {
	done := pipe(last, pipe(third, pipe(second, pipe(first, prepare()))))
	for d := range done {
		fmt.Println(d)
	}
}
