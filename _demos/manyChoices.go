package main

// go run manyChoices.go

import (
	"fmt"
	"github.com/geeks-dev/go-choices"
	"strconv"
)

func main() {
	array := []string{}
	for index := 0; index < 50000; index++ {
		array = append(array, strconv.Itoa(index)+". data")
	}

	res, answered := choices.Ask(array, "Many choices")

	if answered {
		fmt.Println(res)
	}
}
