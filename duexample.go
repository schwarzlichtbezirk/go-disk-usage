package main

import (
	"fmt"

	"github.com/schwarzlichtbezirk/go-disk-usage/du"
)

func main() {
	var usage = du.NewDiskUsage(".")
	if usage == nil {
		fmt.Println("Can not obtain disk info")
		return
	}
	fmt.Println("Free:", du.VisualSize(usage.Free()))
	fmt.Println("Available:", du.VisualSize(usage.Available()))
	fmt.Println("Size:", du.VisualSize(usage.Size()))
	fmt.Println("Used:", du.VisualSize(usage.Used()))
	fmt.Println("Usage:", usage.Usage()*100, "%")
}
