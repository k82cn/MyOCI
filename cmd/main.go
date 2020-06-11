package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/k82cn/myoci/cmd/app"

	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()

	rootCmd := cobra.Command{
		Use: "myoci",
	}

	rootCmd.AddCommand(app.RunCommand())
	rootCmd.AddCommand(app.InitCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Failed to execute command: %v\n", err)
		os.Exit(2)
	}
}
