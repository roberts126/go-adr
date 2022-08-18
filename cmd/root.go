package cmd

import (
	"fmt"
	"os"

	"github.com/roberts126/go-adr/app"
	"github.com/spf13/cobra"
)

const fmtEnvFlag = "Can also be set using the %s environment variable"

var rootCmd = &cobra.Command{
	Use:   "adr",
	Short: "Tools for creating ADRs",
	Long:  `Tools for creating ADRs`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(
		initCommand(),
		versionCommand(),
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Displays the current version",
		Long:  `Displays the current version`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(app.Version)
		},
	}
}
