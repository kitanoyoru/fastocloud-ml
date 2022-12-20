package ws

import (
	"fmt"

	"gitlab.com/fastogt/gofastocloud/gofastocloud/media"

	"gitlab.com/fastogt/machine-learning/face_detection/golang/pkg/utils"
)

type LpdWSClient struct{}

func NewLpdWSClient() *LpdWSClient {
	return &LpdWSClient{}
}

func (w *LpdWSClient) OnStreamMlNotification(notif media.MlNotificationInfo) {
	for _, image := range notif.Images {
		labels := image.Labels
		for _, label := range labels {
			if len(label.Label) == 7 {
				min, sec := utils.GetMinSecTimestamp(image.Timestamp)
				fmt.Printf("%d:%d | %s %f \n", min, sec, label.Label, label.Probability)
			}
		}
	}
}

func (w *LpdWSClient) OnStreamStatistic(stat media.StreamStatisticInfo) {}

func (w *LpdWSClient) OnStreamQuit(quit media.QuitStatusInfo) {}
