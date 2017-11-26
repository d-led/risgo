package render

import (
	"github.com/d-led/risgo/app"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type resource struct {
	Name        string
	Member_name string //optional
	Source_type string
	Source      string
	Compression string
}

type resource_collection struct {
	ResourceSource string // where
	Namespace      string
	Header         string
	Source         string
	Class          string
	Resources      []resource
}

func (r *Ris) loadResources(filename string) resource_collection {
	res := resource_collection{}

	data, err := ioutil.ReadFile(filename)
	app.QuitOnError(err)

	err = yaml.Unmarshal(data, &res)
	app.QuitOnError(err)

	res.ResourceSource = filename

	return res
}
