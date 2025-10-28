package du

import (
	"fmt"
	"testing"
)

func TestNewDiskUsage(t *testing.T) {
	var usage = NewDiskUsage(".")
	if usage == nil {
		t.FailNow()
	}
	fmt.Println("Free:", VisualSize(usage.Free()))
	fmt.Println("Available:", VisualSize(usage.Available()))
	fmt.Println("Size:", VisualSize(usage.Size()))
	fmt.Println("Used:", VisualSize(usage.Used()))
	fmt.Println("Usage:", usage.Usage()*100, "%")
}
