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
	input, err := readInput(r)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%v\n", firstPart(input))
	fmt.Printf("%v\n", secondPart(input))
}

func secondPart(str map[int][]int) int {
	bingoNumbers := str[0]
	delete(str, 0)

	boards := makeBoards(str)
	amt := len(boards)

	for _, nbr := range bingoNumbers {
		boards = markBoards(boards, nbr)
		winLoop := true
		for winLoop {
			win, boardNbr := checkIfWinner(boards)
			winLoop = win
			if win {
				if amt == 1 {
					return getResult(boards[boardNbr], nbr)
				}
				delete(boards, boardNbr)
				amt--
			}
		}
	}

	return 0
}

func firstPart(str map[int][]int) int {
	bingoNumbers := str[0]
	delete(str, 0)

	boards := makeBoards(str)

	for _, nbr := range bingoNumbers {
		boards = markBoards(boards, nbr)
		win, boardNbr := checkIfWinner(boards)
		if win {
			return getResult(boards[boardNbr], nbr)
		}
	}

	return 0
}

func getResult(board *[5][5]int, nbr int) int {
	totalBoard := 0

	/* fmt.Printf("%v\n", board)
	fmt.Printf("%v\n", nbr) */

	for i := range board {
		for _, val := range board[i] {
			if val != -1 {
				totalBoard += val
			}
		}
	}

	return nbr * totalBoard
}

func checkIfWinner(boards map[int]*[5][5]int) (bool, int) {
	for i := range boards {
		// check all rows
		for j := range boards[i] {
			if checkIfWinnerColl(boards[i][j]) {
				return true, i
			}
		}

		transposed := transpose(boards[i])
		for j := range transposed {
			if checkIfWinnerColl(transposed[j]) {
				return true, i
			}
		}
	}

	return false, -1
}

func transpose(board *[5][5]int) *[5][5]int {
	var newBoard [5][5]int

	for i := range board {
		for j := range board[i] {
			newBoard[j][i] = board[i][j]
		}
	}

	return &newBoard
}

func checkIfWinnerColl(coll [5]int) bool {
	win := true

	for _, val := range coll {
		if val != -1 {
			win = false
			break
		}
	}

	return win
}

func markBoards(boards map[int]*[5][5]int, number int) map[int]*[5][5]int {
	internal := boards
	for i := range internal {
		for j := range internal[i] {
			for k, b := range internal[i][j] {
				if b == number {
					internal[i][j][k] = -1
				}
			}
		}
	}

	return internal
}

func makeBoards(str map[int][]int) map[int]*[5][5]int {
	boards := make(map[int]*[5][5]int)

	for i := range str {
		var b [5][5]int

		for j, val := range str[i] {
			b[j/5][j%5] = val
		}

		boards[i] = &b
	}

	return boards
}

func readInput(r io.Reader) (map[int][]int, error) {
	res := make(map[int][]int, 0)
	line := 0

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		x := scanner.Text()
		if x == "" {
			line++
			continue
		}
		if line == 0 {
			res[line] = append(res[line], stringToInt(strings.Split(x, ","))...)
			continue
		}
		res[line] = append(res[line], stringToInt(strings.Split(x, " "))...)
	}

	return res, scanner.Err()
}

func stringToInt(str []string) []int {
	res := make([]int, 0)

	for _, val := range str {
		if val == "" || val == " " {
			continue
		}
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}

	return res
}
