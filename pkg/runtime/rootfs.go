package runtime

import (
	"os"
	"path/filepath"
	"syscall"

	"k8s.io/klog"
)

func pivotRoot(root string) error {
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		klog.Errorf("Failed to mount root <%s> with 'bind'", root)
		return err
	}

	pivotDir := filepath.Join(root, "pivot_root")
	if err := os.Mkdir(pivotDir, 0777); err != nil {
		return err
	}
	defer os.Remove(pivotDir)

	if err := syscall.PivotRoot(root, pivotDir); err != nil {
		klog.Infof("Failed to pivot root to <%s> from <%s>", root, pivotDir)
		return err
	}

	pivotDir = filepath.Join("/", "pivot_root")
	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
		klog.Infof("Failed to unmount old pivot dir <%s>", pivotDir)
		return err
	}

	return os.Remove(pivotDir)
}
