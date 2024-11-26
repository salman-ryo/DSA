package main

import "fmt"

func GetMinMax(arr []int32) (int32, int32) {
	if len(arr) == 0 {
		return -1, -1
	}
	minimum := arr[0]
	maximum := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < minimum {
			minimum = arr[i]
		}
		if arr[i] > maximum {
			maximum = arr[i]
		}
	}
	return minimum, maximum

}

func main() {
	fmt.Println("Array in the memory")
	mymin, myMax := GetMinMax([]int32{8, 1, 2, 3, 4, 5})
	fmt.Printf("Min is: %d, Max is: %d", mymin, myMax)
}
