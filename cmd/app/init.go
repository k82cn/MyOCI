package app

import (
	"github.com/spf13/cobra"

	rt "github.com/k82cn/myoci/pkg/runtime"
)

var initFlags rt.InitFlags

// InitCommand builds an "init" command instance.
func InitCommand() *cobra.Command {
	initCmd := &cobra.Command{
		Use:     "init",
		Short:   "Init current process as container",
		Long:    "Init current process as container",
		Example: `myoci init -c /bin/sh`,
		Run: func(cmd *cobra.Command, args []string) {
			// initFlags.Args = args
			rt.RunContainerInitProcess(&initFlags)
		},
	}

	setInitFlags(initCmd)

	return initCmd
}

func setInitFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&initFlags.Command, "command", "c", "", "command")
}
