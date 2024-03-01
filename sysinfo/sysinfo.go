package sysinfo

import (
	"syscall"
)

type SysInfo struct {
	MaxFD syscall.Rlimit `json:"maxfd"`
}

func newSysInfo() (*SysInfo, error) {
	obj := SysInfo{}
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &obj.MaxFD)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
