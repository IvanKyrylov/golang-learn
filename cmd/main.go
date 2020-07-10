package main

import (
	"fmt"

	"github.com/IvanKyrylov/goland-learn/internal/app/homework/taskone"
	"github.com/IvanKyrylov/goland-learn/internal/app/homework/tasktwo"
)

func main() {
	taskone.CorrectTime()

	t, err := tasktwo.UnpackingString("4c")

	fmt.Printf("%v\t%v", t, err)

}
