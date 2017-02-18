package controllers

import (
	"dark_forest/models"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
)

type UniverseController struct {
	*models.Universe
}

func NewUniverseController(uni *models.Universe) *UniverseController {
	return &UniverseController{
		Universe: uni,
	}
}

func (u *UniverseController) UpdateUniverse(screen *ebiten.Image) error {
	u.UpdateAndDrawUniverse(10000, screen)
	if err := ebitenutil.DebugPrint(screen, fmt.Sprintf("%s\nUniverse Age: %d", u.UniversalMessage, u.NumYears)); err != nil {
		return err
	}
	return nil
}
