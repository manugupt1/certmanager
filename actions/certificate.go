package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/manugupt1/certmanager/models"
)

type CertificateActions struct{}

func (cr CertificateActions) ListCertificate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(500, r.JSON(&err))
	}
	certificate := &models.Certificates{}
	uid := c.Param("uid")
	if uid == "" {
		err := "User id cannot be empty"
		return c.Render(400, r.JSON(&err))
	}

	var active *bool
	if c.Param("active") == "true" {
		temp := true
		active = &temp
	} else if c.Param("active") == "false" {
		temp := false
		active = &temp
	} else if len(c.Param("active")) > 0 {
		msg := "value of active param is not recognized"
		return c.Render(400, r.JSON(&msg))
	}
	err := certificate.ListCertificate(tx, uid, active)
	if err != nil {
		return c.Render(500, r.JSON(err))
	}
	return c.Render(200, r.JSON(certificate))
}
