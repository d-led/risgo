package render

import (
	"fmt"
	"github.com/aymerick/raymond"
)

const defaultHeaderTemplate = `#pragma once
{{header_preamble}}
{{header_includes}}
namespace {{namespace_name}} {
class {{class_name}} /*final*/ {
public:
{{header_declarations}}
public:
    typedef std::string(*ResourceGetter)();
public: // key/value api
template <typename TInserter>
static void GetKeys(TInserter inserter) {
    static const char* keys[] = {
{{header_resource_names}}    };
    for (auto key : keys) {
        inserter(key);
    }
}
public: // key/value api
    static std::string Get(std::string const& key);
{{header_on_no_key}}
};
}
`

func getTemplate(t string) template {
	if t == "" {
		return defaultTemplate()
	}

	return loadTemplate(t)
}

func defaultTemplate() template {
	return template{
		name:   "<default>",
		header: defaultHeaderTemplate,
		source: defaultHeaderTemplate,
	}
}

func loadTemplate(filename string) template {
	return template{
		name: filename,
	}
}

func renderTemplate(tpl string, ctx map[string]interface{}) (string, error) {
	return raymond.Render(tpl, ctx)
}

func test() {
	tpl := `<div class="entry">
  <h1>{{title}}</h1>
  <div class="body">
    {{body}}
  </div>
</div>
`

	ctx := map[string]string{
		"title": "My New Post",
		"body":  "This is my first post!",
	}

	result, err := raymond.Render(tpl, ctx)
	if err != nil {
		panic("Please report a bug :)")
	}

	fmt.Print(result)
}
