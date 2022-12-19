package ws

import (
	"fmt"

	"gitlab.com/fastogt/gofastocloud/gofastocloud/media"
)

type LpdWSClient struct{}

func NewLpdWSClient() *LpdWSClient {
	return &LpdWSClient{}
}

func (w *LpdWSClient) OnStreamMlNotification(notif media.MlNotificationInfo) {
	for _, image := range notif.Images {
		labels := image.Labels
		if len(labels) != 0 {
			fmt.Println(labels)
		}
	}
}

func (w *LpdWSClient) OnStreamStatistic(stat media.StreamStatisticInfo) {}

func (w *LpdWSClient) OnStreamQuit(quit media.QuitStatusInfo) {}
