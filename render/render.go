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
	resource := r.loadResource(r.resource)
	fmt.Println(resource)
}

func templateName(t string) string {
	if t == "" {
		return "<default>"
	}

	return t
}
