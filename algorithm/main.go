package main

import "fmt"

func main() {
	objArr := []int{2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3}
	//bubbleSort(sortObj)
	//insertionSort(sortObj)
	fmt.Println(binarySearchFirstValue(objArr, 0, len(objArr)-1, 2))
	//fmt.Println(binarySearchLastValue(objArr, 0, len(objArr)-1, 3))
}

