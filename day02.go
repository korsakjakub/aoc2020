package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("day02.txt")
	if err != nil {
		panic(err)
	}
	txt := string(dat)
	lines := strings.Split(txt, "\n")
	valid := 0
	for _, l := range lines {
		if solve(l) {
			valid += 1
		}
		fmt.Println(valid)
	}
}

func solve(line string) bool {
	str := strings.Split(line, " ")
	r := strings.Split(str[0], "-")
	min, _ := strconv.Atoi(r[0])
	max, _ := strconv.Atoi(r[1])
	letter := []rune(strings.Replace(str[1], ":", "", -1))[0]
	sequence := str[2]

	return isValid2(min, max, letter, sequence)
}

func isValid(min int, max int, letter rune, sequence string) bool {
	count := 0
	for _, s := range sequence {
		if s == letter {
			count += 1
		}
	}
	if count <= max && count >= min {
		return true
	}
	return false
}

func isValid2(min int, max int, letter rune, sequence string) bool {
	if getNth(sequence, min - 1) == letter || getNth(sequence, max - 1) == letter {
		if getNth(sequence, min - 1) == letter && getNth(sequence, max - 1) == letter {
			return false
		}
		return true
	}
	return false
}

func getNth(str string, n int) rune {
	for i, el := range str {
		if i == n {
			return el
		}
	}
	return '0'
}