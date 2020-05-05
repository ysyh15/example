package main

import "fmt"

/**
冒泡排序
*/
func bubbleSort(sortObj []int) {
	//如果数组只有一个数，则不需要排序
	if len(sortObj) < 1 {
		return
	}
	//循环需要排序的数据
	for n := 0; n < len(sortObj)-1; n++ {
		breakFlag := false
		//将本次排序的值与所有值进行循环比较
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
		//记录当前要插入的值
		value := sortObj[n]
		//标记当前位置的前一个位置
		m := n - 1
		//标记应当插入的地址
		insertPlace := 0
		//从后往前开始遍历比较
		for ; m >= 0; m-- {
			//如果要插入的值小于循环比对的值，
			if value < sortObj[m] {
				//将循环比对的值后移一位
				sortObj[m+1] = sortObj[m]
			} else {
				//没有比要插入的值小的，代表上一个循环的值就是要插入的位置，因为上一个循环的值已经被移到后面去了
				insertPlace = m + 1
				break
			}
		}
		//将值插入
		sortObj[insertPlace] = value
	}

}
/**
归并排序
 */
func mergeSort(mergeSort []int){

}

/**
快速排序
 */
func quickSort(quickSort []int){

}

/**
桶排序
 */


/**
计数排序
 */