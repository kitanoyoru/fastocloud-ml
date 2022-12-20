package main

import (
	"flag"
	"fmt"

	"gitlab.com/fastogt/machine-learning/face_detection/golang/internal/app"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/config"
)

func main() {
	modeFlag := flag.Int("mode", -1, "Which fastocloud ml mode you want to use")

	flag.Parse()

  fmt.Println(*modeFlag)

	cliConfig := config.NewCliConfig(*modeFlag)

	app := app.NewApp(cliConfig)
	app.Run()
}
