package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nierot/aoc21/v2/util"
)

func D2P1() int {
	var (
		puzzle     = util.GrabInput(2)
		horizontal = 0
		depth      = 0
	)

	for _, command := range puzzle {
		var (
			cmds      = strings.Fields(command)
			_, _      = fmt.Println(cmds)
			number, _ = strconv.Atoi(cmds[1])
		)

		switch cmds[0] {
		case "forward":
			horizontal += number
			break
		case "down":
			depth += number
			break
		case "up":
			depth -= number
			break
		}
	}

	return horizontal * depth
}

func D2P2() int {
	var (
		puzzle     = util.GrabInput(2)
		horizontal = 0
		depth      = 0
		aim        = 0
	)

	for _, command := range puzzle {
		var (
			cmds      = strings.Fields(command)
			_, _      = fmt.Println(cmds)
			number, _ = strconv.Atoi(cmds[1])
		)

		switch cmds[0] {
		case "forward":
			horizontal += number
			depth += aim * number
			break
		case "down":
			aim += number
			break
		case "up":
			aim -= number
			break
		}
	}

	return horizontal * depth
}
