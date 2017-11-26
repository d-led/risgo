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
	Resource       string
	Template       string
	HeaderOverride string
	SourceOverride string
}

func (r *Ris) Render() {
	r.render()
}

func (r *Ris) render() {
	t := getTemplate(r.Template)
	fmt.Println("Template:", t.Name)
	fmt.Println("Resource:", r.Resource)

	resources := r.loadResources(r.Resource)

	ctx := contextFrom(resources)

	header, err := renderTemplate(t.Header, ctx)
	app.QuitOnError(err)

	source, err := renderTemplate(t.Source, ctx)
	app.QuitOnError(err)

	header_filename := resources.Header
	if r.HeaderOverride != "" {
		header_filename = r.HeaderOverride
	}
	fmt.Println("Writing", header_filename)
	writeAllText(header, header_filename)

	source_filename := resources.Source
	if r.SourceOverride != "" {
		source_filename = r.SourceOverride
	}
	fmt.Println("Writing", source_filename)
	writeAllText(source, source_filename)
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
		return string(readAllBytes(next_to_template))
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

func readAllBytes(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	app.QuitOnError(err)
	return data
}

func writeAllText(text string, filename string) {
	err := ioutil.WriteFile(filename, []byte(text), 0644)
	app.QuitOnError(err)
}
