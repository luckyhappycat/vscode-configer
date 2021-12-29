/*
Copyright © 2021 luckyhappycat Beijing China <luckyhappycat@126.com>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/luckyhappycat/vscode-configer/pkg/cpp"

	"github.com/spf13/cobra"
)

var cmdType string

var cppCmd = &cobra.Command{
	Use:   "cpp",
	Short: "Add .vscode configuration file for C++ language",
	Long: `Add .vscode configuration file for C++ language.
It provides configuration files, such as:
.vscode/xxx.json,
.vscode/yyy.json`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cpp called")
		fmt.Println(cmdType)
		cpp.CreateVscodeConfig()
	},
}

func init() {
	log.Print("cpp.init()")
	cppCmd.PersistentFlags().StringVarP(&cmdType, "type", "t", "g++", "type of compiler")
	rootCmd.AddCommand(cppCmd)
}
