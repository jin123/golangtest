package main

import (
	"bytes"
	"fmt"

	"github.com/bitly/go-simplejson"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people People

func JsonToStructDemo() {
	jsonStr := `
                        [
                {
                        "BizDate":"2020-04-25+00:00:00",
                        "Price":299.72,
                        "PriceCheck":"HBDr0XrG5FLqsX9+9WtgLVqJfa4RIjj6mIeb0tUN7YgBJ/1SbsGrVl4hskoqmTBj6c5ixfKKYk5qDyegz6R3anGMx09O6SSljSAli0oAYnKAqQuMCfGwUpa9W0c49Mqp54ep/CA5ETPvjOWWQKHixw=="
                },
                {
                        "BizDate":"2020-04-26+00:00:00",
                        "Price":299.72,
                        "PriceCheck":"RuYQb686vmcJe8CupDAKLTpVZIw2iPqxPaJQudPe6pEsPHoYalcjtIPl+RRCnZRq6c5ixfKKYk5qDyegz6R3ajLgXtx/X4q8PdqGATtJ5KKAqQuMCfGwUpa9W0c49Mqp54ep/CA5ETPvjOWWQKHixw=="
                }
                ]
                        `
	buf := bytes.NewBuffer([]byte(jsonStr))
	js, _ := simplejson.NewFromReader(buf)
	fmt.Println("js=", js)
}

func main() {
	JsonToStructDemo()
}
