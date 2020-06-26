package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("3.pool.ntp.org")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Println(time)
}
