package render

import (
	"bytes"
	"github.com/aymerick/raymond"
	"github.com/d-led/risgo/app"

	"gopkg.in/yaml.v2"
	"strconv"
)

type template struct {
	Name   string
	Header string
	Source string
}

func getTemplate(t string) template {
	if t == "" {
		return defaultTemplate()
	}

	return loadTemplate(t)
}

func defaultTemplate() template {
	return template{
		Name:   "<default>",
		Header: defaultHeaderTemplate,
		Source: defaultSourceTemplate,
	}
}

func DefaultTemplateYaml() string {
	t := defaultTemplate()
	res, err := yaml.Marshal(&t)
	app.QuitOnError(err)
	return string(res)
}

func loadTemplate(filename string) template {
	t := template{}
	err := yaml.Unmarshal(readAllBytes(filename), &t)
	app.QuitOnError(err)

	if t.Name == "" {
		t.Name = filename
	}

	return t
}

func renderTemplate(tpl string, ctx map[string]interface{}) (string, error) {
	return raymond.Render(tpl, ctx)
}

func renderBytesFor(what []byte) func() string {
	return func() string {
		var buffer bytes.Buffer
		const maxBytesOnLine = 80
		for i, b := range what {
			if i != 0 {
				buffer.WriteString(", ")
			}

			buffer.WriteString(strconv.Itoa(int(b)))

			if (i+1)%maxBytesOnLine == 0 {
				buffer.WriteRune('\n')
			}
		}
		return buffer.String()
	}
}

const defaultHeaderTemplate = `#pragma once
/* This file has been generated using risgo, do not modify! */
#include <string>

namespace {{namespace_name}} {

class {{class_name}} /*final*/ {
public:
  {{#each resource}}
    static std::string {{member_name}}();
  {{/each}}
public:
    typedef std::string(*ResourceGetter)();
public: // key/value api
template <typename TInserter>
static void GetKeys(TInserter inserter) {
    static const char* keys[] = {
        {{#each resource}}
        "{{name}}",
        {{/each}}
    };
    for (auto key : keys) {
        inserter(key);
    }
}
public: // key/value api
    static std::string Get(std::string const& key);

public:
    static std::string OnNoKey(std::string const& key="") {
        return "";
    }
};

}
`

const defaultSourceTemplate = `/* This file has been generated using risgo, do not modify! */
#include <unordered_map>
#include "{{source_include}}"

namespace {{namespace_name}} {

{{#each resource}}
    std::string {{class_name}}::{{member_name}}() {
        static char const literal[] = {
{{bytes}} }
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
{{/each}}

std::string {{class_name}}::Get(std::string const& key) {
    static std::unordered_map<std::string,ResourceGetter> getters = {
{{#each resource}}
    { "{{name}}", {{class_name}}::{{member_name}} },
{{/each}}
    };
    auto getter = getters.find(key);
    if (getter == getters.end())
         return OnNoKey(key);
    return getter->second();
}

}
`
