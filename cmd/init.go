package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new config",
	Run: func(cmd *cobra.Command, args []string) {

		home, err := os.UserHomeDir()

		err = os.MkdirAll(fmt.Sprintf("%s/.sshboy", home), os.ModePerm)

		if err != nil {
			fmt.Println("An error occured while writing ~/.sshboy/inventory.yaml", err)
			return
		}
		_, err = os.Create(fmt.Sprintf("%s/.sshboy/inventory.yaml", home))
		if err != nil {
			fmt.Println("An error occured while writing ~/.sshboy/inventory.yaml", err)
			os.Exit(1)
		}
	},
}
