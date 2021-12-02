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
	inputLink = "https://adventofcode.com/2021/day/%d/input"
)

// Grab the input for the current puzzle in the form of an int array
// Day 0 is for the example
func GrabIntInput(day int) []int {
	var input []string
	var output []int

	input = GrabInput(day)

	for _, line := range input {
		integer, _ := strconv.Atoi(line)
		output = append(output, integer)
	}

	return output[:len(output)-1]
}

func GrabInput(day int) []string {
	link := fmt.Sprintf(inputLink, day)

	if fileExists(day) {
		log.Println("File exists: reading from file")
		return readFromFile(day)
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

	writeToFile(day, strings.Split(string(body), "\n"))

	slice := strings.Split(string(body), "\n")
	slice = slice[:len(slice)-1]

	return slice
}

func readFromFile(day int) []string {
	var filename string

	filename = fmt.Sprintf("input/d%d.txt", day)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func fileExists(day int) bool {
	var filename string
	filename = fmt.Sprintf("input/d%d.txt", day)

	if _, err := ioutil.ReadFile(filename); err == nil {
		return true
	}
	return false
}

func writeToFile(day int, data []string) {
	var filename string

	filename = fmt.Sprintf("input/d%d.txt", day)

	err := ioutil.WriteFile(filename, []byte(strings.Join(data, "\n")), 0644)
	if err != nil {
		panic(err)
	}
}
