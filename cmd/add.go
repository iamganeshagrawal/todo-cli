package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add task string",
	Aliases: []string{"a", "create", "c"},
	Short:   "Add a TODO",
	Long: `Add a new todo task by passing task description.
Use double quotes for string if it contain special chars.`,
	Example: `todo add do review of PR#45879
todo add "I have to do something, that was more fishy."`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("adding new todo:\t\"%s\"", strings.Join(args, " "))
	},
}
