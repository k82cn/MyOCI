package app

import (
	"github.com/spf13/cobra"

	rt "github.com/k82cn/myoci/pkg/runtime"
)

var Run = &cobra.Command{
	Use:     "run",
	Short:   "Print the version information",
	Long:    "Print the version information",
	Example: `myoci run -it /bin/sh`,
	Run: func(cmd *cobra.Command, args []string) {
		rt.Run()
	},
}
