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

	command := delay.Message{}

	flag.StringVar(&command.AuthToken, "auth", "", "https://developer.atlassian.com/hipchat/guide/hipchat-rest-api/api-access-tokens#APIaccesstokens-userUsertoken")
	flag.StringVar(&command.Room, "room", "", "room id or name")
	flag.DurationVar(&command.NeedSilence, "after-silence", 1 * time.Minute, "Don't post until the silence of the duration")
	on := "now"
	flag.StringVar(&on, "on", on, "when to post the message, HH:MM or duration")

	flag.Parse()

	if "" == command.Room {
		log.Fatal("Need room")
	}

	if "" == command.AuthToken {
		flag.PrintDefaults()
		log.Fatal("Need authentication token")
		return
	}

	if "" == on || "now" == on {
		command.On = time.Now()
	} else {
		command.On, err = time.Parse("15:04", on)
		if err != nil {
			offset, err := time.ParseDuration(on)
			if err != nil {
				log.Fatal("Could not parse 'on': ", on, " expected HH:SS or duration: 5m3s etc.")
				return
			}
			command.On = time.Now().Add(offset)
		} else {
			command.On = util.MergeDateTime(time.Now(), command.On)
		}
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
