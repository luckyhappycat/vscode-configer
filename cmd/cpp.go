/*
Copyright © 2021 luckyhappycat HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// cppCmd represents the cpp command
var cppCmd = &cobra.Command{
	Use:   "cpp",
	Short: "Add .vscode configuration file for C++ language",
	Long: `Add .vscode configuration file for C++ language.
It provides configuration files, such as:
.vscode/xxx.json, .vscode/yyy.json`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cpp called")
	},
}

func init() {
	log.Print("cpp.init()")
	rootCmd.AddCommand(cppCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cppCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cppCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
