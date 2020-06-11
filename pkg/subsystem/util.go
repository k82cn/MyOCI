package subsystem

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

// FindCgroupMountPoint finds the Cgroup mount point of subsystem.
func FindCgroupMountPoint(subsystem string) (string, error) {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4], nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("unknown error")
}

// GetCgroupPath get the Cgroup path of the subsystem.
func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot, err := FindCgroupMountPoint(subsystem)
	if err != nil {
		return "", err
	}

	cgpath := path.Join(cgroupRoot, cgroupPath)
	if _, err := os.Stat(cgpath); err != nil {
		if os.IsNotExist(err) && autoCreate {
			if err := os.Mkdir(cgpath, 0755); err != nil {
				return "", err
			}

			return "", nil
		}

		return "", err
	}

	return cgpath, nil
}
