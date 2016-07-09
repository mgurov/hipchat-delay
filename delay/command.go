package delay

import (
	"time"
	"github.com/tbruyelle/hipchat-go/hipchat"
	"github.com/pkg/errors"
	"log"
)

type Message struct {
	Text        string
	Room        string
	AuthToken   string
	NeedSilence time.Duration
	On          time.Time
}

func (c Message) Send() error {
	sender := sender{c, hipchat.NewClient(c.AuthToken), time.Now()}
	return sender.send()
}

type sender struct {
	Message
	cli *hipchat.Client
	now time.Time
}

func (s sender) send() error {
	s.waitForTheFuture()
	err := s.waitForSilence()
	if nil != err {
		return errors.Wrap(err, "While waiting for silence")
	}
	_, err = s.cli.Room.Message(s.Room, &hipchat.RoomMessageRequest{Message: s.Message.Text})
	return errors.Wrap(err, "Could not send message")
}

func (s sender) waitForSilence() error {
	if s.NeedSilence <= 0 {
		return nil
	}
	log.Println("Need silence at least", s.NeedSilence)

	// check the time passed since start and wait for the silence interval
	// otherwise little point in checking last timestamps yet
	// except for the unlikely case of empty chat history, where silence request would not make much sense in the first place
	time.Sleep(s.now.Add(s.NeedSilence).Sub(time.Now()))

	for {
		durationSinceLastMessage, err := s.readDurationSinceLastMessage()
		if err != nil {
			return err
		}
		if durationSinceLastMessage >= s.NeedSilence {
			return nil
		} else {
			needToWait := s.NeedSilence - durationSinceLastMessage
			log.Println("Waiting", needToWait, "more")
			time.Sleep(needToWait)
		}
	}
}

// returns s.NeedSliencde + 1 if no history records found
func (s sender) readDurationSinceLastMessage() (time.Duration, error) {
	history, _, err := s.cli.Room.Latest(s.Room, &hipchat.LatestHistoryOptions{MaxResults:1})
	if err != nil {
		return 0, errors.Wrap(err, "reading room last message")
	}

	if (len(history.Items) == 0) {
		return s.NeedSilence + 1, nil
	}

	if (len(history.Items) > 1) {
		log.Println("Unexpected volume of history items received: ", len(history.Items))
	}

	lastMessageDate := history.Items[0].Date
	lastMessageTimestamp, err := time.Parse(time.RFC3339, lastMessageDate)
	if err != nil {
		return 0, errors.Wrap(err, "Parsing message date: " + lastMessageDate)
	}
	return time.Now().Sub(lastMessageTimestamp), nil
}

func (s sender) waitForTheFuture() {
	howLongToWait := s.On.Sub(s.now)
	if (howLongToWait > 0) {
		log.Println("Will wait", howLongToWait, "till", s.On)
		time.Sleep(howLongToWait)
	}
}
