package subsystem

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubsystem struct {
}

// Name returns the name of subsystem.
func (m *MemorySubsystem) Name() string {
	return "memory"
}

func (m *MemorySubsystem) Set(cpath string, res *ResourceConfig) error {
	subsyspath, err := GetCgroupPath(m.Name(), cpath, true)
	if err != nil {
		return err
	}

	// TODO: check numberic
	if res.MemoryLimit != "" {
		mpath := path.Join(subsyspath, "memory.limit_in_bytes")
		return ioutil.WriteFile(mpath, []byte(res.MemoryLimit), 0644)
	}

	return nil
}

func (m *MemorySubsystem) Apply(cpath string, pid int) error {
	subsyspath, err := GetCgroupPath(m.Name(), cpath, true)
	if err != nil {
		return err
	}

	taskpath := path.Join(subsyspath, "tasks")
	return ioutil.WriteFile(taskpath, []byte(strconv.Itoa(pid)), 0644)
}

func (m *MemorySubsystem) Remove(cpath string) error {
	subsyspath, err := GetCgroupPath(m.Name(), cpath, true)
	if err != nil {
		return err
	}

	return os.Remove(subsyspath)
}
