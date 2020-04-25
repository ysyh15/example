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
	if len(sortObj) < 1 {
		return
	}
	//只需要从第二个数开始比较
	for n := 1; n < len(sortObj); n = n + 1 {
		//记录当前的值
		value := sortObj[n]
		//标记前一个位置
		m := n - 1
		//从后往前开始遍历比较
		for ; m >= 0; m-- {
			//如果当前值小于循环比对的值，
			if sortObj[m] > value {
				sortObj[m+1] = sortObj[m]
				fmt.Println(m)
			} else {
				//没有比value小的，就已经标记处当前m的位置了
				break
			}
		}
		fmt.Println(m)
		//m的值在经过循环后-1插入数据
		sortObj[m+1] = value
	}

	fmt.Println(sortObj)
}
