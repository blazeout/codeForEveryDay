package main

import "fmt"

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	destinationSun := 0
	destinationYi := 0
	sum := 0
	// 顺时针
	for i := 0; i < len(distance); i++ {
		sum += distance[i]
	}
	if start < destination {
		for i := start; i < destination; i++ {
			destinationSun += distance[i]
		}
	} else {
		for i := start - 1; i >= destination; i-- {
			destinationSun += distance[i]
		}
	}

	destinationYi = sum - destinationSun
	if destinationSun < destinationYi {
		return destinationSun
	}
	return destinationYi
}

func main() {
	distance := []int{7, 10, 1, 12, 11, 14, 5, 0}
	start := 7
	destination := 2
	fmt.Println(distanceBetweenBusStops(distance, start, destination))
}
