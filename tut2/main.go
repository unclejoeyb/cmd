package main

import "fmt"

func removeFromSlice(slice []int, index int) []int{
	slice[index] = slice[len(slice) - 1]
	return slice[:len(slice) - 1]
}

func removeFromSliceWithOrder(slice []int, index int) []int{
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(removeFromSlice(nums, 2))
	fmt.Println(removeFromSliceWithOrder(nums, 2))
}	