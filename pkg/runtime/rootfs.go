package runtime

import (
	"os"
	"path/filepath"
	"syscall"

	"k8s.io/klog"
)

func pivotRoot(root string) error {
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return err
	}

	pivotDir := filepath.Join(root, ".pivot_root")
	oldPivotDir := filepath.Join("/", ".pivot_root")

	if err := os.Mkdir(pivotDir, 0777); err != nil {
		return err
	}

	defer func() {
		if _, err := os.Stat(pivotDir); err == nil {
			os.Remove(pivotDir)

		}

		if _, err := os.Stat(oldPivotDir); err == nil {
			os.Remove(oldPivotDir)
		}
	}()

	if err := syscall.PivotRoot(root, pivotDir); err != nil {
		return err
	}

	if err := syscall.Unmount(oldPivotDir, syscall.MNT_DETACH); err != nil {
		klog.Error(err)
	}

	return nil
}
