package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
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
	gamma := calculateGammaRate(str)
	epsilon := calculateEpsilonRate(gamma)

	return getInt(gamma) * getInt(epsilon)
}

func getInt(str []string) int {
	i, err := strconv.ParseInt(fmt.Sprintf(strings.Join(str[:], "")), 2, 64)
	if err != nil {
		return 0
	}

	return int(i)
}

func calculateGammaRate(str []string) []string {
	res := make([]string, len(str[0]))
	common := make(map[int]int, len(str))

	for _, binary := range str {
		spl := strings.Split(binary, "")

		for i := range spl {
			if spl[i] == "0" {
				common[i] += 1
			}
		}
	}

	for k, v := range common {
		if v > len(str)/2 {
			res[k] = "0"
		} else {
			res[k] = "1"
		}
	}

	return res
}

func calculateEpsilonRate(str []string) []string {
	res := make([]string, len(str))

	for _, val := range str {
		if val == "0" {
			res = append(res, "1")
		} else {
			res = append(res, "0")
		}
	}

	return res
}

func secondPart(str []string) int {
	co2 := calculateCO2ScrubberRate(str, 0)
	oxygen := calculateOxygenGeneratorRate(str, 0)

	return getInt(strings.Split(co2[0], "")) * getInt(strings.Split(oxygen[0], ""))
}

func calculateCO2ScrubberRate(str []string, index int) []string {
	count := 0
	mp := make(map[int][]string, 2)

	for _, binary := range str {
		spl := strings.Split(binary, "")

		if spl[index] == "0" {
			count++
			mp[1] = append(mp[1], binary)
		} else {
			mp[0] = append(mp[0], binary)
		}
	}

	if count <= len(str)/2 {
		if len(mp[1]) > 1 {
			return calculateCO2ScrubberRate(mp[1], index+1)
		} else {
			return mp[1]
		}
	} else {
		if len(mp[0]) > 1 {
			return calculateCO2ScrubberRate(mp[0], index+1)
		} else {
			return mp[0]
		}
	}
}

func calculateOxygenGeneratorRate(str []string, index int) []string {
	count := 0
	mp := make(map[int][]string, 2)

	for _, binary := range str {
		spl := strings.Split(binary, "")

		if spl[index] == "1" {
			count++
			mp[1] = append(mp[1], binary)
		} else {
			mp[0] = append(mp[0], binary)
		}
	}

	if float64(count) >= math.Ceil(float64(len(str))/2) {
		if len(mp[1]) > 1 {
			return calculateOxygenGeneratorRate(mp[1], index+1)
		} else {
			return mp[1]
		}
	} else {
		if len(mp[0]) > 1 {
			return calculateOxygenGeneratorRate(mp[0], index+1)
		} else {
			return mp[0]
		}
	}
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
