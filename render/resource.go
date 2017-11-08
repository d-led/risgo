package render

import (
	"github.com/d-led/risgo/app"
	"github.com/pelletier/go-toml"
)

type resource interface{}

func (r *Ris) loadResource(filename string) resource {
	tree, err := toml.LoadFile(filename)
	app.QuitOnError(err)
	return tree
}
