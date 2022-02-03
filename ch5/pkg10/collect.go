package main

import "text/scanner"

type intermediate map[string][]scanner.Position

func (m intermediate) addPartial(p partial) {
	positions, ok := m[p.token]
	if !ok {
		positions = make([]scanner.Position, 1)
	}
	positions = append(positions, p.Position)
	m[p.token] = positions
}

func collect(in <-chan partial) intermediate {
	tokenPositions := make(intermediate, 10)
	for t := range in {
		tokenPositions.addPartial(t)
	}
	return tokenPositions
}
