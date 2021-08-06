package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/slack-go/slack"
)

func main() {
	token, ok := os.LookupEnv("SLACK_TOKEN")
	if !ok {
		fmt.Println("Missing SLACK_TOKEN in environment")
		os.Exit(1)
	}
	api := slack.New(
		token,
		slack.OptionDebug(false),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			// Replace C2147483705 with your Channel ID

		case *slack.MessageEvent:
			if strings.Contains(ev.Text, "Ping") {
				rtm.SendMessage(rtm.NewOutgoingMessage("Pong", ev.Channel))
			} else if ev.BotID == "B01V8SG8Q94" {

				fmt.Println(ev)
				msg := parseQuote(ev)
				fmt.Println(msg)
				sendSMS(msg)
			}

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:

		}
	}
}
