package render

import (
	"github.com/d-led/risgo/app"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type resource struct {
	Name        string
	MemberName  string `yaml:"member_name"` //optional
	SourceType  string `yaml:"source_type"`
	Source      string
	Compression string
}

type resourceCollection struct {
	ResourceSource string // where
	Namespace      string
	Header         string
	Source         string
	Class          string
	Resources      []resource
}

func (r *Ris) loadResources(filename string) resourceCollection {
	res := resourceCollection{}

	data, err := ioutil.ReadFile(filename)
	app.QuitOnError(err)

	err = yaml.Unmarshal(data, &res)
	app.QuitOnError(err)

	res.ResourceSource = filename

	return res
}
