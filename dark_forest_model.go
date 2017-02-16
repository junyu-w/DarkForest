package dark_forest

import (
	"dark_forest/models"
)

func BuildAndRunModel() {
	u := models.NewUniverse()
	for {
		u.Evovle(1000)
	}
}
