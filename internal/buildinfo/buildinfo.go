package buildinfo

import "time"

var (
	GitTag    string
	BuildDate = time.Now().Format("2006-01-02")
)
