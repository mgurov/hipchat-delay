package delay

import (
	"time"
	"github.com/tbruyelle/hipchat-go/hipchat"
	"github.com/pkg/errors"
)

type Message struct {
	Message   string
	Room      string
	AuthToken string
	Delay     time.Duration
}

func (c Message) Send() error {
	time.Sleep(c.Delay)
	cli := hipchat.NewClient(c.AuthToken)
	_, err := cli.Room.Message(c.Room, &hipchat.RoomMessageRequest{Message: c.Message})
	return errors.Wrap(err, "Could not send message")
}