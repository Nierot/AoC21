package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	part1Link = "https://adventofcode.com/2021/day/%d/input"
	part2Link = "https://adventofcode.com/2021/day/%d/answer"
)

func GrabIntInput(day int, part int) []int {
	var input []string
	var output []int

	input = GrabInput(day, part)

	for _, line := range input {
		integer, _ := strconv.Atoi(line)
		output = append(output, integer)
	}

	return output
}

func GrabInput(day int, part int) []string {
	var link string
	if part == 1 {
		link = fmt.Sprintf(part1Link, day)
	} else {
		link = fmt.Sprintf(part2Link, day)
	}

	if fileExists(day, part) {
		log.Println("File exists: reading from file")
		return readFromFile(day, part)
	}

	req, err := http.NewRequest("GET", link, nil)
	req.Header.Set("Cookie", GetAdventOfCodeSessionCookie())
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	writeToFile(day, part, strings.Split(string(body), "\n"))

	return strings.Split(string(body), "\n")
}

func readFromFile(day, part int) []string {
	var filename string

	filename = fmt.Sprintf("input/d%dp%d.txt", day, part)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func fileExists(day int, part int) bool {
	var filename string
	filename = fmt.Sprintf("input/d%dp%d.txt", day, part)

	if _, err := ioutil.ReadFile(filename); err == nil {
		return true
	}
	return false
}

func writeToFile(day int, part int, data []string) {
	var filename string

	filename = fmt.Sprintf("input/d%dp%d.txt", day, part)

	err := ioutil.WriteFile(filename, []byte(strings.Join(data, "\n")), 0644)
	if err != nil {
		panic(err)
	}
}
