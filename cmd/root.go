package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "ToDo is a cli based todo(s) manager.",
	Long: `ToDo is a cli based todo(s) manager application.
That you can use for tracking your daily tasks
without leaving your command prompt or bash.
This tool create for geeks 
"Who Love CLI and Want to Live There".`,
}
