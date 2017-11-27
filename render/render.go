package render

import (
	"errors"
	"fmt"
	"github.com/d-led/risgo/app"
	"io/ioutil"
	"path"
	"strings"
)

// Ris : input for the code generator run
type Ris struct {
	Resource       string
	Template       string
	HeaderOverride string
	SourceOverride string
}

// Render : run the code generation
func (r *Ris) Render() {
	r.render()
}

func (r *Ris) render() {
	t := getTemplate(r.Template)

	fmt.Println("Template:", t.Name)
	fmt.Println("Resource:", r.Resource)

	resources := r.loadResources(r.Resource)

	applyOverrides(&resources, r)

	ctx := contextFrom(resources)

	renderToFile(t.Header, ctx, resources.Header)
	renderToFile(t.Source, ctx, resources.Source)
}

func applyOverrides(resources *resourceCollection, r *Ris) {
	if r.HeaderOverride != "" {
		resources.Header = r.HeaderOverride
	}

	if r.SourceOverride != "" {
		resources.Source = r.SourceOverride
	}
}

func renderToFile(template string, ctx map[string]interface{}, filename string) {
	res, err := renderTemplate(template, ctx)
	app.QuitOnError(err)
	fmt.Println("Writing", filename)
	writeAllText(res, filename)
}

func contextFrom(resources resourceCollection) map[string]interface{} {
	return map[string]interface{}{
		"source_include":   strings.Replace(resources.Header, "\\", "/", -1),
		"include_filename": path.Base(resources.Header),
		"namespace_name":   resources.Namespace,
		"class_name":       resources.Class,
		"resource":         prepareResources(resources),
	}
}

func prepareResources(resources resourceCollection) []map[string]interface{} {
	res := []map[string]interface{}{}
	// path.Dir doesn't seem to handle windows paths correctly
	baseDir := path.Clean(path.Dir(strings.Replace(resources.ResourceSource, "\\", "/", -1)))
	for _, r := range resources.Resources {

		if r.Compression != "" {
			app.QuitOnError(errors.New("Compression with " + r.Compression + " not supported yet!"))
		}

		memberName := memberNameFor(r)

		res = append(res,
			map[string]interface{}{
				"member_name": memberName,
				"name":        r.Name,
				"bytes":       renderBytesFor(resourceContent(r, baseDir)),
			},
		)
	}

	return res
}

func resourceContent(r resource, baseDir string) []byte {
	if r.SourceType == "string" {
		return []byte(r.Source)
	} else if r.SourceType == "file" {
		nextToTemplate := path.Clean(path.Join(baseDir, r.Source))
		return readAllBytes(nextToTemplate)
	} else {
		app.QuitOnError(errors.New("unknown source type: " + r.SourceType))
	}
	return []byte("")
}

func memberNameFor(r resource) string {
	if len(strings.TrimSpace(r.MemberName)) == 0 {
		return strings.TrimSpace(r.Name)
	}
	return strings.TrimSpace(r.MemberName)
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
