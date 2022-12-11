package app

import (
	fc "gitlab.com/fastogt/gofastocloud_http/gofastocloud_http"

	"gitlab.com/fastogt/machine-learning/face_detection/golang/internal/resolver"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/internal/ws"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/config"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/utils"
)

type App struct {
	fastocloudConfig *config.FastocloudConfig
	resolverConfig   *config.ResolverConfig
}

func NewApp() *App {
	fc, rc := config.ReadConfigFile()

	return &App{
		fastocloudConfig: fc,
		resolverConfig:   rc,
	}
}

func (a *App) Run() {
	fastocloud := fc.NewFastoCloud(
		a.fastocloudConfig.Endpoint,
		&a.fastocloudConfig.Login,
		&a.fastocloudConfig.Password,
	)

	embeddings := utils.LoadNPZFile(a.resolverConfig.PathToEmbeddings)
	names := utils.LoadNamesFile(a.resolverConfig.PathToNames)

	r := resolver.NewFaceResolver(embeddings, names)

	wsClient := ws.NewWSClient(r)

	fastocloud.Run(wsClient)
}
