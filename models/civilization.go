package models

import (
	"dark_forest/utils"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math/rand"
	"time"
)

type Civilization struct {
	Id                int
	Position          *Coordinate
	Type              string
	NumYears          int
	Range             float64 // unit: lightyear
	MatterOwned       float64
	Level             int
	ContainerUniverse *Universe
	Revealed          bool
	MessageChannel    chan CivilMessage
	Color             color.NRGBA
}

func NewCivilization(id int, pos *Coordinate, category string, universe *Universe) *Civilization {
	return &Civilization{
		Id:                id,
		Position:          pos,
		Type:              category,
		NumYears:          0,
		Range:             10e-2,
		MatterOwned:       10e-2,
		Level:             utils.LIGHTSPEED_x0001,
		ContainerUniverse: universe,
		Revealed:          false,
		MessageChannel:    make(chan CivilMessage),
		Color:             HIDDEN_COLOR,
	}
}

var REVEAL_COLOR color.NRGBA = color.NRGBA{0xff, 0x00, 0x00, 0xff}
var HIDDEN_COLOR color.NRGBA = color.NRGBA{0xff, 0xff, 0xff, 0xff}
var DISCOVER_COLOR color.NRGBA = color.NRGBA{0x7f, 0xff, 0x00, 0xff}
var DEATH_COLOR color.NRGBA = color.NRGBA{0x00, 0x00, 0x00, 0x00}

/**
 * this function evolves civilization, including increase its
 * explored area (Range) and MatterOwned, which might cause its
 * civilization level to increase as well
 */
func (c *Civilization) Evovle(num_year int) {
	c.NumYears += num_year
	c.MatterOwned += (c.Range * c.Range) / c.ContainerUniverse.GetArea() * float64(rand.Int31n(100))
	c.Range += 0.02 // TODO: let's not use fixed numebrs
	matterPercentage := c.MatterOwned / c.ContainerUniverse.TotalMatter
	c.Level = getLevel(matterPercentage)
}

func getLevel(matterPercentage float64) int {
	switch {
	case matterPercentage >= 10e-5:
		return utils.LIGHTSPEED_x0001
	case matterPercentage >= 10e-4:
		return utils.LIGHTSPEED_x001
	case matterPercentage >= 10e-3:
		return utils.LIGHTSPEED_x1
	case matterPercentage >= 10e-2:
		return utils.LIGHTSPEED_x2
	case matterPercentage >= 10e-1:
		return utils.LIGHTSPEED_x10
	}
	return utils.LIGHTSPEED_x0001
}

/**
 * Civilization chose to reveal itself in the dark forest
 */
func (c *Civilization) BroadcastPosition() {
	speed := getInfoSpeed(c.Level)
	c.Revealed = true
	c.Color = REVEAL_COLOR
	nearby_civils := c.ContainerUniverse.GetNearbyCivilizations(c, 10)
	c.ContainerUniverse.ChangeUniversalMessage(fmt.Sprintf("[REVEAL] Civilization %d choose to broadcast position", c.Id))
	for _, civil := range nearby_civils {
		dist := GetDistance(c.Position, civil.Position)
		arrival_time := int(dist/speed) + c.NumYears
		go c.SendMessage(arrival_time, civil.MessageChannel, c.Position)
		go civil.ProcessMessage()
	}
}

type CivilMessage interface{}

/**
 * Send message to another civilization via its message channel, this
 * simulates the late arrival of message using a timer
 */
func (c *Civilization) SendMessage(arrival_time int, channel chan CivilMessage, msg CivilMessage) {
	for {
		// fmt.Println("arrive at ", arrival_time, " now: ", c.NumYears)
		if c.NumYears >= arrival_time {
			channel <- msg
			break
		}
		time.Sleep(time.Millisecond * 20)
	}
}

// TODO: what to do after message arrived
func (civil *Civilization) ProcessMessage() {
	for {
		info := <-civil.MessageChannel
		switch info.(type) {
		case *Coordinate:
			civil.ContainerUniverse.ChangeUniversalMessage(fmt.Sprintf("[DISCOVER] Civilization %d Found (%d, %d)", civil.Id, info.(*Coordinate).x, info.(*Coordinate).y))
			if civil.Revealed == false {
				civil.Color = DISCOVER_COLOR
				civil.InitiateDarkForestAttack(info.(*Coordinate))
			}
		case *DarkForestAttack:
			dfa := info.(*DarkForestAttack)
			if err := dfa.Execute(); err == ALREADY_DESTROTYED {
				dfa.attacker.Color = HIDDEN_COLOR
				continue
			}
			civil.ContainerUniverse.ChangeUniversalMessage(fmt.Sprintf("[DFA] Civilization %d Destroyed %d", dfa.attacker.Id, civil.ContainerUniverse.GetCivilAtPosition(dfa.victim_pos).Id))
		}
		time.Sleep(time.Millisecond * 20)
	}
}

func (civil *Civilization) InitiateDarkForestAttack(coord *Coordinate) {
	target_civil := civil.ContainerUniverse.GetCivilAtPosition(coord)
	speed := getInfoSpeed(civil.Level)
	dist := GetDistance(civil.Position, coord)
	arrival_time := int(dist/speed) + civil.NumYears
	dfa := NewDFAttack(civil, coord)
	go civil.SendMessage(arrival_time, target_civil.MessageChannel, dfa)
}

func getInfoSpeed(civil_level int) float64 {
	switch civil_level {
	case utils.LIGHTSPEED_x0001:
		return utils.LIGHTSPEED * 0.0001
	case utils.LIGHTSPEED_x001:
		return utils.LIGHTSPEED * 0.001
	case utils.LIGHTSPEED_x1:
		return utils.LIGHTSPEED
	case utils.LIGHTSPEED_x2:
		return utils.LIGHTSPEED * 2
	case utils.LIGHTSPEED_x10:
		return utils.LIGHTSPEED * 10
	}
	return utils.LIGHTSPEED * 0.0001
}

func (civil *Civilization) ChooseToRevealPosition() bool {
	rand_num := rand.Int31n(100000)
	if rand_num > 99990 && civil.Revealed == false {
		return true
	}
	return false
}

func (civil *Civilization) Shape() (*ebiten.Image, error) {
	size := 2
	if civil.Revealed == true {
		size = 4
	}
	square, err := ebiten.NewImage(size, size, ebiten.FilterNearest)
	return square, err
}

func (civil *Civilization) GameWindowPosition() (float64, float64) {
	return civil.Position.TranslateToGameWindowPosition()
}

func (civil *Civilization) Draw(screen *ebiten.Image) {
	utils.DrawShapeAtPositionWithColor(civil, screen, civil.Color)
}
