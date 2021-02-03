package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		return
	}

	var myDate string = os.Args[1]

	d, err := time.Parse("02 Jan 2006", myDate)

	if err == nil {
		fmt.Printf("Full: %v\n", d)
		fmt.Printf("Time: %v %v %v\n", d.Day(), d.Month(), d.Year())
	} else {
		fmt.Println(err)
	}

}
