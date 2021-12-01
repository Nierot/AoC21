package main

import (
	"log"
	"time"
)

type solver func() int

func main() {
	Run(1, 2, D1P2)
}

func Run(day int, part int, solver solver) {
	log.Printf("Running day %d part %d\n", day, part)
	start := time.Now()
	log.Println(solver())
	elapsed := time.Since(start)
	log.Printf("Day %d part %d took %s\n", day, part, elapsed)
}
