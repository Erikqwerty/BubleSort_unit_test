package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generationArray(lenArr int, numberRange int) []int {
	var array []int = make([]int, lenArr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(numberRange)
	}
	return array
}

//
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for x := i; x < len(arr); x++ {
			var tmp int
			if arr[i] > arr[x] {
				tmp = arr[i]
				arr[i] = arr[x]
				arr[x] = tmp
			}
		}
	}
}

func main() {
	var array []int = generationArray(20, 100)
	BubbleSort(array)
	fmt.Println(array)

}
