package models

import (
	"dark_forest/utils"
	"math"
	"sort"
)

type Coordinate struct {
	x int64
	y int64
}

type CoordDistPair struct {
	coord, dist interface{}
}

type ByDistance []*CoordDistPair

func (s ByDistance) Len() int {
	return len(s)
}

func (s ByDistance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByDistance) Less(i, j int) bool {
	return s[i].dist.(float64) < s[j].dist.(float64)
}

func GetDistance(a, b *Coordinate) float64 {
	x_diff := a.x - b.x
	y_diff := a.y - b.y
	return math.Sqrt(float64(x_diff*x_diff + y_diff*y_diff))
}

func SortByDistance(orig *Coordinate, all []*Coordinate) []*Coordinate {
	all = append(all, orig)
	coord_dist_list := make([]*CoordDistPair, 0, len(all))
	for _, other := range all {
		dist := GetDistance(orig, other)
		coord_dist_list = append(coord_dist_list, &CoordDistPair{other, dist})
	}
	sort.Sort(ByDistance(coord_dist_list))
	sorted_coord := make([]*Coordinate, len(all), len(all))
	for i := 0; i < len(all); i++ {
		sorted_coord[i] = coord_dist_list[i].coord.(*Coordinate)
	}
	return sorted_coord[2:len(all)]
}

func (coord *Coordinate) TranslateToGameWindowPosition() (float64, float64) {
	scale_x := float64(utils.WIDTH) / (float64(utils.G_WIDTH))
	scale_y := float64(utils.HEIGHT) / (float64(utils.G_HEIGHT))
	return float64(coord.x) / scale_x, float64(coord.y) / scale_y
}
