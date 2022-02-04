package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var (
	path    = flag.String("path", "..", "find path")
	pattern = flag.String("pattern", ".go$", "grep pattern")
)

const BufSize = 10

func find(dir, pattern string) <-chan string {
	out := make(chan string, 1000)
	regex, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	go func() {
		filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			if regex.MatchString(path) {
				out <- path
			}
			return nil
		})
		close(out)
	}()
	return out
}

func parseArgs() (string, string) {
	flag.Parse()
	return *path, *pattern
}

func mainOne() {
	fmt.Println(runReduce(collect(runMap(find(parseArgs())))))
}

func mainMulti() {
	fmt.Println(runConcurrentReduce(collect(runConcurrentMap(find(parseArgs())))))
}

func main() {
	//mainOne()
	mainMulti()
}
