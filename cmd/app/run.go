package app

import (
	"github.com/spf13/cobra"

	rt "github.com/k82cn/myoci/pkg/runtime"
)

var runFlags rt.RunFlags

// RunCommand get the run command instance.
func RunCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:     "run",
		Short:   "Run an image as container",
		Long:    "Run an image as container",
		Example: `myoci run -it /bin/sh`,
		Run: func(cmd *cobra.Command, args []string) {
			rt.Run(&runFlags)
		},
	}

	setRunFlags(runCmd)

	return runCmd
}

func setRunFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&runFlags.Terminal, "terminal", "t", true, "true")
	cmd.Flags().BoolVarP(&runFlags.Interactive, "interactive", "i", true, "true")
	cmd.Flags().StringVarP(&runFlags.Command, "command", "c", "", "/bin/sh")
}
