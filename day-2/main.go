package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	byteData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	r := bytes.NewReader(byteData)
	str, err := readString(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", firstPart(str))
	fmt.Printf("%v\n", secondPart(str))
}

func firstPart(str []string) int {
	fwd := 0
	depth := 0

	for _, value := range str {
		split := strings.Split(value, " ")
		val, err := strconv.Atoi(split[1])
		if err != nil {
			return 0
		}
		switch split[0] {
		case "forward":
			fwd += val
		case "up":
			depth -= val
		case "down":
			depth += val
		}

	}

	return fwd * depth
}

func secondPart(str []string) int {
	fwd := 0
	aim := 0
	depth := 0

	for _, value := range str {
		split := strings.Split(value, " ")
		val, err := strconv.Atoi(split[1])
		if err != nil {
			return 0
		}
		switch split[0] {
		case "forward":
			fwd += val
			depth += val * aim
		case "up":
			aim -= val
		case "down":
			aim += val
		}

	}

	return fwd * depth
}

func readString(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		x := scanner.Text()
		result = append(result, x)
	}
	return result, scanner.Err()
}
