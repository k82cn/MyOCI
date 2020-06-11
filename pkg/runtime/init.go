package runtime

import (
	"os"
	"syscall"

	"k8s.io/klog"
)

// InitFlags is the flags of init command instance.
type InitFlags struct {
	Command string
	Args    []string
}

// RunContainerInitProcess start init process of the container
func RunContainerInitProcess(flags *InitFlags) error {
	klog.Infof("command %s", flags.Command)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	argv := []string{flags.Command}
	if len(flags.Args) != 0 {
		argv = append(argv, flags.Args...)
	}

	if err := syscall.Exec(flags.Command, argv, os.Environ()); err != nil {
		klog.Error(err)
		return err
	}

	return nil
}
