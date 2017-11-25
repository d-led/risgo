package render

import (
	"github.com/d-led/risgo/app"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type resource map[interface{}]interface{}

func (r *Ris) loadResource(filename string) resource {
	m := make(map[interface{}]interface{})
	data, err := ioutil.ReadFile(filename)
	app.QuitOnError(err)
	err = yaml.Unmarshal(data, &m)
	app.QuitOnError(err)
	return m
}
