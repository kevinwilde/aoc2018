package main

import (
	"aoc2018/day7/priorityqueue"
	"aoc2018/utils"
	"fmt"
	"log"
	"regexp"
)

func main() {
	input := utils.GetInput("input.txt")
	fmt.Println(part1(convertInputToDependencyTree(input, 0)))
	fmt.Println(part2(convertInputToDependencyTree(input, 60), 5))
}

type task struct {
	id        string
	duration  int
	blockedBy int
	blocks    []*task
}

type worker struct {
	currentTask *task
	startedAt   int
}

func part1(deps map[string]*task) string {
	pq := priorityqueue.NewPriorityQueue()

	// add initial unblocked tasks to queue
	for id, t := range deps {
		if t.blockedBy == 0 {
			pq.Push(id)
		}
	}

	res := ""

	for pq.Len() > 0 {
		curTaskID := pq.Pop()
		res = res + curTaskID
		curTask := deps[curTaskID]
		for _, t := range curTask.blocks {
			t.blockedBy--
			if t.blockedBy == 0 {
				pq.Push(t.id)
			}
		}
	}

	return res
}

func part2(deps map[string]*task, numWorkers int) int {
	var workers []*worker
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, &worker{nil, 0})
	}

	pq := priorityqueue.NewPriorityQueue()

	// add initial unblocked tasks to queue
	for id, t := range deps {
		if t.blockedBy == 0 {
			pq.Push(id)
		}
	}

	time := 0
	for pq.Len() > 0 || workerInProgressOnTask(workers) {
		// check if there are any tasks in pq
		// if so, look for an available worker to assign it to
		available := availableWorker(workers)
		for pq.Len() > 0 && available != nil {
			available.currentTask = deps[pq.Pop()]
			available.startedAt = time
			available = availableWorker(workers)
		}

		time++

		// check if any workers have finished a task
		// if so, add anything that was blocked by only that task to pq
		// and set that worker as available
		for _, w := range workers {
			if w.currentTask != nil && w.currentTask.duration+w.startedAt == time {
				for _, t := range w.currentTask.blocks {
					t.blockedBy--
					if t.blockedBy == 0 {
						pq.Push(t.id)
					}
				}
				w.currentTask = nil
			}
		}
	}

	return time
}

func convertInputToDependencyTree(input []string, baseTaskDuration int) map[string]*task {
	reg := regexp.MustCompile(`^Step (.+) must be finished before step (.+) can begin.$`)

	tasks := make(map[string]*task)

	for _, s := range input {
		match := reg.FindStringSubmatch(s)
		if len(match) < 1 {
			log.Fatal("Input did not match regex")
		}

		for _, taskID := range match[1:] {
			if tasks[taskID] == nil {
				t := &task{
					taskID,
					int(byte(taskID[0])-'A') + baseTaskDuration + 1,
					0,
					make([]*task, 0),
				}
				tasks[taskID] = t
			}
		}
	}

	for _, s := range input {
		match := reg.FindStringSubmatch(s)
		if len(match) < 1 {
			log.Fatal("Input did not match regex")
		}

		tasks[match[1]].blocks = append(tasks[match[1]].blocks, tasks[match[2]])
		tasks[match[2]].blockedBy++
	}

	return tasks
}

func workerInProgressOnTask(workers []*worker) bool {
	for _, w := range workers {
		if w.currentTask != nil {
			return true
		}
	}
	return false
}

func availableWorker(workers []*worker) *worker {
	for _, w := range workers {
		if w.currentTask == nil {
			return w
		}
	}
	return nil
}
