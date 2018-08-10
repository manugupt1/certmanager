package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/manugupt1/certmanager/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
