package models

import (
	"errors"
)

type DarkForestAttack struct {
	attacker   *Civilization
	victim_pos *Coordinate
}

var ALREADY_DESTROTYED error = errors.New("Civilization already destroyed")

func (dfa *DarkForestAttack) Execute() error {
	target := dfa.attacker.ContainerUniverse.GetCivilAtPosition(dfa.victim_pos)
	if target.Color == DEATH_COLOR {
		return ALREADY_DESTROTYED
	}
	target.Color = DEATH_COLOR
	if dfa.attacker.Color != DEATH_COLOR {
		dfa.attacker.Color = HIDDEN_COLOR
	}
	return nil
}

func NewDFAttack(a *Civilization, p *Coordinate) *DarkForestAttack {
	return &DarkForestAttack{
		attacker:   a,
		victim_pos: p,
	}
}
