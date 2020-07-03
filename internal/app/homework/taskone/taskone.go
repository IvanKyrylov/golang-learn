package taskone

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const server = "3.pool.ntp.org"

func CorrectTime() time.Time {
	time, err := ntp.Time(server)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return time
	}
	fmt.Printf("Time: %v", time)
	return time
}
