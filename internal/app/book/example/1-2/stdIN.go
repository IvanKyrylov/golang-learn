package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scaner := bufio.NewScanner(f)
	for scaner.Scan() {
		fmt.Println(">", scaner.Text())
	}
}
