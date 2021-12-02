package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
)

func main() {
	byteData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	r := bytes.NewReader(byteData)
	integers, err := readInts(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", firstPart(integers))
	fmt.Printf("%v\n", secondPart(integers))
}

func firstPart(integers []int) int {
	timesIncreased := -1
	prev := 0

	for _, value := range integers {
		if value > prev {
			timesIncreased++
		}
		prev = value
	}

	return timesIncreased
}

func secondPart(integers []int) int {
	window := getSlidingWindowMeasurements(integers)

	timesIncreased := firstPart(window)

	return timesIncreased
}

func getSlidingWindowMeasurements(ints []int) []int {
	resp := make([]int, len(ints)/3)

	for i := 2; i < len(ints); i++ {
		resp = append(resp, ints[i-2]+ints[i-1]+ints[i])
	}

	return resp
}

func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}
