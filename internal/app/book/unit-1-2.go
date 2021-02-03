package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatalln("Please give my one argument!")
	}

	arg := os.Args
	var r float64

	for i := 1; i < len(arg); i++ {
		n, err := strconv.ParseFloat(arg[i], 64)
		if err == nil {
			r += n
		}
	}
	r /= float64(len(arg) - 1)
	fmt.Println(r)

}
