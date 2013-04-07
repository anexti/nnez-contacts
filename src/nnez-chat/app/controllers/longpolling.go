package controllers

import (
	"github.com/robfig/revel"
	"nnez-chat/app/chat"
)

type LongPolling struct {
	*revel.Controller
}

func (c LongPolling) Room(user string) revel.Result {
	chat.Join(user)
	return c.Render(user)
}

func (c LongPolling) Say(user, message string) revel.Result {
	chat.Say(user, message)
	return nil
}

func (c LongPolling) WaitMessages(lastReceived int) revel.Result {
	subscription := chat.Subscribe()
	defer subscription.Cancel()

	var events []chat.Event
	for _, event := range subscription.Archive {
		if event.Timestamp > lastReceived {
			events = append(events, event)
		}
	}

	if len(events) > 0 {
		return c.RenderJson(events)
	}

	event := <-subscription.New
	return c.RenderJson([]chat.Event{event})
}

func (c LongPolling) Leave(user string) revel.Result {
	chat.Leave(user)
	return c.Redirect(Application.Index)
}
