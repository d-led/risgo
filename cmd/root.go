package cmd

import (
	"fmt"
	render "github.com/d-led/risgo/render"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var template string
var sourceOverride string
var headerOverride string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "risgo <path_to>/<resources>.yml",
	Short: "a simple cross-platform resource compiler",
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.risgo.toml)")

	RootCmd.PersistentFlags().StringVar(&template, "template", "", "alternative template file to use")
	RootCmd.PersistentFlags().StringVar(&sourceOverride, "source", "", "alternative output source file (optional)")
	RootCmd.PersistentFlags().StringVar(&headerOverride, "header", "", "alternative output header file (optional)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".risgo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".risgo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
