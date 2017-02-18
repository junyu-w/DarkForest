package models

import (
	"dark_forest/utils"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"math/rand"
)

type Universe struct {
	width  int64
	height int64

	NumYears               int // unit: year
	TotalMatter            float64
	ContainedCivilizations map[*Coordinate]*Civilization
}

func NewUniverse() *Universe {
	return &Universe{
		width:                  utils.WIDTH,
		height:                 utils.HEIGHT,
		NumYears:               0,
		TotalMatter:            utils.TOTAL_MATTER,
		ContainedCivilizations: make(map[*Coordinate]*Civilization),
	}
}

func (u *Universe) GetArea() float64 {
	return float64(u.width * u.height)
}

func (u *Universe) GetNearbyCivilizations(c *Civilization, limit int) []*Civilization {
	all_pos := make([]*Coordinate, 0, len(u.ContainedCivilizations)+1)
	for pos, _ := range u.ContainedCivilizations {
		all_pos = append(all_pos, pos)
	}
	sorted_pos := SortByDistance(c.Position, all_pos)
	nearby_civils := make([]*Civilization, limit, limit)
	for i := 0; i < limit; i++ {
		nearby_civils[i] = u.ContainedCivilizations[sorted_pos[i]]
	}
	return nearby_civils
}

/**
 * update universe & civil and reflect on UI
 */
func (u *Universe) UpdateAndDrawUniverse(num_year int, screen *ebiten.Image) {
	u.Evovle(num_year)
	u.UpdateAndDrawCivilization(num_year, screen)
}

/**
 * this function evolves the universe, and existing
 * civilizations, and create new civlization based on
 * randomeness
 */
func (u *Universe) Evovle(num_year int) {
	u.NumYears += num_year
	fmt.Println("Universe has evovled ", u.NumYears, "years")
	// create civilization
	shouldCreateCivilization := rand.Intn(100) > 80 // TODO: this should change to something based on num of existing civil
	if shouldCreateCivilization {
		new_pos := &Coordinate{
			x: rand.Int63n(utils.WIDTH),
			y: rand.Int63n(utils.HEIGHT),
		}
		id := len(u.ContainedCivilizations) + 1
		new_c := NewCivilization(id, new_pos, utils.CONSERVATIVE, u) // TODO: create conservative civilization for now
		u.ContainedCivilizations[new_pos] = new_c
		// start message receiving process
		fmt.Println("[CREATE] Civilization ", id, " is created at position ", new_c.Position)
	}
}

func (u *Universe) UpdateAndDrawCivilization(num_year int, screen *ebiten.Image) {
	for _, civil := range u.ContainedCivilizations {
		civil.Evovle(num_year)
		if len(u.ContainedCivilizations) > 50 && civil.ChooseToRevealPosition() {
			civil.BroadcastPosition()
		}
		civil.Draw(screen)
	}
}
