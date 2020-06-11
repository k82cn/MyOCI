package runtime

import (
	"os"
	"os/exec"
	"syscall"

	"k8s.io/klog"
)

// RunFlags is the flags of run command.
type RunFlags struct {
	Terminal    bool
	Interactive bool
	Command     string
}

// Run run target command in container
func Run(flags *RunFlags) {
	parent := newParentProcess(flags)
	if err := parent.Start(); err != nil {
		klog.Errorf("Failed to start parent process: %v", err)
	}

	parent.Wait()
	os.Exit(-1)
}

func newParentProcess(flags *RunFlags) *exec.Cmd {
	args := []string{"init", "-c", flags.Command}

	klog.Infof("command line args: %+v", args)

	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if flags.Terminal || flags.Interactive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd
}
