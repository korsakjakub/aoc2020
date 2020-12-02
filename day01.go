package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("day01.txt")
	if err != nil {
		panic(err)
	}
	txt := string(dat)
	lines := strings.Split(txt, "\n")
	var values []int
	for _, el := range lines {
		val, _ := strconv.Atoi(el)
		values = append(values, val)
	}

	for _, e := range values {
		for _, f := range values {
			for _, g := range values {
				if e+f+g == 2020 {
					fmt.Println(e, f, g)
					fmt.Println(e * f * g)
				}
			}
		}
	}
}
