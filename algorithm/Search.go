package main

//二分查找，没找到返回-1，数组长度必须大于1
func binarySearch(searchObj []int, low int, high int, searchVal int) (index int) {
	//可用移位运算 来优化
	var middle int
	if low == high {
		//最后一次比较
		if searchObj[low] != searchVal {
			return -1
		} else {
			return low
		}
	} else if low+1 == high {
		middle = low
	} else {
		middle = high - (high-low)/2
	}

	if searchObj[middle] > searchVal {
		return binarySearch(searchObj, low, middle-1, searchVal)
	} else if searchObj[middle] == searchVal {
		return middle
	} else {
		return binarySearch(searchObj, middle+1, high, searchVal)
	}
}

//二分查找第一个值
func binarySearchFirstValue(searchObj []int, low int, high int, firstVal int) (index int) {
	//获取中间数
	var middle int

	if low == high {
		//最后一次比较
		if searchObj[low] != firstVal {
			return -1
		} else {
			return low
		}
	} else if low+1 == high {
		middle = low
	} else {
		middle = high - (high-low)/2
	}

	if searchObj[middle] == firstVal {
		//此时要确认是否为第一个数
		if low == middle {
			return middle
		} else {
			return binarySearchFirstValue(searchObj, low, middle, firstVal)
		}
	} else if searchObj[middle] > firstVal {
		return binarySearchFirstValue(searchObj, low, middle-1, firstVal)
	} else {
		return binarySearchFirstValue(searchObj, middle+1, high, firstVal)
	}
}

//二分查找最后个值
func binarySearchLastValue(searchObj []int, low int, high int, lastVal int) (index int) {
	//获取中间数
	var middle int

	if low == high {
		//最后一次比较
		if searchObj[low] != lastVal {
			return -1
		} else {
			return low
		}
	} else {
		middle = high - (high-low)/2
	}

	if searchObj[middle] == lastVal {
		//此时要确认是否为最后一个数
		if high == middle {
			return middle
		} else {
			return binarySearchLastValue(searchObj, middle, high, lastVal)
		}
	} else if searchObj[middle] > lastVal {
		return binarySearchLastValue(searchObj, low, middle-1, lastVal)
	} else {
		return binarySearchLastValue(searchObj, middle+1, high, lastVal)
	}
}
