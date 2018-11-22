package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Printing randoms.")
	const modifier int = 100
	var nums [modifier]int
	numOfTests := len(nums) * modifier
	for i := 0; i < numOfTests; i++ {
		rand := getRandom(100)
		nums[rand]++
		if i%(numOfTests/10) == 0 {
			fmt.Print(i/(numOfTests/10)*10+10, "% ")
		}
	}
	fmt.Println()
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, "->", nums[i])
	}
}

func getRandom(max int) int {
	// create and start helpers
	bufferSize := max
	helperSize := time.Now().Second() / 2
	if helperSize == 0 {
		helperSize = 1
	}
	var helpers = make([]chan int, helperSize, helperSize)
	for i := 0; i < helperSize; i++ {
		helpers[i] = make(chan int, bufferSize)
	}
	for i := 0; i < len(helpers); i++ {
		go randHelperRoutine(helpers[i], bufferSize)
	}
	// get helper values
	var values = make([]int, helperSize, helperSize)
	for i := 0; i < max; i++ {
		for j := 0; j < len(helpers); j++ {
			select {
			case temp := <-helpers[j]:
				values[j] += temp
			}
		}
	}
	// get sum
	sum := 0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}

	return int(math.Abs(float64(sum * time.Now().Nanosecond() % max)))
}

func randHelperRoutine(helper chan int, iterations int) {
	for i := 1; i < iterations+1; i++ {
		helper <- i * time.Now().Nanosecond()
	}
	close(helper)
}
