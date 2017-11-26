package render

import (
	"errors"
	"fmt"
	"github.com/d-led/risgo/app"
	"io/ioutil"
	"path"
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
		"resource":       prepareResources(resources),
	}
}

func prepareResources(resources resource_collection) []map[string]interface{} {
	res := []map[string]interface{}{}
	base_dir := path.Dir(resources.ResourceSource)
	for _, r := range resources.Resources {

		if r.Compression != "" {
			app.QuitOnError(errors.New("Compression with " + r.Compression + " not supported yet!"))
		}

		member_name := memberName(r)

		res = append(res,
			map[string]interface{}{
				"member_name": member_name,
				"name":        r.Name,
				"bytes":       renderBytesFor(resourceContent(r, base_dir)),
			},
		)
	}

	return res
}

func resourceContent(r resource, base_dir string) string {
	if r.Source_type == "string" {
		return r.Source
	} else if r.Source_type == "file" {
		next_to_template := path.Join(base_dir, r.Source)
		return readAllText(next_to_template)
	} else {
		app.QuitOnError(errors.New("unknown source type: " + r.Source_type))
	}
	return ""
}

func memberName(r resource) string {
	if len(strings.TrimSpace(r.Member_name)) == 0 {
		return strings.TrimSpace(r.Name)
	} else {
		return strings.TrimSpace(r.Member_name)
	}
}

func readAllText(filename string) string {
	data, err := ioutil.ReadFile(filename)
	app.QuitOnError(err)
	return string(data)
}

func writeAllText(text string, filename string) {
	err := ioutil.WriteFile(filename, []byte(text), 0644)
	app.QuitOnError(err)
}
