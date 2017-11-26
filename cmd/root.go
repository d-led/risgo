package cmd

import (
	"fmt"
	render "github.com/d-led/risgo/render"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string
var template string
var sourceOverride string
var headerOverride string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "risgo <path_to>/<resources>.yml",
	Short: "A simple cross-platform resource compiler",
	Long:  longHelp(),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ris := render.Ris{
			Resource:       args[0],
			Template:       template,
			SourceOverride: sourceOverride,
			HeaderOverride: headerOverride,
		}
		ris.Render()
	},
}

func longHelp() string {
	return `Example input yaml:

--------------------------------
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
--------------------------------

which will result in examples/resource.h and .cpp being generated

default template (rendered via github.com/aymerick/raymond):
------------------------------------------------------------
` + render.DefaultTemplateYaml() + `------------------------------------------------------------`
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	flags := RootCmd.PersistentFlags()

	flags.StringVar(&template, "template", "", "alternative template file to use")
	flags.StringVar(&sourceOverride, "source", "", "alternative output source file (optional)")
	flags.StringVar(&headerOverride, "header", "", "alternative output header file (optional)")
}
