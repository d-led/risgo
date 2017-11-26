package render

import (
	// "fmt"
	"github.com/aymerick/raymond"
)

type template struct {
	name             string
	header           string
	source           string
	context_defaults map[string]interface{}
}

// ---
// class: Resource
// header: ris_lib/template.h
// namespace: ris
// resources:
//   -
//     name: header_preamble
//     source: "/* This file has been generated using ris, do not modify! */\n"
//     source_type: string
//   -
//     name: header_includes
//     source: "#include <string>\n"
//     source_type: string
//   -
//     name: class_name
//     source: Resource
//     source_type: string
//   -
//     name: source_preamble
//     source: "/* This file has been generated using ris, do not modify! */"
//     source_type: string
//   -
//     name: source_includes
//     source: "#include <unordered_map>"
//     source_type: string
//   -
//   -
//     name: source
//     source: |
//     source_type: string
//   -
//     name: header_single_declaration
//     source: "    static std::string {{resource_member_name}}();\n"
//     source_type: string
//   -
//     name: source_single_definition
//     source: |
//         std::string {{class_name}}::{{resource_member_name}}() {
//             static char const literal[] = {{{source_literal_bytes}}
//             };
//         {{source_return_literal}}
//         }
//     source_type: string
//   -
//     name: source_single_getter
//     source: "        { \"{{resource_name}}\", {{class_name}}::{{resource_member_name}} },\n"
//     source_type: string
//   -
//     name: header_single_resource_name
//     source: "        \"{{resource_name}}\",\n"
//     source_type: string
//   -
//     name: source_return_plain_literal
//     source: "    return std::string(literal, sizeof(literal)/sizeof(char));"
//     source_type: string
//   -
//     name: source_return_compressed_literal
//     source: "    return bundle::unpack(std::string(literal, sizeof(literal)/sizeof(char)));"
//     source_type: string
//   -
//     name: header_on_no_key
//     source: |
//         public:
//             static std::string OnNoKey(std::string const& key="") {
//                 return "";
//             }
//     source_type: string
// source: ris_lib/template.cpp

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
