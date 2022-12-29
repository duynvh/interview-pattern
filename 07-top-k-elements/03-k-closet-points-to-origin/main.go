package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func kClosest(points []Point, k int) []Point {
	pointsMaxHeap := newMaxHeap()
	res := make([]Point, 0)

	for i := 0; i < k; i++ {
		heap.Push(pointsMaxHeap, points[i])
	}

	for i := k; i < len(points); i++ {
		top := pointsMaxHeap.Top().(Point)
		if points[i].Less(top) {
			heap.Pop(pointsMaxHeap)
			heap.Push(pointsMaxHeap, points[i])
		}
	}

	for !pointsMaxHeap.Empty() {
		res = append(res, heap.Pop(pointsMaxHeap).(Point))
	}

	return res
}

func pointArrayToString(points []Point) string {
	str := "["
	for i, p := range points {
		str += p.String()
		if i < len(points)-1 {
			str += ", "
		}
	}
	str += "]"
	return str
}

// Driver Code
func main() {
	pointsOne := []Point{Point{1, 3}, Point{3, 4}, Point{2, -1}}
	pointsTwo := []Point{Point{1, 3}, Point{2, 4}, Point{2, -1}, Point{-2, 2}, Point{5, 3}, Point{3, -2}}
	pointsThree := []Point{Point{1, 3}, Point{5, 3}, Point{3, -2}, Point{-2, 2}}
	pointsFour := []Point{Point{2, -1}, Point{-2, 2}, Point{1, 3}, Point{2, 4}}
	pointsFive := []Point{Point{1, 3}, Point{2, 4}, Point{2, -1}, Point{-2, 2}, Point{5, 3}, Point{3, -2}, Point{5, 3}, Point{3, -2}}
	kList := []int{2, 3, 1, 4, 5}
	points := [][]Point{pointsOne, pointsTwo, pointsThree, pointsFour, pointsFive}
	for i, point := range points {
		fmt.Printf("%d.\tSet of points: %s\n", i+1, pointArrayToString(point))
		fmt.Printf("\tk = %d\n", kList[i])
		result := kClosest(point, kList[i])
		fmt.Printf("\tHere are the k = %d points closest to the origin (0, 0): %s\n", kList[i], pointArrayToString(result))
		fmt.Printf("%s\n", strings.Repeat("-", 95))
	}
}
