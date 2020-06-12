package runtime

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/k82cn/myoci/pkg/subsystem"

	"k8s.io/klog"
)

// RunFlags is the flags of run command.
type RunFlags struct {
	Terminal    bool
	Interactive bool
	Command     string
	Args        []string

	subsystem.ResourceConfig
}

// Run run target command in container
func Run(flags *RunFlags) {
	parent := newParentProcess(flags)
	if err := parent.Start(); err != nil {
		klog.Errorf("Failed to start parent process: %v", err)
		return
	}

	mgrID := fmt.Sprintf("myoci-cgroup-%d", parent.Process.Pid)
	cgroupManager := subsystem.NewManager(mgrID)
	defer cgroupManager.Destroy()

	cgroupManager.Set(&flags.ResourceConfig)
	cgroupManager.Apply(parent.Process.Pid)

	parent.Wait()
	os.Exit(-1)
}

func newParentProcess(flags *RunFlags) *exec.Cmd {
	args := []string{"init", flags.Command}
	if len(flags.Args) != 0 {
		args = append(args, flags.Args...)
	}

	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWIPC,
	}

	if flags.Terminal || flags.Interactive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd
}
