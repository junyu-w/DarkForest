package models

type Civilization struct {
	Id int
	Position *coordinate
	Type string
	NumYears int
	Range float64 // unit: lightyear
	MatterOwned int
	Level level
}

type coordinate struct {
	x int
	y int
}

type level int
const (
	DESOLATE level = 1 + iota
	EARTH
	TRISOLARIS
	DIMENSTION_TECH
	MINI_UNIVERSE
)

const (
	CONQUERER string = "conquerer"
	CONSERVATIVE string = "conservative"
)

func NewCivilization(id int, pos *coordinate, category string) *Civilization {
	return &Civilization {
		Id: id,
		Position: pos,
		Type: category,
		NumYears: 0,
		Range: 1e-10,
		Level: DESOLATE,
	}
}