# risgo


[![Build Status](https://travis-ci.org/d-led/risgo.svg?branch=master)](https://travis-ci.org/d-led/risgo)


a simple cross-platform C++ resource compiler
> embed data / strings as constants in source

C++ implementation: [ris](https://github.com/d-led/ris)

## usage

basic usage:

```
risgo <path_to>/<resources>.yml
``

Given the following input

```yaml
header: "examples/resource.h"
source: "examples/resource.cpp"
namespace: "test"
class: "res"
resources:
  -
    name: "string_test"
    source: "plain text"
    source_type: "string"
  -
    name: "binary_file_test"
    source: "test.bin"
    source_type: "file"
```

the files `examples/resource.h` and `.cpp` will be generated, where `test::res::string_test()` will return `plain text` and `test::res::binary_file_test()` will return the binary contents of `test.bin` (read relative to the input YAML file).

More help on template and output overrides and a template example: `risgo --help`

## build

- checkout in `$GOPATH/src/<user>`
- Linux/Mac: `make test`
- Windows: `make.bat`

## credits

- [spf13/cobra](https://github.com/spf13/cobra) CLI implementation
- [go-yaml/yaml](https://github.com/go-yaml/yaml) YAML parsing/rendering
- [aymerick/raymond](https://github.com/aymerick/raymond) Handlebars (templates) for output customization
- [onqtam/doctest](https://github.com/onqtam/doctest) ([examples/doctest.h](examples/doctest.h)) C++ tests
