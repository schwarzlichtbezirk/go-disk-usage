package du

import "fmt"

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB

	cap1 = 1.5
	cap2 = 0.85
)

func VisualSize(size uint64) string {
	var sf = float32(size)
	if sf < KB*cap1 {
		return fmt.Sprintf("%4d B ", size)
	}
	if sf < MB*cap2 {
		return fmt.Sprintf("%4.3g KB", sf/KB)
	}
	if sf < MB {
		return fmt.Sprintf("%4.2g MB", sf/MB)
	}
	if sf < GB*cap2 {
		return fmt.Sprintf("%4.3g MB", sf/MB)
	}
	if sf < GB {
		return fmt.Sprintf("%4.2g GB", sf/GB)
	}
	if sf < TB*cap2 {
		return fmt.Sprintf("%4.3g GB", sf/GB)
	}
	if sf < TB {
		return fmt.Sprintf("%4.2g TB", sf/TB)
	}
	if sf < 1024*TB*cap2 {
		return fmt.Sprintf("%4.3g TB", sf/TB)
	}
	return fmt.Sprintf("%4d TB", uint64(sf/TB))
}
