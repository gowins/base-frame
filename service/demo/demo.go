package main

import (
	base_frame "base-frame"
	"base-frame/cmd/zerocmd"
)

func main() {
	zc := zerocmd.New()
	base_frame.Start("test", "base-frame", zc)
}
