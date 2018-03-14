package templates

import (
	"github.com/gobuffalo/packr"
)

func GetTemplates() packr.Box {
	box := packr.NewBox("../../../templates")
	return box
}
