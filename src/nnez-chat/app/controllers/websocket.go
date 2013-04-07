package controllers 


import (
	"code.google.com/p/go.net/websocket"
	"github.com/robfig/revel"
	"nnez-chat/app/chat"
)

type WebSocket struct {
	*revel.Controller
}


func (c WebSocket) Room(user string) revel.Result {
	return c.Render(user)
}

func (c WebSocket) RoomSocket(user string, ws *websocket.Conn) revel.Result {
	subscription := chat.Subscribe()
	defer subscription.Cancel()

	chat.Join(user)
	defer chat.Leave(user)

	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			return nil
		}
	}

	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				return nil
			}
		case msg, ok := <-newMessages:
			if !ok {
				return nil
			}
			chat.Say(user, msg)
		}
	}
	return nil
}
