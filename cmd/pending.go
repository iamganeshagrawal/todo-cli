package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pendingListCmd = &cobra.Command{
	Use:     "pending",
	Aliases: []string{"p"},
	Short:   "List pending todo(s)",
	Long: `List all saved Todo(s) and
You can filter them based on pending/completed by passing subcommand.`,
	Example: `todo list pending`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list of pending todos")
	},
}
