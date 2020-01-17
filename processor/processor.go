package processor

import (
	"log"
	"strconv"
)

func Process(config, data []string) []int {
	maxSlices, err := strconv.Atoi(config[0])
	if err != nil {
		log.Fatal(err)
	}

	var pizzen []int
	var currentSlice int = 0
	for key, pizza := range data {
		numberPizza, err := strconv.Atoi(pizza)

		if err != nil {
			log.Fatal(err)
		}

		if (currentSlice + numberPizza) < maxSlices{
			currentSlice += numberPizza
			pizzen = append(pizzen, key)
		}
	}

	return pizzen
}