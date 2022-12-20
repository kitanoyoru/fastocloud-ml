package handler

import (
	"gitlab.com/fastogt/machine-learning/face_detection/golang/internal/resolver"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/internal/ws"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/config"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/utils"
)

func PrepareFaceRecognitionClient(c *config.FaceRecognitionResolverConfig) *ws.FaceRecognitionWSClient {
	embeddings := utils.LoadNPZFile(c.PathToEmbeddings)
	names := utils.LoadNamesFile(c.PathToNames)

	r := resolver.NewFaceResolver(embeddings, names)

	wsClient := ws.NewFaceRecognitionWSClient(r)

	return wsClient
}

func PrepareLPDClient() *ws.LpdWSClient {
	return ws.NewLpdWSClient()
}
