package main

import "fmt"

// There are no classes in Go, data hiding is achieved through exporting

type automobile struct {
	manufacturer string
	model        string
	year         int
}

func main() {
	car := automobile{
		manufacturer: "Honda",
		model:        "City",
		year:         2009,
	}

	fmt.Println(car.manufacturer, car.model, car.year)
}
