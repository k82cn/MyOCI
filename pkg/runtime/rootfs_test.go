package runtime

import (
	"os"
	"testing"

	"k8s.io/klog"
)

func Test_pivotRoot(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		klog.Error(err)
	}

	klog.Infof("cwd %s", cwd)

	if err = pivotRoot(cwd); err != nil {
		klog.Error(err)
	}
}
