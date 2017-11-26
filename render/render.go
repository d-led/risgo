package render

import (
	"fmt"
	"github.com/d-led/risgo/app"
	"io/ioutil"
	"strings"
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
		"resource":       prepareResources(resources.Resources),
	}
}

func prepareResources(resources []resource) []map[string]interface{} {
	res := []map[string]interface{}{}
	for _, r := range resources {
		member_name := memberName(r)

		res = append(res,
			map[string]interface{}{
				"member_name": member_name,
				"name":        r.Name,
				"bytes":       renderBytesFor(resourceContent(r)),
			},
		)
	}

	return res
}

func resourceContent(r resource) string {
	if r.Source_type == "string" {
		return r.Source
	} else {
		return "...todo"
	}
}

func memberName(r resource) string {
	if len(strings.TrimSpace(r.Member_name)) == 0 {
		return strings.TrimSpace(r.Name)
	} else {
		return strings.TrimSpace(r.Member_name)
	}
}

func writeAllText(text string, filename string) {
	err := ioutil.WriteFile(filename, []byte(text), 0644)
	app.QuitOnError(err)
}
