package controller

import (
	"backend/ent/gen"
	"backend/internal/core/stack"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WebSocketController struct {
	ControllerServices
}

func NewWebSocketController(controller ControllerServices) WebSocketController {
	return WebSocketController{controller}
}

var (
	upgrader = websocket.Upgrader{}
)

func (c WebSocketController) PileLiveShow(ctx echo.Context) error {
	type liveShowDTO struct {
		FilePath    string `json:"filePath"`
		Legend      string `json:"legend"`
		FileType    string `json:"fileType"`
		CurrentTime int    `json:"currentTime"`
	}

	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	var oldMedia *gen.LiveShow
	for {
		if oldMedia != stack.CurrentMedia {
			media := stack.CurrentMedia
			if media == nil {
				err = ws.WriteJSON("stop")
				if err != nil {
					ctx.Logger().Error(err)
				}
			} else {
				now := time.Now()
				diff := now.Sub(media.StartedTime).Seconds()
				err = ws.WriteJSON(liveShowDTO{
					FilePath:    "/upload/" + media.Edges.Upload.FilePath,
					Legend:      media.Legend,
					FileType:    string(media.Edges.Upload.FileType),
					CurrentTime: int(diff),
				})
				if err != nil {
					ctx.Logger().Error(err)
				}
			}
		}
		oldMedia = stack.CurrentMedia
		_, _, err := ws.ReadMessage()
		if err != nil {
			ctx.Logger().Error(err)
		}
	}
}
