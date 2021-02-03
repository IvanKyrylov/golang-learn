package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Need more arguments")
	}

	inDate := os.Args

	for i := 1; i < len(inDate); i++ {
		d, err := time.Parse("Jan 2 2006", inDate[i])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Full: %v\n", d)
		fmt.Printf("My date: %v %v %v\n", d.Year(), d.Month(), d.Day())
	}
}
