package cmd

import (
	"fmt"
	"mattiamancina/sshboy/config"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{

		Use:   "sshboy",
		Short: "sshboy is a simple tool to connect to a server",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

	home, _ := os.UserHomeDir()

	config.Init(fmt.Sprintf(
		"%s/.sshboy/inventory.yaml",
		home,
	))

}
