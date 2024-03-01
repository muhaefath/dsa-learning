package main

import (
	"sort"
)

// k => number of workers
// task => duration of task must be compelted by the workers
// each worker need complete 2 task
// each worker only can work on one task at a time
// the number task will awalys equal to 2k
// all task are independent
// workers will complete their task in parerel
// return optimal assignment task to each workers
//
// 1. initiate resp [][]int
// 2. order task
// 3. loop task
// 4. initiate taskIndex := []int
// 5. append taskIndex = task[I]
// 6. append taskIndex = task[len(array)-i]
// 7. return resp
func TaskAssignment(k int, tasks []int) [][]int {
	resp := [][]int{}
	tempTask := tasks

	for _, v := range tasks {
		tempTask = append(tempTask, v)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})

	for i := 0; i < k; i++ {
		taskA := 0
		taskB := 0

		taskA, tempTask = FindIndex(tasks[i], tempTask)
		taskB, tempTask = FindIndex(tasks[len(tasks)-1-i], tempTask)
		taskIndex := []int{
			taskA,
			taskB,
		}

		resp = append(resp, taskIndex)
	}

	return resp
}

func FindIndex(value int, tempTask []int) (int, []int) {
	for idx, v := range tempTask {
		if v == value {
			tempTask[idx] = -1
			return idx, tempTask
		}
	}

	return 0, tempTask
}
