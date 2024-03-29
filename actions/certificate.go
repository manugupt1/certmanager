package actions

import (
	"net/http"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/manugupt1/certmanager/models"
)

type CertificateActions struct{}

func (cr CertificateActions) ListCertificate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(http.StatusInternalServerError, r.JSON(&err))
	}
	certificate := &models.Certificates{}
	uid := c.Param("cust_id")
	if uid == "" {
		err := "User id cannot be empty"
		return c.Render(http.StatusBadRequest, r.JSON(&err))
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
		return c.Render(http.StatusBadRequest, r.JSON(&msg))
	}
	err := certificate.ListCertificate(tx, uid, active)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}
	return c.Render(http.StatusOK, r.JSON(certificate))
}

func (cr CertificateActions) UpdateStatus(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(http.StatusInternalServerError, r.JSON(&err))
	}

	id, err := strconv.Atoi(c.Param("cert_id"))
	if err != nil {
		err := "Certificate id is required and it needs to be an integer"
		return c.Render(http.StatusBadRequest, r.JSON(&err))
	}

	active := c.Param("active")
	var toActivate bool
	if active == "true" {
		toActivate = true
	} else if active != "false" {
		msg := "Needs active parameter as a true or false value only"
		return c.Render(http.StatusBadRequest, r.JSON(&msg))
	}

	cert := &models.Certificate{}

	err = cert.UpdateStatus(tx, id, toActivate)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(err.Error()))
	}
	return nil
}

func (cr CertificateActions) CreateCertificate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(http.StatusInternalServerError, r.JSON(&err))
	}

	id := c.Param("cust_id")
	if id == "" {
		err := "Customer id is required"
		return c.Render(http.StatusBadRequest, r.JSON(&err))
	}

	certificate := &models.Certificate{}
	err := certificate.CreateCertificate(tx, id)

	if err != nil {
		return err
	}
	return nil
}

func (cr CertificateActions) DownloadKey(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(http.StatusInternalServerError, r.JSON(&err))
	}

	cert_id := c.Param("cert_id")
	if cert_id == "" {
		err := "cert_id  is required"
		return c.Render(http.StatusBadRequest, r.JSON(&err))
	}

	cust_id := c.Param("cust_id")
	if cust_id == "" {
		err := "cust_id  is required"
		return c.Render(http.StatusBadRequest, r.JSON(&err))
	}

	cert := &models.Certificate{}
	keydata, err := cert.DownloadKey(tx, cert_id, cust_id)
	if err != nil {
		errMsg := err.Error()
		return c.Render(http.StatusInternalServerError, r.JSON(&errMsg))
	}
	return c.Render(http.StatusOK, r.JSON(keydata))
}

func (cr CertificateActions) DownloadBody(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	if tx == nil {
		err := "Database connection lost"
		return c.Render(http.StatusInternalServerError, r.JSON(&err))
	}

	cert_id := c.Param("cert_id")
	if cert_id == "" {
		err := "cert_id  is required"
		return c.Render(http.StatusBadRequest, r.JSON(&err))
	}

	cust_id := c.Param("cust_id")
	if cust_id == "" {
		err := "cust_id  is required"
		return c.Render(http.StatusBadRequest, r.JSON(&err))
	}

	cert := &models.Certificate{}
	bodydata, err := cert.DownloadBody(tx, cust_id, cert_id)
	if err != nil {
		errMsg := err.Error()
		return c.Render(http.StatusInternalServerError, r.JSON(&errMsg))
	}
	return c.Render(http.StatusOK, r.JSON(bodydata))

}
