package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// 创建一个命令
var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "An example CLI application",
	Long:  "A detailed description of the CLI application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

var Person struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

var versionCmd = &cobra.Command{
	Use: "version",
	//Short: "Print the version number of CLI application",
	//Long: "All software has version, This is CLI application's ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
