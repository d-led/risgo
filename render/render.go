package render

import (
	"bytes"
	"fmt"
	"strconv"
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
	ctx := map[string]interface{}{
		"header":         resources.Header,
		"source_include": "#include \"resource.h\"",
		"namespace_name": "my_namespace",
		"class_name":     "Resource",
		"resource": []map[string]interface{}{
			{
				"member_name": "bla",
				"name":        "bla",
				"bytes": func() string {
					var buffer bytes.Buffer
					for _, b := range []byte("hello, world") {
						buffer.WriteString(strconv.Itoa(int(b)))
						buffer.WriteString(", ")
					}
					return buffer.String()
				},
			},
		},
	}
	header, err := renderTemplate(t.header, ctx)
	source, err := renderTemplate(t.source, ctx)
	fmt.Println("Header", header, err)
	fmt.Println("Source", source, err)
}
