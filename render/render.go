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
	t := getTemplate(r.template)
	fmt.Println("Template:", t.name)
	fmt.Println("Resource:", r.resource)
	resources := r.loadResources(r.resource)
	header, err := renderTemplate(t.header, map[string]interface{}{
		"header":         resources.Header,
		"namespace_name": "my_namespace",
		"class_name":     "Resource",
		"resource": []map[string]interface{}{
			{
				"member_name": "bla",
				"name":        "bla",
			},
		},
	})
	fmt.Println("Header", header, err)
}
