package utils

import "fmt"

func PrintPath(path []int) {
	for i := 0; i < len(path); i++ {
		fmt.Printf(" %d ", path[i])
	}
	fmt.Printf(" %d \n", path[0])
}

func Contains[K comparable](slice []K, value K) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}
