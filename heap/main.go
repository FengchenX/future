

package main

import (
	"fmt"
	"container/heap"

)

func main() {
	Example_PriorityQueue()
}

type Item struct {
	value string
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority>pq[j].priority
}

func (pq PriorityQueue) Swap(i,j int) {
	pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n:=len(*pq)
	item:=x.(*Item)
	item.index=n
	*pq= append(*pq,item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old:=*pq
	n:=len(old)
	item:=old[n-1]
	item.index=-1
	*pq=old[0:n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item,value string,priority int) {
	item.value=value
	item.priority=priority
	heap.Fix(pq,item.index)
}

func Example_PriorityQueue(){
	var items = map[string]int {
		"apple":4,"banana":2,"pear":3,
	}
	var pq =make(PriorityQueue,len(items))
	i:=0
	for value,priority := range items {
		pq[i]=&Item{
			value: value,
			priority: priority,
			index: i,
		}
		i++
	}
	item:=&Item{
		value: "oriange",
		priority: 1,
	}
	pq.Push(item)
	pq.update(item,item.value,5)
	for pq.Len()>0 {
		item:=heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ",item.priority,item.value)
	}
}

