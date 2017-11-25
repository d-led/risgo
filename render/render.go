package render

import (
	"fmt"
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
	fmt.Println("Template:", templateName(r.template))
	fmt.Println("Resource:", r.resource)
	resources := r.loadResources(r.resource)
	fmt.Println(resources)
}

func templateName(t string) string {
	if t == "" {
		return "<default>"
	}

	return t
}
