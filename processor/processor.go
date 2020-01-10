package processor

import (
	"strconv"
)

func Process(config, data []string) []int {
	var currentSlice int = 0
	maxSlices, err := strconv.Atoi(config[0])
	if err != nil {
		panic(maxSlices)
	}

	pizzen := []int{}
	for key, pizza := range data {
		numberPizza, err := strconv.Atoi(pizza)

		if err != nil {
			panic(err)
		}

		if (currentSlice + numberPizza) < maxSlices{
			currentSlice += numberPizza
			pizzen = append(pizzen, key)
		}
	}

	return pizzen
}