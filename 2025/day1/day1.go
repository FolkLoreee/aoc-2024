package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	maxDial   = 100
	dialStart = 50
)

var re = regexp.MustCompile(`([RL])(\d+)`)

func part1() {
	zeroCount := 0
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	currentPos := dialStart
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		matches := re.FindStringSubmatch(str)
		if matches == nil {
			print("no match found")
			return
		}
		direction := matches[1]
		distance, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Println("string: ", str)
			log.Println("strArray: ", matches)
			log.Fatal(err)
		}

		initialPos := currentPos
		distance %= maxDial

		if direction == "L" {
			distance *= -1
		}

		currentPos += distance
		if currentPos < 0 {
			currentPos = maxDial + currentPos
		} else {
			currentPos %= maxDial
		}
		if currentPos == 0 {
			zeroCount++
		}

		fmt.Printf(
			`
			Moving from: %d
			Distance: %d
			Destination: %d
			`, initialPos, distance, currentPos,
		)
	}
	fmt.Println("Total zero count: ", zeroCount)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2() {
	zeroCount := 0
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	currentPos := dialStart
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		matches := re.FindStringSubmatch(str)
		if matches == nil {
			print("no match found")
			return
		}
		direction := matches[1]
		distance, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Println("string: ", str)
			log.Println("strArray: ", matches)
			log.Fatal(err)
		}

		// Count how many times we pass through 0 during this rotation
		var firstHit int
		if direction == "R" {
			// Going right: we hit 0 at (100 - currentPos) clicks, then every 100 after
			firstHit = (maxDial - currentPos) % maxDial
			if firstHit == 0 {
				firstHit = maxDial
			}
		} else {
			// Going left: we hit 0 at currentPos clicks, then every 100 after
			firstHit = currentPos % maxDial
			if firstHit == 0 {
				firstHit = maxDial
			}
		}

		// Count how many times we hit 0 during this rotation
		if distance >= firstHit {
			count := (distance-firstHit)/maxDial + 1
			zeroCount += count
		}

		// Update current position (same logic as part1)
		distanceMod := distance % maxDial
		if direction == "L" {
			distanceMod *= -1
		}
		currentPos += distanceMod
		if currentPos < 0 {
			currentPos = maxDial + currentPos
		} else {
			currentPos %= maxDial
		}
	}

	fmt.Println("Total zero count (part 2):", zeroCount)
}

func main() {
	part2()
}
