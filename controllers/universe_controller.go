package controllers

import (
	"dark_forest/models"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
)

type UniverseController struct {
	u *models.Universe
}

func NewUniverseController(uni *models.Universe) *UniverseController {
	return &UniverseController{
		u: uni,
	}
}

func (uc *UniverseController) UpdateUniverse(screen *ebiten.Image) error {
	uc.u.Evovle(1000)
	if err := ebitenutil.DebugPrint(screen, "Our first game in Ebiten!"); err != nil {
		return err
	}
	return nil
}
