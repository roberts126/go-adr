package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/roberts126/go-adr/adr"
	"github.com/roberts126/go-adr/app"
	"github.com/roberts126/go-adr/internal/env"
	"github.com/roberts126/go-adr/internal/terminal"
	"github.com/spf13/cobra"
)

type initHandler struct {
	path         string
	templatePath string
}

func initCommand() *cobra.Command {
	h := initHandler{}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initializes the settings for ADRs",
		Long:  `Initializes the settings for ADRs`,
		Run:   h.handle,
	}

	usage := fmt.Sprintf(fmtEnvFlag, app.EnvConfigDir)
	cmd.Flags().StringVarP(&h.path, "dir", "d", env.GetString(app.EnvConfigDir, ""), "The directory to create an adr in. Defaults to ${HOME}/.config/adr. "+usage)

	return cmd
}

func (h *initHandler) handle(_ *cobra.Command, _ []string) {
	if h.path == "" {
		cwd, _ := os.UserHomeDir()
		h.path = filepath.Join(cwd, ".config", "adr")
	}

	h.templatePath = filepath.Join(h.path, "templates")

	_, err := os.Stat(h.path)
	if !os.IsNotExist(err) {
		terminal.Warnf("Directory %s already exists", h.path)
		os.Exit(0)
	}

	h.createDirectories(h.path, h.templatePath)
	h.saveInitialConfig()

	terminal.Success("ADR directory initialized")
}

func (h *initHandler) createDirectories(dirs ...string) {
	for i := 0; i < len(dirs); i++ {
		if err := os.MkdirAll(dirs[i], app.DefaultDirPerms); err != nil {
			terminal.Panicf("Error while creating directory %s; error: %v", dirs[i], err)
		}
	}
}

func (h *initHandler) saveInitialConfig() {
	path := filepath.Join(h.path, "config.yaml")

	if err := ioutil.WriteFile(path, []byte(adr.ExampleConfiguration), app.DefaultFilePerms); err != nil {
		terminal.Infof("ADR directory created at %s", h.path)
		terminal.Panicf("Unable to save initial configuration; error: %v", err)
	}

	path = filepath.Join(h.templatePath, "default.tpl")
	if err := ioutil.WriteFile(path, []byte(adr.DefaultTemplate), app.DefaultFilePerms); err != nil {
		terminal.Infof("ADR directory created at %s", h.path)
		terminal.Panicf("Unable to save initial configuration; error: %v", err)
	}
}
