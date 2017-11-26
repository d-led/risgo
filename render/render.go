package render

import (
	"fmt"
)

type Ris struct {
	resource string
	template string
}

type template struct {
	name   string
	header string
	source string
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
	header, err := renderTemplate(t.header, map[string]interface{}{
		"header": resources.Header,
	})
	fmt.Println("Header", header, err)
}
