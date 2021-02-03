package main

import "fmt"

func arrayToMap(slice []interface{}) map[int]interface{} {

	tempMap := make(map[int]interface{})

	for i, v := range slice {
		tempMap[i] = v
	}
	return tempMap
}

func main() {

	testcase := [...]struct {
		i int
		s string
		b bool
	}{
		{i: 1, s: "2", b: false},
		{i: 3, s: "4", b: true},
	}

	tempSlice := make([]interface{}, len(testcase))

	for i, v := range testcase {
		tempSlice[i] = v
	}

	fmt.Printf("%v", arrayToMap(tempSlice))
}
