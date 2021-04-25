package cmd

import (
	"fmt"

	"github.com/iamganeshagrawal/todo-cli/pkg/update"
	"github.com/iamganeshagrawal/todo-cli/pkg/version"
	"github.com/spf13/cobra"
)

var checkUpdateFlag bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Long:  `Print the build information of current installed build.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if checkUpdateFlag {
			update.CheckUpdate(version.Version)
		} else {
			fmt.Println("Version\t\t:", version.Version)
			fmt.Println("Build Time\t:", version.BuildTime)
		}
	},
}

func init() {
	versionCmd.Flags().BoolVar(&checkUpdateFlag, "update", false, "Check available update")
}
