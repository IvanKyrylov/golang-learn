package main

import "fmt"

const (
	_ = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Printf(`Monday: %d day, Tuesday: %d day, Wednesday: %d day, Thursday: %d day,
Friday: %d day, Saturday: %d day, Sunday: %d day`, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}
