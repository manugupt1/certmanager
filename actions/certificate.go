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
	uid := c.Param("cust_id")
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
		msg := "value of active param is should be true, false or empty (to display every certificate)"
		return c.Render(400, r.JSON(&msg))
	}
	err := certificate.ListCertificate(tx, uid, active)
	if err != nil {
		return c.Render(500, r.JSON(err))
	}
	return c.Render(200, r.JSON(certificate))
}

func (cr CertificateActions) UpdateStatus(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(500, r.JSON(&err))
	}

	id := c.Param("cert_id")
	if id == "" {
		err := "Certificate id is required"
		return c.Render(400, r.JSON(&err))
	}

	active := c.Param("active")
	var toActivate bool
	if active == "true" {
		toActivate = true
	} else if active != "false" {
		msg := "Needs active parameter as a true or false value only"
		return c.Render(400, r.JSON(&msg))
	}

	cert := &models.Certificate{}

	return cert.UpdateStatus(tx, id, toActivate)
}

func (cr CertificateActions) CreateCertificate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(500, r.JSON(&err))
	}

	id := c.Param("cust_id")
	if id == "" {
		err := "Customer id is required"
		return c.Render(400, r.JSON(&err))
	}

	cert := &models.Certificate{}
	return cert.Create(tx, id)
}
