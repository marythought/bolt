package grifts

import (
	"github.com/markbates/grift/grift"
	"github.com/marythought/bolt/models"
)

var _ = grift.Add("db:seed", func(c *grift.Context) error {
	err := models.DB.RawQuery("delete from campaigns").Exec()
	if err != nil {
		return err
	}
	campaign := &models.Campaign{Name: "Purina Cat Chow", Type: "Social"}
	verrs, err := models.DB.ValidateAndCreate(campaign)
	if verrs.HasAny() {
		return verrs
	}
	return err
})
