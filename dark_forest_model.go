package dark_forest

import (
	"github.com/DrakeW/DarkForest/controllers"
	"github.com/DrakeW/DarkForest/models"
	"github.com/DrakeW/DarkForest/utils"
	"github.com/hajimehoshi/ebiten"
)

func BuildAndRunModel() {
	u := models.NewUniverse()
	u_c := controllers.NewUniverseController(u)
	if err := ebiten.Run(u_c.UpdateUniverse, utils.G_WIDTH, utils.G_HEIGHT, utils.G_SCALE, "Dark Forest"); err != nil {
		panic(err)
	}
}
