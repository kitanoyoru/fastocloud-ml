package ws

import (
	"fmt"

	"gitlab.com/fastogt/gofastocloud/gofastocloud/media"
	res "gitlab.com/fastogt/machine-learning/face_detection/golang/internal/resolver"
	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/utils"
)

type WSClient struct {
	resolver res.IResolver
}

func NewWSClient(resolver res.IResolver) *WSClient {
	return &WSClient{
		resolver,
	}
}

func (w *WSClient) OnStreamMlNotification(notif media.MlNotificationInfo) {
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
		fmt.Printf("%d:%d | %s | %f\n", min, sec, *ans, dist)
	}
}

func (w *WSClient) OnStreamStatistic(stat media.StreamStatisticInfo) {}

func (w *WSClient) OnStreamQuit(quit media.QuitStatusInfo) {}
