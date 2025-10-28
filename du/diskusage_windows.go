package du

import (
	"syscall"
	"unsafe"
)

type DiskUsage struct {
	freeBytes  int64
	totalBytes int64
	availBytes int64
}

// NewDiskUsages returns an object holding the disk usage of volumePath
// or nil in case of error (invalid path, etc)
func NewDiskUsage(volumePath string) *DiskUsage {
	var h = syscall.MustLoadDLL("kernel32.dll")
	var c = h.MustFindProc("GetDiskFreeSpaceExW")

	var du DiskUsage

	var pv, err = syscall.UTF16PtrFromString(volumePath)
	if err != nil {
		return nil
	}
	var ok, _, _ = c.Call(
		uintptr(unsafe.Pointer(pv)),
		uintptr(unsafe.Pointer(&du.freeBytes)),
		uintptr(unsafe.Pointer(&du.totalBytes)),
		uintptr(unsafe.Pointer(&du.availBytes)))
	if ok == 0 {
		return nil
	}

	return &du
}

// Free returns total free bytes on file system
func (du *DiskUsage) Free() uint64 {
	return uint64(du.freeBytes)
}

// Available returns total available bytes on file system to an unprivileged user
func (du *DiskUsage) Available() uint64 {
	return uint64(du.availBytes)
}

// Size returns total size of the file system
func (du *DiskUsage) Size() uint64 {
	return uint64(du.totalBytes)
}

// Used returns total bytes used in file system
func (du *DiskUsage) Used() uint64 {
	return du.Size() - du.Free()
}

// Usage returns percentage of use on the file system
func (du *DiskUsage) Usage() float32 {
	return float32(du.Used()) / float32(du.Size())
}
