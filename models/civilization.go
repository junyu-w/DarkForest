package models

import (
	"math/rand"
)

type Civilization struct {
	Id int
	Position *Coordinate
	Type string
	NumYears int
	Range float64 // unit: lightyear
	MatterOwned float64
	Level level
	ContainerUniverse *Universe
}

type Coordinate struct {
	x int64
	y int64
}

type level int
const (
	DESOLATE level = 1 + iota
	EARTH
	TRISOLARIS
	DIMENSION_TECH
	MINI_UNIVERSE
)

const (
	CONQUERER string = "conquerer"
	CONSERVATIVE string = "conservative"
)

func NewCivilization(id int, pos *Coordinate, category string, universe *Universe) *Civilization {
	return &Civilization {
		Id: id,
		Position: pos,
		Type: category,
		NumYears: 0,
		Range: 10e-10,
		MatterOwned: 10e-5,
		Level: DESOLATE,
		ContainerUniverse: universe,
	}
}

/**
 * this function evolves civilization, including increase its
 * explored area (Range) and MatterOwned, which might cause its
 * civilization level to increase as well
 */
func (c *Civilization) Evovle(num_year int) {
	c.NumYears += 1
	c.MatterOwned += (c.Range * c.Range)/c.ContainerUniverse.GetArea() * float64(rand.Int31n(100))
	c.Range += 0.02 // TODO: let's not use fixed numebrs
	matterPercentage := c.MatterOwned / c.ContainerUniverse.TotalMatter
	c.Level = getLevel(matterPercentage)
}

func getLevel(matterPercentage float64) level {
	switch {
	case matterPercentage >= 10e-13:
		return DESOLATE
	case matterPercentage >= 10e-11:
		return EARTH
	case matterPercentage >= 10e-7:
		return TRISOLARIS
	case matterPercentage >= 10e-5:
		return DIMENSION_TECH
	case matterPercentage >= 10e-2:
		return MINI_UNIVERSE
	}
	return DESOLATE
}