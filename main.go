package main

import (
	"github.com/pkg/errors"
	"os"
	"io/ioutil"
	"flag"
	"log"
	"github.com/mgurov/hipchat-delay/delay"
	"time"
	"github.com/mgurov/hipchat-delay/util"
)

func main() {

	var err error

	command := delay.Message{
		NeedSilence: 1 * time.Minute,
	}

	flag.StringVar(&command.AuthToken, "auth", "", "https://developer.atlassian.com/hipchat/guide/hipchat-rest-api/api-access-tokens#APIaccesstokens-userUsertoken")
	flag.StringVar(&command.Room, "room", "", "room id or name")
	flag.DurationVar(&command.NeedSilence, "silence", command.NeedSilence, "Don't post until the silence of the duration")
	at := ""
	flag.StringVar(&at, "at", at, "when to post the message, HH:MM.")
	var in time.Duration = 0
	flag.DurationVar(&in, "in", in, "when to post the message, duration e.g. 5m")

	flag.Parse()

	if "" == command.Room {
		log.Fatal("Need room")
	}

	if "" == command.AuthToken {
		if command.AuthToken = os.Getenv("HIPCHAT_AUTH_TOKEN"); "" == command.AuthToken {
			flag.PrintDefaults()
			log.Fatal("Need authentication token")
			return
		}
	}

	now := time.Now()
	if "" != at && 0 != in {
		log.Fatal("Either -on or -in should be specified, not both of them")
		return
	} else if "" != at {
		command.On, err = time.Parse("15:04", at)
		if err != nil {
			log.Fatal("Could not parse 'on': ", at, " expected HH:SS")
			return
		}
		command.On = util.MergeDateTime(now, command.On)
		if now.After(command.On) {
			command.On = command.On.AddDate(0, 0, 1)
		}
	} else if 0 != in {
		command.On = now.Add(in)
	} else {
		command.On = now
	}

	message, err := ioutil.ReadAll(os.Stdin)
	if (err != nil) {
		log.Fatal(errors.Wrap(err, "reading stdin failed"))
	}

	command.Text = string(message)

	err = command.Send()
	if err != nil {
		log.Fatal(err)
	}
}
