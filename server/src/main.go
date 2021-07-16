package main

import (
	"log"
	"net/http"
	"time"

	"allesfresser/watchdog/bot"
	"allesfresser/watchdog/routes"

	/*
		secrets/secrets.go

		const BotToken = Your_bot_token_here""
	*/
	"allesfresser/watchdog/secrets"
)

var isMotionDetected bool

//var posts []Post

func main() {
	isMotionDetected = false
	myBot := bot.CreateBot(secrets.BotToken, secrets.ChatID)
	go myBot.HandleMessages()
	go func() {
		lasDetect := time.Now()
		for {
			// If motion detected and last motion was 10 second or more time ago, sends message
			if isMotionDetected && time.Since(lasDetect) > time.Second*10 {
				// Update last motion detection time
				lasDetect = time.Now()
				myBot.SendMotionMessage()
				// Set isMotionDetected false
				isMotionDetected = false
			}
		}
	}()

	r := routes.NewRouter(&isMotionDetected)
	http.Handle("/", r)
	log.Println("Listening on port :8020")
	http.ListenAndServe(":8020", nil)
}
