package du

import (
	"fmt"
	"testing"
)

const (
	KB float32 = 1024
	MB float32 = 1024 * 1024
	TB float32 = 1024 * 1024 * 1024
)

func VisualSize(size uint64) string {
	var sf = float32(size)
	if sf < KB*1.5 {
		return fmt.Sprintf("%d bytes", size)
	}
	if sf < MB*0.9 {
		return fmt.Sprintf("%.3g KB", sf/KB)
	}
	if sf < TB*0.9 {
		return fmt.Sprintf("%.3g MB", sf/MB)
	}
	return fmt.Sprintf("%.3g TB", sf/TB)
}

func TestNewDiskUsage(t *testing.T) {
	var usage = NewDiskUsage(".")
	if usage == nil {
		t.FailNow()
	}
	fmt.Println("Free:", VisualSize(usage.Free()))
	fmt.Println("Available:", VisualSize(usage.Available()))
	fmt.Println("Size:", VisualSize(usage.Size()))
	fmt.Println("Used:", VisualSize(usage.Used()))
	fmt.Printf("Usage: %.3g%%\n", usage.Usage()*100)
}
