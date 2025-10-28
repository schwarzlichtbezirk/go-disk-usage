# go-disk-usage

Get disk usage information like how much space is available, free, and used.

## Documentation

Documentation can be [found here](https://godoc.org/github.com/schwarzlichtbezirk/go-disk-usage/du) and there is a [short example](https://github.com/schwarzlichtbezirk/go-disk-usage/blob/master/duexample.go) of how to use this library.

## Compatibility

This works for Windows, MacOS, and Linux although there may some minor variability between what this library reports and what you get from `df`.  This library will maintain reverse compatability, any breaking changes will be made to a forked repository.

## Install

```bash
go get -u "github.com/schwarzlichtbezirk/go-disk-usage/du"
```

## Usage

```go
import "github.com/schwarzlichtbezirk/go-disk-usage/du"
usage := du.NewDiskUsage("/path/to")
```
