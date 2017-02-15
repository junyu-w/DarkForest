package models

import (
	"fmt"
	"math/rand"
)

// assume a fixed size universe
const (
	WIDTH int64= 10e6
	HEIGHT int64 = 10e6
)

type Universe struct {
	width int64
	height int64

	Duration int // unit: year
	TotalMatter float64
	ContainedCivilizations map[*Coordinate]*Civilization
}

const (
	TOTAL_MATTER float64 = 10e10
)

func NewUniverse() *Universe {
	return &Universe {
		width: WIDTH,
		height: HEIGHT,
		Duration: 0,
		TotalMatter: TOTAL_MATTER,
		ContainedCivilizations: make(map[*Coordinate]*Civilization),
	}
}

func (u *Universe) GetArea() float64 {
	return float64(u.width * u.height)
}

/**
 * this function evolves the universe, and existing
 * civilizations, and create new civlization based on
 * randomeness
 */
func (u *Universe) Evovle(num_year int) {
	u.Duration += 1
	for pos, civil := range u.ContainedCivilizations {
		civil.Evovle(num_year)
		fmt.Println("Civilization ", civil.Id, " at position ", pos, " has evovled")
	}
	shouldCreateUniverse := rand.Intn(10) > 5 // TODO: this should change to something based on num of existing civil
	if shouldCreateUniverse {
		new_pos := &Coordinate {
			x: rand.Int63n(WIDTH),
			y: rand.Int63n(HEIGHT),
		}
		id := len(u.ContainedCivilizations) + 1
		new_c := NewCivilization(id, new_pos, CONSERVATIVE, u) // TODO: create conservative civilization for now
		u.ContainedCivilizations[new_pos] = new_c
		fmt.Println("Civilization ", id, " is created!!")
	}
}

