package cmd

import (
	"github.com/spf13/cobra"

	listCmd "github.com/ronaldlflau/tmp/cmd/list"
	updateCmd "github.com/ronaldlflau/tmp/cmd/update"
)

func NewCmdRoot() *cobra.Command {

	cmdRoot := &cobra.Command{
		Use:   "branch-rules <command> [flags]",
		Short: "List and update branch protection rules.",
		Long:  "List and update branch protection rules for repositories in an organization.",
	}

	cmdRoot.AddCommand(listCmd.NewCmdList())
	cmdRoot.AddCommand(updateCmd.NewCmdUpdate())
	cmdRoot.CompletionOptions.DisableDefaultCmd = true
	cmdRoot.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	return cmdRoot
}
