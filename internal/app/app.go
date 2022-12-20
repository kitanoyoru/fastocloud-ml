package app

import (
	fasto "gitlab.com/fastogt/gofastocloud_http/gofastocloud_http"
	fastows "gitlab.com/fastogt/gofastocloud_http/gofastocloud_http/public"

	"gitlab.com/fastogt/machine-learning/face_detection/golang/internal/handler"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/config"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/logger"
)

type App struct {
	cliConfig        *config.CliConfig
	fastocloudConfig *config.FastocloudConfig
	resolverConfig   *config.FaceRecognitionResolverConfig
}

func NewApp(cc *config.CliConfig) *App {
	fc, fr := config.ReadConfigFile()

	return &App{
		cliConfig:        cc,
		fastocloudConfig: fc,
		resolverConfig:   fr,
	}
}

func (a *App) Run() {
	fastocloud := fasto.NewFastoCloud(
		a.fastocloudConfig.Endpoint,
		&a.fastocloudConfig.Login,
		&a.fastocloudConfig.Password,
	)

	var wsClient fastows.WsClient
	switch a.cliConfig.Mode {
	case 1:
		wsClient = handler.PrepareFaceRecognitionClient(a.resolverConfig)
	case 2:
		wsClient = handler.PrepareLPDClient()
	}

	logger.Infof("Start listening Fastocloud websocket on host %s", a.fastocloudConfig.Endpoint)
	fastocloud.Run(wsClient)
	logger.Info("Stop listening Fastocloud websocket")
}
