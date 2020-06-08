package runtime

import (
	"os"
	"syscall"

	"k8s.io/klog"
)

func RunContainerInitProcess(command string, args []string) error {
	klog.Infof("command %s", command)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		klog.Error(err)
		return err
	}
	return nil
}
