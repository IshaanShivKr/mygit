package cmd

import (
	"fmt"
	"os"

	"github.com/IshaanShivKr/mygit/internal/repo"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new mygit repository",
	Long:  `Creates a new .mygit directory with the required repository structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if err := repo.InitRepo(path); err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing repository: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Initialized empty mygit repository in .mygit/")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
