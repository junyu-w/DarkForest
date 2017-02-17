package controllers

import (
	"dark_forest/models"
)

type CivilizationController struct {
	c *models.Civilization
}

func NewCivilizationController(civil *models.Civilization) *CivilizationController {
	return &CivilizationController{
		c: civil,
	}
}
