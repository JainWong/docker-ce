package cgroups

import (
	"errors"

	"github.com/dotcloud/docker/pkg/libcontainer/devices"
)

var (
	ErrNotFound = errors.New("mountpoint not found")
)

type Cgroup struct {
	Name   string `json:"name,omitempty"`
	Parent string `json:"parent,omitempty"` // name of parent cgroup or slice

	AllowAllDevices   bool              `json:"allow_all_devices,omitempty"` // If this is true allow access to any kind of device within the container.  If false, allow access only to devices explicitly listed in the allowed_devices list.
	AllowedDevices    []*devices.Device `json:"allowed_devices,omitempty"`
	Memory            int64             `json:"memory,omitempty"`             // Memory limit (in bytes)
	MemoryReservation int64             `json:"memory_reservation,omitempty"` // Memory reservation or soft_limit (in bytes)
	MemorySwap        int64             `json:"memory_swap,omitempty"`        // Total memory usage (memory + swap); set `-1' to disable swap
	CpuShares         int64             `json:"cpu_shares,omitempty"`         // CPU shares (relative weight vs. other containers)
	CpuQuota          int64             `json:"cpu_quota,omitempty"`          // CPU hardcap limit (in usecs). Allowed cpu time in a given period.
	CpuPeriod         int64             `json:"cpu_period,omitempty"`         // CPU period to be used for hardcapping (in usecs). 0 to use system default.
	CpusetCpus        string            `json:"cpuset_cpus,omitempty"`        // CPU to use
	Freezer           string            `json:"freezer,omitempty"`            // set the freeze value for the process

	Slice string `json:"slice,omitempty"` // Parent slice to use for systemd
}

type ActiveCgroup interface {
	Cleanup() error
}
