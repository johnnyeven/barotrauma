package main

import (
	"github.com/johnnyeven/barotrauma/modules/core"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(
		&logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05.000",
		},
	)

	world := core.NewWorld(50)
	for {
		world.Update()
	}
}
