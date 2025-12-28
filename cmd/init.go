/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dirs := []string{".mygit", ".mygit/objects", ".mygit/refs"}
		for _, dirPath := range dirs {
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", dirPath, err)
				os.Exit(1)
			}
		}

		headFile := ".mygit/HEAD"
		headContents := []byte("ref: refs/heads/main\n")
		if err := os.WriteFile(headFile, headContents, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write file %s: %v\n", headFile, err)
			os.Exit(1)
		}

		fmt.Println("Initialized empty Git repository in .mygit/")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
