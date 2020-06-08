package runtime

import (
	"os/exec"
	"syscall"

	"k8s.io/klog"
)

func NewParentProcess(tty bool, command string) *exec.Cmd {

	args := []string{"init", command}

	cmd := exec.Command("/proc/self/exec", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CloneFlags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd
}

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
