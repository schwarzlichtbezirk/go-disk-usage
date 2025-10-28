package main

import (
	"fmt"

	"github.com/schwarzlichtbezirk/go-disk-usage/du"
)

const MB uint64 = 1024 * 1024

func main() {
	var usage = du.NewDiskUsage(".")
	if usage == nil {
		fmt.Println("Can not obtain disk info")
		return
	}
	fmt.Println("Free:", usage.Free()/MB)
	fmt.Println("Available:", usage.Available()/MB)
	fmt.Println("Size:", usage.Size()/MB)
	fmt.Println("Used:", usage.Used()/MB)
	fmt.Println("Usage:", usage.Usage()*100, "%")
}
