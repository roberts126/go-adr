package cmd

import (
	"github.com/roberts126/go-adr/internal/terminal"
	"github.com/spf13/cobra"
)

type projectHandler struct {
	name    string
	path    string
	verbose bool
}

func projectCommand() *cobra.Command {
	h := &projectHandler{}

	c := &cobra.Command{
		Use:     "project",
		Aliases: []string{"projects"},
		Short:   "Tools for creating or modifying projects.",
		Long:    "Tools for creating or modifying projects.",
		Run:     h.handle,
	}

	c.AddCommand(
		h.addCommand(),
		h.configCommand(),
		h.listCommand(),
	)

	return c
}

func (h *projectHandler) handle(_ *cobra.Command, _ []string) {}

func (h *projectHandler) addCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "add",
		Short: "Adds a new project.",
		Long:  "Adds a new project.",
		Run:   h.addProject,
	}

	addPathFlag(&h.path, c)
	c.Flags().StringVarP(&h.name, "name", "n", "", "The name of the project")

	return c
}

func (h *projectHandler) addProject(_ *cobra.Command, args []string) {
	if h.name == "" && len(args) > 0 {
		h.name = args[0]
	}

	if h.name == "" {
		terminal.Panic("Could not determine project name")
	}

	terminal.Info(h.name)
}

func (h *projectHandler) configCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Configures an existing project.",
		Long:  "Configures an existing project.",
		Run:   noop,
	}
}

func (h *projectHandler) listCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: "Lists all project.",
		Long:  "Lists all project.",
		Run:   h.listProjects,
	}

	addPathFlag(&h.path, c)

	return c
}

func (h *projectHandler) listProjects(_ *cobra.Command, _ []string) {
}
