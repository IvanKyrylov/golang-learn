package taskone

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

const server = "3.pool.ntp.org"

func CorrectTime() {
	time, err := ntp.Time(server)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Println(time)
}
