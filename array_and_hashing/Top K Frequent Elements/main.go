package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// fmt.Println("final: ", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println("final: ", topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
	// fmt.Println("final: ", topKFrequent([]int{1, 4, 3, 2, 2, 3, 3, 2, 3, 1}, 3))

}

// 1. initiate map [key]counter
// 2. initiate resp func []int and set the size with k
// 3. initiate counter min resp int
// 4. initiate index min resp int
// 5. loop nums
// 5.1 assign the num to map and counter it
// 5.2 compare to min resp, if more than will asign the value
// 6. return resp func
func topKFrequent(nums []int, k int) []int {
	mapNums := make(map[int]int)
	respFunc := []int{}
	counterMinReturn := 0
	indexMinReturn := 0

	for _, v := range nums {
		value, _ := mapNums[v]
		mapNums[v] = value + 1
	}

	for key, value := range mapNums {
		if counterMinReturn < value {
			if len(respFunc) < k {
				respFunc = append(respFunc, key)

				if len(respFunc) == k {
					counterMinReturn = mapNums[respFunc[0]]
					for i := 1; i < len(respFunc); i++ {
						if counterMinReturn > mapNums[respFunc[i]] {
							counterMinReturn = mapNums[respFunc[i]]
							indexMinReturn = i
						}
					}
				}

				fmt.Println("indexMinReturn: ", indexMinReturn)
				fmt.Println("counterMinReturn: ", counterMinReturn)
				fmt.Println("respFunc: ", respFunc)

			} else {
				respFunc[indexMinReturn] = key

				indexMinReturn, counterMinReturn = MappingSmallestNumber(mapNums, respFunc)

				fmt.Println("indexMinReturn2: ", indexMinReturn)
				fmt.Println("counterMinReturn2: ", counterMinReturn)
				fmt.Println("respFunc2: ", respFunc)
			}
		}
	}

	return respFunc
}

func MappingSmallestNumber(mapReq map[int]int, req []int) (int, int) {
	resp := mapReq[req[0]]
	index := 0
	for i := 1; i < len(req); i++ {
		if resp > mapReq[req[i]] {
			resp = mapReq[req[i]]
			index = i
		}
	}

	return index, resp
}

func topKFrequent2(nums []int, k int) []int {
	// Step 1: Count the frequency of each element in nums
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Step 2: Use a priority queue (min heap) to keep track of the k most frequent elements
	priorityQueue := &MinHeap{}
	heap.Init(priorityQueue)

	// Step 3: Populate the priority queue with the k most frequent elements
	for num, freq := range frequencyMap {
		heap.Push(priorityQueue, Element{num, freq})
		if priorityQueue.Len() > k {
			heap.Pop(priorityQueue) // Remove the least frequent element if the size exceeds k
		}
	}

	// Step 4: Extract the elements from the priority queue
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(priorityQueue).(Element).Value
	}

	return result
}

// Element represents an element in the priority queue with a value and frequency.
type Element struct {
	Value     int
	Frequency int
}

// MinHeap is a min-heap implementation for Element.
type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Frequency < h[j].Frequency }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
