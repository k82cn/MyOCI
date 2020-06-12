package subsystem

import "k8s.io/klog"

type Manager struct {
	Path     string
	Resource *ResourceConfig
}

func NewManager(path string) *Manager {
	return &Manager{
		Path: path,
	}
}

func (m *Manager) Apply(pid int) error {
	for _, subSysIns := range subsystemIns {
		if err := subSysIns.Apply(m.Path, pid); err != nil {
			klog.Errorf("Failed to apply cgrouop to %s for %d: %v", m.Path, pid, err)
		}
	}

	return nil
}

func (m *Manager) Set(res *ResourceConfig) error {
	for _, subSysIns := range subsystemIns {
		if err := subSysIns.Set(m.Path, res); err != nil {
			klog.Errorf("Failed to set cgroup to %s: %v", m.Path, err)
		}
	}

	return nil
}

func (m *Manager) Destroy() error {
	for _, subSysIns := range subsystemIns {
		if err := subSysIns.Remove(m.Path); err != nil {
			klog.Errorf("Failed to remove cgroup %s: %v", m.Path, err)
		}
	}

	return nil
}
