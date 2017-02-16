package models

import (
	"math/rand"
	"dark_forest/utils"
	"fmt"
)

type Civilization struct {
	Id int
	Position *Coordinate
	Type string
	NumYears int
	Range float64 // unit: lightyear
	MatterOwned float64
	Level int
	ContainerUniverse *Universe
	Revealed bool
	MessageChannel chan *Coordinate
}

func NewCivilization(id int, pos *Coordinate, category string, universe *Universe) *Civilization {
	return &Civilization {
		Id: id,
		Position: pos,
		Type: category,
		NumYears: 0,
		Range: 10e-10,
		MatterOwned: 10e-5,
		Level: utils.LIGHTSPEED_x0001,
		ContainerUniverse: universe,
		Revealed: false,
		MessageChannel: make(chan *Coordinate),
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


func getLevel(matterPercentage float64) int {
	switch {
	case matterPercentage >= 10e-13:
		return utils.LIGHTSPEED_x0001
	case matterPercentage >= 10e-11:
		return utils.LIGHTSPEED_x001
	case matterPercentage >= 10e-7:
		return utils.LIGHTSPEED_x1
	case matterPercentage >= 10e-5:
		return utils.LIGHTSPEED_x2
	case matterPercentage >= 10e-2:
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
	nearby_civils := c.ContainerUniverse.GetNearbyCivilizations(c, 10)
	for _, civil := range nearby_civils {
		dist := GetDistance(c.Position, civil.Position)
		arrival_time := int(dist / speed) + c.NumYears
		go c.SendMessage(arrival_time, civil.MessageChannel)
	}
}

/**
 * Send message to another civilization via its message channel, this 
 * simulates the late arrival of message using a timer
 */
func (c *Civilization) SendMessage(arrival_time int, channel chan *Coordinate) {
	for {
		if c.NumYears >= arrival_time {
			channel <- c.Position
			break
		}
	}
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


// TODO: what to do after message arrived
func (civil *Civilization) ProcessMessage() {
	for {
		info := <-civil.MessageChannel
		fmt.Printf("Civilization %d got position (%d, %d)", civil.Id, info.x, info.y)
	}
}