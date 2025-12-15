package cmd

import (
	"fmt"
	"mattiamancina/sshboy/config"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all the available servers",
	Run: func(cmd *cobra.Command, args []string) {
		prettyPrint(config.Get())
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func prettyPrint(cfg *config.Config) {
	var b strings.Builder

	b.WriteString("Hosts:\n")

	for _, s := range cfg.Servers {
		fmt.Fprintf(&b,
			"   %s\n"+
				"      Host: %s\n"+
				"      User: %s\n",
			s.Name,
			s.Host,
			s.User,
		)
	}

	fmt.Println(b.String())
}
