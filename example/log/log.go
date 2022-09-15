package main

import (
	"base-frame/log"
)

func main() {
	log.Setup(log.SetProjectName("Test"))
	log.Debug("haha")
}
