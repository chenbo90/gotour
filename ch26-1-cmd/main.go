package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myapp",
		Short: "MyApp is a sample CLI application",
		Long:  `MyApp is a sample CLI application built with Go and Cobra.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to MyApp! Use --help for more information.")
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of MyApp",
	Long:  `All software has versions. This is MyApp's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MyApp v1.0")
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the MyApp server",
	Long:  `Start the MyApp server at the specified port`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		fmt.Printf("Starting MyApp server on port %d\n", port)
	},
}

func init() {
	serveCmd.Flags().IntP("port", "p", 8080, "Port to run the server on")
}
