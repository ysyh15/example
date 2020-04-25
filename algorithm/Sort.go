package main

import "fmt"

/**
冒泡排序
*/
func bubbleSort(sortObj []int) {
	if len(sortObj) < 1 {
		return
	}
	for n := 0; n < len(sortObj)-1; n++ {
		breakFlag := false
		for m := 0; m < len(sortObj)-1; m++ {
			tmp := 0
			if sortObj[m] > sortObj[m+1] {
				tmp = sortObj[m]
				sortObj[m] = sortObj[m+1]
				sortObj[m+1] = tmp
				breakFlag = true
			}
		}
		if !breakFlag {
			break
		}
	}
	fmt.Println(sortObj)
}

/**
插入排序
*/
func insertionSort(sortObj []int) {
	fmt.Println(sortObj)
}
