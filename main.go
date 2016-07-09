package main

import (
	"github.com/tbruyelle/hipchat-go/hipchat"
	"github.com/pkg/errors"
	"time"
	"os"
	"io/ioutil"
	"flag"
	"log"
)

func main() {

	authToken := flag.String("auth", "", "https://developer.atlassian.com/hipchat/guide/hipchat-rest-api/api-access-tokens#APIaccesstokens-userUsertoken")
	room := flag.String("room", "", "room id or name")

	flag.Parse()

	if "" == *room {
		log.Fatal("Need room")
	}

	if "" == *authToken {
		flag.PrintDefaults()
		log.Fatal("Need authentication token")
		return
	}

	c := hipchat.NewClient(*authToken)

	message, err := ioutil.ReadAll(os.Stdin)
	if (err != nil) {
		panic(errors.Wrap(err, "reading stdin failed"))
	}

	/*
	rooms, _, err := c.Room.List()
	if err != nil {
		panic(err)
	}
	*/

	//fmt.Printf("%+v\n", rooms)

	/*
	for _, room := range rooms.Items {
		fmt.Printf("%+v\n", room)
	}
	*/

	/*
	latest, _, err := c.Room.Latest(*room, &hipchat.LatestHistoryOptions{NotBefore:"7187aa29-7c2d-4850-a37f-5b03e86dc98c"})

	if err != nil {
		panic(err)
	}

	var latestMessageId string
	for _, l := range latest.Items {
		latestMessageId = l.ID
		//fmt.Printf("%+v\n", l)
	}
	fmt.Println(latestMessageId)
	*/

	time.Sleep(3 * time.Second)
	_, err = c.Room.Message(*room, &hipchat.RoomMessageRequest{Message: string(message)})
	if nil != err {
		panic(err)
	}
}