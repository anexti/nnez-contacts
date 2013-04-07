package controllers

import (
	"github.com/robfig/revel"
	"nnez-chat/app/chat"
)

type Refresh struct {
	*revel.Controller
}

func (c Refresh) Index(user string) revel.Result {
	chat.Join(user)
	return c.Room(user)
}

func (c Refresh) Room(user string) revel.Result {
	subscription := chat.Subscribe()
	defer subscription.Cancel()
	events := subscription.Archive
	for i, _ := range events {
		if events[i].User == user {
			events[i].User = "you"
		}
	}
	return c.Render(user, events)
}

func (c Refresh) Say(user, message string) revel.Result {
	chat.Say(user, message)
	return c.Room(user)
}

func (c Refresh) Leave(user string) revel.Result {
	chat.Leave(user)
	return c.Redirect(Application.Index)
}
