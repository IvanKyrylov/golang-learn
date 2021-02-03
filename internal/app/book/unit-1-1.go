package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("Need more arguments")
	}

	arg := os.Args
	var sum int
	for i := 1; i < len(arg); i++ {
		num, err := strconv.Atoi(arg[i])

		if err != nil {
			log.Fatal(err)
		}
		sum += num

	}
	fmt.Printf("Sum: %v", sum)
}
