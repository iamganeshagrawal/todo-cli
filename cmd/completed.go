package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completedListCmd = &cobra.Command{
	Use:     "completed",
	Aliases: []string{"c"},
	Short:   "List completed todo(s)",
	Long: `List all saved Todo(s) and
You can filter them based on pending/completed by passing subcommand.`,
	Example: `todo list completed`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list of completed todos")
	},
}
