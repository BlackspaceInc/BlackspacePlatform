package api

import (
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"k8s.io/klog/v2"
)

var wsCon = websocket.Upgrader{}

func (s *Server) sendHostWs(ws *websocket.Conn, in chan interface{}, done chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			var status = struct {
				Time time.Time `json:"ts"`
				Host string    `json:"server"`
			}{
				Time: time.Now(),
				Host: s.config.Hostname,
			}
			in <- status
		case <-done:
			klog.Info("websocket exit")
			return
		}
	}
}

func (s *Server) writeWs(ws *websocket.Conn, in chan interface{}) {
	for {
		select {
		case msg := <-in:
			if err := ws.WriteJSON(msg); err != nil {
				if !strings.Contains(err.Error(), "close") {
					klog.Info("websocket write error", zap.Error(err))
				}
				return
			}
		}
	}
}
