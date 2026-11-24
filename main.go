package main

import (
	"fmt"
	"sort"
)

// Job is a unit-length job with a deadline (1-based day) and a profit,
// earned only if the job runs on some day ≤ its deadline.
type Job struct {
	ID       int
	Deadline int
	Profit   int
}

// Schedule returns how many jobs to run and the total profit, taking jobs
// in descending profit and slotting each into the latest free day within
// its deadline.
func Schedule(jobs []Job) (count, profit int) {
	sorted := append([]Job{}, jobs...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Profit > sorted[j].Profit })
	maxDay := 0
	for _, j := range sorted {
		if j.Deadline > maxDay {
			maxDay = j.Deadline
		}
	}
	slot := make([]bool, maxDay+1) // slot[d]: day d taken (1-based)
	for _, j := range sorted {
		for d := j.Deadline; d >= 1; d-- {
			if !slot[d] {
				slot[d] = true
				count++
				profit += j.Profit
				break
			}
		}
	}
	return count, profit
}

func main() {
	jobs := []Job{{1, 4, 20}, {2, 1, 10}, {3, 1, 40}, {4, 1, 30}}
	c, p := Schedule(jobs)
	fmt.Println(c, p)
}
