package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Range struct {
	Start int
	End   int
	Sum   int
}

func FromString(in string) (*Range, error) {
	arr := strings.Split(in, "-")
	start, err := strconv.Atoi(arr[0])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, err
	}
	return &Range{
		Start: start,
		End:   end,
	}, nil
}

func (r *Range) CountSum() {
	for num := r.Start; num <= r.End; num++ {
		if isInvalid(num) {
			r.Sum += num
		}
	}
}

func (r *Range) CountSum2() {
	for num := r.Start; num <= r.End; num++ {
		if isInvalid2(num) {
			r.Sum += num
		}
	}
}

func isInvalid(num int) bool {
	var (
		str = strconv.Itoa(num)
		l   = 0
		r   = len(str) / 2
	)

	if len(str)%2 != 0 {
		return false
	}

	for r < len(str) {
		if str[l] != str[r] {
			return false
		}
		r++
		l++
	}

	println("INVALID: ", num)
	return true
}

func isInvalid2(num int) bool {
	var (
		str      = strconv.Itoa(num)
		divisors []int
	)
	for i := 1; i <= len(str); i++ {
		if len(str)%i == 0 {
			divisors = append(divisors, i)
		}
	}

	// try all r starting points
	for _, divisor := range divisors {
		var (
			l = 0
			r = len(str) / divisor
		)
		duplicate := false
		for r < len(str) {
			if str[l] != str[r] {
				duplicate = false
				break
			} else {
				duplicate = true
			}
			r++
			l++
		}
		if duplicate {
			println("INVALID: ", num)
			return true
		}
	}

	return false
}

func CountSum(rangeChan <-chan *Range, wg *sync.WaitGroup) {
	defer wg.Done()
	for rng := range rangeChan {
		rng.CountSum()
	}
}

func CountSum2(rangeChan <-chan *Range, wg *sync.WaitGroup) {
	defer wg.Done()
	for rng := range rangeChan {
		rng.CountSum2()
	}
}

const (
	inputPath = "./input.txt"
	numWorker = 10
)

func part1() {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var ranges []*Range
	for scanner.Scan() {
		rangeString := scanner.Text()
		rangeStrArr := strings.Split(rangeString, ",")
		for _, rangeStr := range rangeStrArr {
			rangeInt, err := FromString(rangeStr)
			if err != nil {
				log.Fatal(err)
			}
			ranges = append(ranges, rangeInt)
		}
	}

	var (
		wg        sync.WaitGroup
		rangeChan = make(chan *Range, numWorker)
	)
	wg.Add(numWorker)

	for range numWorker {
		go CountSum(rangeChan, &wg)
	}

	for _, rng := range ranges {
		rangeChan <- rng
	}
	close(rangeChan)
	wg.Wait()

	var sum int64
	for _, rng := range ranges {
		sum += int64(rng.Sum)
	}
	println("sum: ", sum)
}

func part2() {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var ranges []*Range
	for scanner.Scan() {
		rangeString := scanner.Text()
		rangeStrArr := strings.Split(rangeString, ",")
		for _, rangeStr := range rangeStrArr {
			rangeInt, err := FromString(rangeStr)
			if err != nil {
				log.Fatal(err)
			}
			ranges = append(ranges, rangeInt)
		}
	}

	var (
		wg        sync.WaitGroup
		rangeChan = make(chan *Range, numWorker)
	)
	wg.Add(numWorker)

	for range numWorker {
		go CountSum2(rangeChan, &wg)
	}

	for _, rng := range ranges {
		rangeChan <- rng
	}
	close(rangeChan)
	wg.Wait()

	var sum int64
	for _, rng := range ranges {
		sum += int64(rng.Sum)
	}
	println("sum: ", sum)
}

func main() {
	part2()
}
