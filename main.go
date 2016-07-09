package main

import (
	"github.com/pkg/errors"
	"os"
	"io/ioutil"
	"flag"
	"log"
	"github.com/mgurov/hipchat-delay/delay"
)

func main() {

	command := delay.Message{}

	flag.StringVar(&command.AuthToken, "auth", "", "https://developer.atlassian.com/hipchat/guide/hipchat-rest-api/api-access-tokens#APIaccesstokens-userUsertoken")
	flag.StringVar(&command.Room, "room", "", "room id or name")
	flag.DurationVar(&command.NeedSilence, "after-silence", 0, "Don't post until the silence of the duration")

	flag.Parse()

	if "" == command.Room {
		log.Fatal("Need room")
	}

	if "" == command.AuthToken {
		flag.PrintDefaults()
		log.Fatal("Need authentication token")
		return
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
