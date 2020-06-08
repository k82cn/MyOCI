package runtime

import (
	"os"
	"os/exec"
	"syscall"

	"k8s.io/klog"
)

func Run(tty bool, command string) {
	parent := newParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		klog.Error(err)
	}

	parent.Wait()
	os.Exit(-1)
}

func newParentProcess(tty bool, command string) *exec.Cmd {
	args := []string{"init", command}

	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd
}
