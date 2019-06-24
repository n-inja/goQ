package main

import (
	"github.com/otiai10/opengraph"
	"github.com/traPtitech/traq-bot"
	"github.com/traPtitech/traq-client"

	"log"
	"os"
	"regexp"
	"sync"
)

var once sync.Once
var URLReg *regexp.Regexp

func extractURL(message string) []string {
	once.Do(func() {
		URLReg = regexp.MustCompile(`(http|https)://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?`)
	})

	us := URLReg.FindAllStringSubmatch(message, -1)
	URLs := make([]string, 0)
	for _, r := range us {
		URLs = append(URLs, r[0])
	}

	return URLs
}

func getOGP(rawURL string) string {
	ogp, err := opengraph.Fetch(rawURL)

	if err != nil {
		return ""
	}

	return ogp.Title + "\n> " + ogp.Description + "\n\n\n"
}

func main() {
	vt := os.Getenv("VERIFICATION_TOKEN")
	at := os.Getenv("ACCESS_TOKEN")
	userID := os.Getenv("USER_ID")
	client := traq.NewClient("q.trap.jp", at)

	handlers := traqbot.EventHandlers{}
	handlers.SetMessageCreatedHandler(func(payload *traqbot.MessageCreatedPayload) {
		if payload.Message.User.ID == userID {
			return
		}

		URLs := extractURL(payload.Message.Text)
		message := ""

		for _, URL := range URLs {
			m := getOGP(URL)
			message += m
		}

		client.PostMessage(payload.Message.ChannelID, message)
	})

	server := traqbot.NewBotServer(vt, handlers)
	log.Fatal(server.ListenAndServe(":3000"))
}
