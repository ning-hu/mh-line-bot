package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ning-hu/mh-line-bot/linebot"
)

func main() {
	secret := os.Getenv("LINE_SECRET")
	accessToken := os.Getenv("LINE_ACCESS_TOKEN")

	bot, err := linebot.New(secret, accessToken)
	if err != nil {
		log.Fatal("Error creating a new http handler")
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	http.HandleFunc("/.well-known/acme-challenge/iejmNt4sz-ZKxe9H0ExzIT8hguHWtbhgBQL9UeLSxA8", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "iejmNt4sz-ZKxe9H0ExzIT8hguHWtbhgBQL9UeLSxA8.PXkXecLYabBTs7tPtwiUDttGgWLCXL3AGQljK8RLd84")
	})

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
