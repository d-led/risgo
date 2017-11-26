package render

import (
	"bytes"
	"strconv"

	"github.com/aymerick/raymond"
)

type template struct {
	name             string
	header           string
	source           string
	context_defaults map[string]interface{}
}

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
		source: defaultSourceTemplate,
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

func renderButesFor(what string) func() string {
	return func() string {
		var buffer bytes.Buffer
		const maxBytesOnLine = 3
		for i, b := range []byte(what) {
			buffer.WriteString(strconv.Itoa(int(b)))
			buffer.WriteString(", ")
			if (i+1)%maxBytesOnLine == 0 {
				buffer.WriteRune('\n')
			}
		}
		return buffer.String()
	}
}

const defaultHeaderTemplate = `#pragma once
/* This file has been generated using ris, do not modify! */
#include <string>

namespace {{namespace_name}} {
class {{class_name}} /*final*/ {
public:
  {{#each resource}}
    static std::string {{member_name}};
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
{{header_on_no_key}}
};
}
`

const defaultSourceTemplate = `/* This file has been generated using ris, do not modify! */
#include <unordered_map>
{{{source_include}}}

namespace {{namespace_name}} {
{{#each resource}}
    std::string {{class_name}}::{{member_name}}() {
        static char const literal[] = {
{{bytes}}
        ;
        return std::string(literal, sizeof(literal)/sizeof(char));
    }
{{/each}}

std::string {{class_name}}::Get(std::string const& key) {
    static std::unordered_map<std::string,ResourceGetter> getters = {
{{source_getters}}    };
    auto getter = getters.find(key);
    if (getter == getters.end())
         return OnNoKey(key);
    return getter->second();
}
}
`
