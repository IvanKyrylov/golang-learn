package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	f := os.Stdin
	defer f.Close()

	s := bufio.NewScanner(f)
	var r []int
	for s.Scan() {
		if s.Text() == "END" {
			return
		}
		i, err := strconv.Atoi(s.Text())
		if err == nil {
			r = append(r, i)
		}
	}
}
