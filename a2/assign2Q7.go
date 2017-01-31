package main

import (
	"log"
	"sort"
)

func countOverlappingIntervals(intervals [][]float64) int {
	if len(intervals) <= 1 {
		return 0
	}
	mid := len(intervals) / 2
	leftCount := countOverlappingIntervals(intervals[:mid])
	rightCount := countOverlappingIntervals(intervals[mid:])

	count := 0
	i := 0
	j := mid
	for i < mid && j < len(intervals) {
		left := intervals[i][1]
		right := intervals[j][1]
		if right > left {
			i++
		} else {
			count += i + 1
			j++
		}
	}

	sorted := make([][]float64, len(intervals))

	// Merge sort by y
	i = 0
	j = mid
	// While there are elements in the left or right runs...
	for k := 0; k < len(intervals); k++ {
		// If left run head exists and is <= existing right run head.
		if i < mid && (j >= len(intervals) || intervals[i][1] <= intervals[j][1]) {
			sorted[k] = intervals[i]
			i++
		} else {
			sorted[k] = intervals[j]
			j++
		}
	}
	copy(intervals, sorted)

	return leftCount + rightCount + count
}

type byX [][]float64

func (s byX) Len() int      { return len(s) }
func (s byX) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byX) Less(i, j int) bool {
	a := s[i][0]
	b := s[j][0]
	if a == b {
		return s[i][1] > s[j][1]
	}
	return a < b
}

func main() {
	intervals := [][]float64{
		{0, 5},
		{0, 6},
		{0, 7},
		{0, 8},
	}

	sort.Sort(byX(intervals))
	log.Println(intervals)
	log.Println(countOverlappingIntervals(intervals))
	log.Println(intervals)
}
