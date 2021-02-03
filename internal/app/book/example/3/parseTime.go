package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s strings\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	var myTime string = os.Args[1]

	d, err := time.Parse("15:04", myTime)

	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("My time:", d.Hour(), d.Minute())
	} else {
		fmt.Println(err)
	}

}
