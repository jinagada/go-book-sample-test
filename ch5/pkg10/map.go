package main

import (
	"fmt"
	"os"
	"text/scanner"
)

type partial struct {
	token string
	scanner.Position
}

func mapper(path string, out chan<- partial) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	if info, err := file.Stat(); err != nil || info.Mode().IsDir() {
		return
	}
	var s scanner.Scanner
	s.Filename = path
	s.Init(file)
	tok := s.Scan()
	for tok != scanner.EOF {
		fmt.Println(s.Pos())
		out <- partial{s.TokenText(), s.Pos()}
		tok = s.Scan()
	}
}

func runMap(paths <-chan string) <-chan partial {
	out := make(chan partial, BufSize)
	go func() {
		for path := range paths {
			mapper(path, out)
		}
		close(out)
	}()
	return out
}
