package logic

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"hanye/pkg/dates"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"hanye/service/order/api/internal/svc"
)

type WebsocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebsocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebsocketLogic {
	return &WebsocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebsocketLogic) Websocket(writer http.ResponseWriter, request *http.Request, userId string) error {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
	}
	node := &svc.Node{
		UserId:        userId,
		Conn:          conn,
		Addr:          conn.RemoteAddr().String(),
		HeartbeatTime: dates.NowTimestamp(),
		LoginTime:     dates.NowTimestamp(),
		DataQueue:     make(chan []byte, 500),
	}
	l.svcCtx.Lock()
	l.svcCtx.ClientMap[userId] = node
	l.svcCtx.Unlock()
	go sendProc(node)
	return nil
}

func sendProc(node *svc.Node) {
	for {
		data := <-node.DataQueue
		fmt.Println("[ws]sendProc >>>> msg :", string(data))
		err := node.Conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
