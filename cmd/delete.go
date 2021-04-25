package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete id number",
	Aliases: []string{"d"},
	Short:   "Delete a TODO",
	Long: `Delete or remove a todo by passing todo ID.
(For get todo ID use "todo list" command)`,
	Example: `todo delete 17`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("deleteing todo (ID): %s", args[0])
	},
}
