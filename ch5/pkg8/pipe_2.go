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

func (j Job) String() string {
	return "job name: " + j.name + "\n[log]\n" + j.log
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

func doFirst(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			job.log += "first stage\n"
			out <- job
		}
		close(out)
	}()
	return out
}

func doSecond(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			job.log += "second stage\n"
			out <- job
		}
		close(out)
	}()
	return out
}

func doThird(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			job.log += "third stage\n"
			out <- job
		}
		close(out)
	}()
	return out
}

func doLast(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			job.log += "last stage\n"
			out <- job
		}
		close(out)
	}()
	return out
}

func main() {
	done := doLast(doThird(doSecond(doFirst(prepare()))))
	for d := range done {
		fmt.Println(d)
	}
}
