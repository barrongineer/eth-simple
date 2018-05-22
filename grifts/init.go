package grifts

import (
	"github.com/barrongineer/simple/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
