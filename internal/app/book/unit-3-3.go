package main

import "fmt"

const (
	p4_0 = 1 << iota << iota
	p4_1
	p4_2
	p4_3
	p4_4
	p4_5
	p4_6
	p4_7
	p4_8
	p4_9
)

func main() {
	fmt.Printf("p4_0: %d, p4_1: %d, p4_2: %d, p4_3: %d, p4_4: %d, p4_5: %d, p4_6: %d, p4_7: %d, p4_8: %d, p4_9: %d",
		p4_0, p4_1, p4_2, p4_3, p4_4, p4_5, p4_6, p4_7, p4_8, p4_9)
}
