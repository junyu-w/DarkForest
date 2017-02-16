package models

import (
	"sort"
	"math"
)

type Coordinate struct {
	x int64
	y int64
}

type ByDistance []*Coordinate

func (s ByDistance) Len() int {
	return len(s)
}

func (s ByDistance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByDistance) Less(i, j int) bool {
	orig := s[len(s)-1]
	dist_i := GetDistance(orig, s[i])
	dist_j := GetDistance(orig, s[j])
	return dist_i < dist_j
}

func GetDistance(a, b *Coordinate) float64 {
	x_diff := a.x - b.x
	y_diff := a.y - b.y
	return math.Sqrt(float64(x_diff*x_diff + y_diff*y_diff))
}


func SortByDistance(orig *Coordinate, others []*Coordinate) []*Coordinate {
	all := append(others, orig)
	sort.Sort(ByDistance(all))
	return all[1:len(all)]
}