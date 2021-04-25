package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var markAsPending bool

var markCmd = &cobra.Command{
	Use:     "mark id <number>",
	Aliases: []string{"m", "done"},
	Short:   "Mark todo as completed / pending",
	Long: `Mark a new todo task by passing task id.
Use uncheck or u flag to mark it as pending.`,
	Example: `todo mark 17
todo mark 17 --uncheck`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if markAsPending {
			fmt.Printf("marked todo (ID): %s as pending.", args[0])
		} else {
			fmt.Printf("marked todo (ID): %s as completed.", args[0])
		}
	},
}

func init() {
	markCmd.Flags().BoolVarP(&markAsPending, "uncheck", "u", false, "Mark as pending.")
}
