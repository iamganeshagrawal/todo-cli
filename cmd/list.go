package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all todo(s)",
	Long: `List all saved Todo(s) and
You can filter them based on pending/completed by passing subcommand.`,
	Example: `todo list`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list of available todos")
	},
}

func init() {
	listCmd.AddCommand(pendingListCmd, completedListCmd)
}
