package main

import (
	"github.com/nierot/aoc21/v2/util"
)

func D1P1() int {
	var (
		input          = util.GrabIntInput(1)
		depthIncreases = -1
		previous       = 0
	)

	for _, depth := range input {
		if depth > previous {
			depthIncreases++
		}
		previous = depth
	}
	return depthIncreases
}

func D1P2() int {
	var (
		input          = util.GrabIntInput(1)
		window1        = 0
		window2        = 0
		depthIncreases = 0
		length         = len(input)
	)

	for idx, depth := range input {
		if idx+3 == length {
			break
		}
		window1 = depth + input[idx+1] + input[idx+2]
		window2 = input[idx+1] + input[idx+2] + input[idx+3]
		if window2 > window1 {
			depthIncreases++
		}
	}

	return depthIncreases
}
