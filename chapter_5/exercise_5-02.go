package main

import (
	"fmt"
	"time"
)

// There are no classes in Go, data hiding is achieved through exporting

type automobile struct {
	manufacturer string
	model        string
	year         int
}

func (a automobile) getDescription() string {
	return fmt.Sprintf("%d %s %s", a.year, a.manufacturer, a.model)
}

func (a automobile) getAge() int {
	return time.Now().Year() - a.year
}

func main() {
	car := automobile{
		manufacturer: "Honda",
		model:        "City",
		year:         2009,
	}

	fmt.Println(car.getDescription(), "\b, currently", car.getAge(), "years old.")
}
