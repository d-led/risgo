package render

import (
	"fmt"
	"github.com/d-led/risgo/app"
	"io/ioutil"
)

type Ris struct {
	resource string
	template string
}

func Render(resource string, template string) {
	r := Ris{resource, template}
	r.render()
}

func (r *Ris) render() {
	t := getTemplate(r.template)
	fmt.Println("Template:", t.name)
	fmt.Println("Resource:", r.resource)
	resources := r.loadResources(r.resource)
	ctx := contextFrom(resources)

	header, err := renderTemplate(t.header, ctx)
	app.QuitOnError(err)

	source, err := renderTemplate(t.source, ctx)
	app.QuitOnError(err)

	fmt.Println("Writing", resources.Header)
	writeAllText(header, resources.Header)

	fmt.Println("Writing", resources.Source)
	writeAllText(source, resources.Source)
}

func contextFrom(resources resource_collection) map[string]interface{} {
	return map[string]interface{}{
		"source_include": resources.Header,
		"namespace_name": resources.Namespace,
		"class_name":     resources.Class,
		"resource": []map[string]interface{}{
			{
				"member_name": "bla",
				"name":        "bla",
				"bytes":       renderButesFor("hello, world"),
			},
		},
	}
}

func writeAllText(text string, filename string) {
	err := ioutil.WriteFile(filename, []byte(text), 0644)
	app.QuitOnError(err)
}
