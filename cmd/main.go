package main

import (
	"os"
	"os/exec"
	"syscall"

	rt "github.com/k82cn/myoci/pkg/runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"k8s.io/klog"
)

func main() {
	tty := true
	cmd := "/bin/sh"

	Run(tty, cmd)
	rt.RunContainerInitProcess(cmd, nil)
}

func Run(tty bool, command string) {
	parent := rt.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		klog.Error(err)
	}

	parent.Wait()
	os.Exit(-1)
}
