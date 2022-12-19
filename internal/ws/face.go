package ws

import (
	"gitlab.com/fastogt/gofastocloud/gofastocloud/media"
	res "gitlab.com/fastogt/machine-learning/face_detection/golang/internal/resolver"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/logger"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/utils"
)

type FaceRecognitionWSClient struct {
	resolver res.IResolver
}

func NewFaceRecognitionWSClient(resolver res.IResolver) *FaceRecognitionWSClient {
	return &FaceRecognitionWSClient{
		resolver,
	}
}

func (w *FaceRecognitionWSClient) OnStreamMlNotification(notif media.MlNotificationInfo) {
	for _, image := range notif.Images {
		if len(image.Layers) == 0 {
			continue
		}
		layer := image.Layers[0]
		min, sec := utils.GetMinSecTimestamp(image.Timestamp)
		ans, dist := w.resolver.Resolve(layer.Data)
		if ans == nil {
			continue
		}
		logger.Infof("%d:%d | %s | %f\n", min, sec, *ans, dist)
	}
}

func (w *FaceRecognitionWSClient) OnStreamStatistic(stat media.StreamStatisticInfo) {}

func (w *FaceRecognitionWSClient) OnStreamQuit(quit media.QuitStatusInfo) {}
