package app

import (
	fc "gitlab.com/fastogt/gofastocloud_http/gofastocloud_http"

	"gitlab.com/fastogt/machine-learning/face_detection/golang/internal/resolver"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/internal/ws"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/config"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/logger"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/utils"
)

type App struct {
	cliConfig                     *config.CliConfig
	fastocloudConfig              *config.FastocloudConfig
	faceRecognitionResolverConfig *config.ResolverConfig
}

func NewApp(cc config.CliConfig) *App {
	fc, fr := config.ReadConfigFile()

	return &App{
		cliConfig:                     cc,
		fastocloudConfig:              fc,
		faceRecognitionResolverConfig: fr,
	}
}

func (a *App) Run() {
	fastocloud := fc.NewFastoCloud(
		a.fastocloudConfig.Endpoint,
		&a.fastocloudConfig.Login,
		&a.fastocloudConfig.Password,
	)

	// TODO: Load files through goroutines
	embeddings := utils.LoadNPZFile(a.resolverConfig.PathToEmbeddings)
	names := utils.LoadNamesFile(a.resolverConfig.PathToNames)

	r := resolver.NewFaceResolver(embeddings, names)

	wsClient := ws.NewWSClient(r)

	logger.Infof("Start listening Fastocloud websocket on host %s", a.fastocloudConfig.Endpoint)
	fastocloud.Run(wsClient)
	logger.Info("Stop listening Fastocloud websocket")
}
