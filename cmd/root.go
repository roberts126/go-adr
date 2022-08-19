package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/roberts126/go-adr/app"
	"github.com/roberts126/go-adr/internal/env"
	"github.com/spf13/cobra"
)

const fmtEnvFlag = "Can also be set using the %s environment variable"

var (
	format  string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "adr",
	Short: "Tools for creating ADRs",
	Long:  `Tools for creating ADRs`,
	Run:   noop,
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Prints verbose output on commands that support it")
	rootCmd.PersistentFlags().StringVar(&format, "format", "", "The format to echo the data in. Can be one of text, json, or yaml")

	rootCmd.PreRun = func(_ *cobra.Command, _ []string) {
		if !strings.Contains(format+",", "text,json,yaml,") {
			format = "text"
		}
	}

	rootCmd.AddCommand(
		initCommand(),
		projectCommand(),
		versionCommand(),
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func noop(_ *cobra.Command, _ []string) {
}

func addPathFlag(v *string, cmd *cobra.Command) {
	usage := fmt.Sprintf(fmtEnvFlag, app.EnvConfigDir)
	cmd.Flags().StringVarP(v, "dir", "d", env.GetString(app.EnvConfigDir, ""), "The directory to create an adr in. Defaults to ${HOME}/.config/adr. "+usage)
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
