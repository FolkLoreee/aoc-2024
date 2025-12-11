package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"
)

const (
	inputPathTest   = "./test.txt"
	inputPathActual = "./input.txt"

	numWorkers = 10
)

type Bank struct {
	Batteries []Battery
	Output    int
}

func (b *Bank) CalculateOutput() {
	var (
		leftBat  Battery = 0
		rightBat Battery = 0
	)
	for i := range b.Batteries {
		bat := b.Batteries[i]
		if bat > leftBat && i != len(b.Batteries)-1 {
			leftBat = bat
			rightBat = b.Batteries[i+1]
		} else if bat > rightBat {
			rightBat = bat
		}
	}

	b.Output = combineBatteries(leftBat, rightBat)
}

func (b *Bank) CalculateOutput2() {
	var inner func(start, remaining int) int
	inner = func(start, remaining int) int {
		var (
			largest    Battery = 0
			largestIdx         = 0
		)

		for i := start; i < len(b.Batteries)-remaining+1; i++ {
			current := b.Batteries[i]
			if current > largest {
				largest = current
				largestIdx = i
			}
		}

		if remaining > 1 {
			return int(largest)*int(math.Pow10(remaining-1)) + inner(largestIdx+1, remaining-1)
		} else {
			return int(largest)
		}
	}

	b.Output = inner(0, 12)
}

type Battery int

func parseBattery(ch rune) Battery {
	return Battery(ch - '0')
}

func combineBatteries(b1, b2 Battery) int {
	return int(b1*10 + b2)
}

func parseBank(line string) *Bank {
	bank := &Bank{
		Batteries: make([]Battery, 0, len(line)),
		Output:    0,
	}
	for _, ch := range line {
		bank.Batteries = append(bank.Batteries, parseBattery(ch))
	}
	return bank
}

func part1(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal("error opening file: ", err)
	}
	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)
		banks   []*Bank
	)
	for scanner.Scan() {
		line := scanner.Text()
		bank := parseBank(line)
		banks = append(banks, bank)
	}

	var (
		wg          sync.WaitGroup
		bankChannel = make(chan *Bank, numWorkers)
	)
	wg.Add(len(banks))
	for range numWorkers {
		go calculateOutput(bankChannel, &wg)
	}
	for _, bank := range banks {
		bankChannel <- bank
	}
	wg.Wait()

	totalOutput := 0
	for _, bank := range banks {
		totalOutput += bank.Output
	}
	fmt.Printf("\nTotal joltage: %d", totalOutput)
}

func calculateOutput(bankChannel <-chan *Bank, wg *sync.WaitGroup) {
	for bank := range bankChannel {
		bank.CalculateOutput()
		wg.Done()
	}
}

func part2(inputPath string) {
	start := time.Now()
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal("error opening file: ", err)
	}
	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)
		banks   []*Bank
	)
	for scanner.Scan() {
		line := scanner.Text()
		bank := parseBank(line)
		banks = append(banks, bank)
	}

	var (
		wg          sync.WaitGroup
		bankChannel = make(chan *Bank, numWorkers)
	)
	wg.Add(len(banks))
	for range numWorkers {
		go calculateOutput2(bankChannel, &wg)
	}
	for _, bank := range banks {
		bankChannel <- bank
	}
	wg.Wait()

	totalOutput := 0
	for _, bank := range banks {
		totalOutput += bank.Output
	}
	elapsed := time.Since(start)
	fmt.Printf("\nTotal joltage: %d\nTime taken: %d µs", totalOutput, elapsed.Microseconds())
}

func calculateOutput2(bankChannel <-chan *Bank, wg *sync.WaitGroup) {
	for bank := range bankChannel {
		bank.CalculateOutput2()
		wg.Done()
	}
}

func main() {
	testFlag := flag.Bool("test", false, "run against test input")
	flag.Parse()

	inputPath := inputPathActual
	if testFlag != nil && *testFlag {
		inputPath = inputPathTest
	}
	part2(inputPath)
}
