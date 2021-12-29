/*
Copyright © 2021 luckyhappycat Beijing China <luckyhappycat@126.com>

*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)


// go build -o vscode main.go

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vscode",
	Short: "vscode is a cli tool to generate .vscode config",
	Long: `vscode is a cli tool to generate .vscode config,
it supports gcc, g++, cmake`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vscode-configer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	log.Print("root.init()")

}
