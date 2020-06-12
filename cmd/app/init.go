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
		Example: `myoci init /bin/sh`,
		Run: func(cmd *cobra.Command, args []string) {
			initFlags.Command = args[0]
			if len(args) > 1 {
				initFlags.Args = args[1:]
			}

			rt.RunContainerInitProcess(&initFlags)
		},
	}

	setInitFlags(initCmd)

	return initCmd
}

func setInitFlags(cmd *cobra.Command) {
}
