package dark_forest

import (
	"dark_forest/controllers"
	"dark_forest/models"
	"dark_forest/utils"
	"github.com/hajimehoshi/ebiten"
)

func BuildAndRunModel() {
	u := models.NewUniverse()
	u_c := controllers.NewUniverseController(u)
	if err := ebiten.Run(u_c.UpdateUniverse, utils.G_WIDTH, utils.G_HEIGHT, utils.G_SCALE, "Dark Forest"); err != nil {
		panic(err)
	}
}
