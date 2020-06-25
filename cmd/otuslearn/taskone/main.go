package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("3.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %d", err)
		os.Exit(1)
	}
	fmt.Println(time)
}
