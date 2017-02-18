package utils

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

type Drawable interface {
	Shape() (*ebiten.Image, error)
	GameWindowPosition() (float64, float64)
}

func DrawShapeAtPositionWithColor(d Drawable, screen *ebiten.Image, color color.NRGBA) error {
	shape, err := d.Shape()
	if err != nil {
		ebitenutil.DebugPrint(screen, "Something went wrong while drawing")
		return err
	}
	shape.Fill(color)
	x, y := d.GameWindowPosition()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
	if err = screen.DrawImage(shape, opts); err != nil {
		return err
	}
	return nil
}
