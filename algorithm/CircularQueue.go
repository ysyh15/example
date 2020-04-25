package main

import "fmt"
//循环队列
type circularQueue struct {
	items []string //数组
	n     int      //数组大小
	head  int      //头
	tail  int      //尾部
}

func (cir *circularQueue) enqueue(item string) (succ bool) {
	if (cir.tail+1)%cir.n == cir.head {
		return false
	}
	cir.items[cir.tail] = item
	cir.tail = (cir.tail + 1) % cir.n
	return true
}

func (cir *circularQueue) dequeue() (item string) {
	if cir.tail == cir.head {
		return ""
	}
	item = cir.items[cir.tail]
	cir.head = (cir.head + 1) % cir.n
	fmt.Println(item)
	return item
}

