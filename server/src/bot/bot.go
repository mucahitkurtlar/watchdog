package bot

import (
	"allesfresser/watchdog/esp"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Creates a Telegram reply keyboard
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("LED Status"),
		tgbotapi.NewKeyboardButton("LED On"),
		tgbotapi.NewKeyboardButton("LED Off"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Lamp Status"),
		tgbotapi.NewKeyboardButton("Lamp On"),
		tgbotapi.NewKeyboardButton("Lamp Off"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Sensor On"),
		tgbotapi.NewKeyboardButton("Sensor Off"),
	),
)

// Defines Bot struct that inherits Telegram bot
type Bot struct {
	TgBot      *tgbotapi.BotAPI
	chatID     int64
	isSensorOn bool
}

// Creates a Bot object and returns its reference
func CreateBot(secret string) *Bot {
	tgBot, err := tgbotapi.NewBotAPI(secret)
	if err != nil {
		log.Fatal(err)
	}
	tgBot.Debug = false
	log.Printf("Authorized on account %s", tgBot.Self.UserName)
	bot := Bot{TgBot: tgBot}

	return &bot
}

// Ping function for Bot
func (b *Bot) Ping() {
	log.Println("Pong!")
}

// Checks message queue, sets/reads pins status
func (b *Bot) HandleMessages() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5

	updates, err := b.TgBot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		// Ignore non-Message updates
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		log.Printf("Message: %s From %s\n", update.Message.Text, update.Message.From)
		switch update.Message.Text {
		case "/open":
			msg.ReplyMarkup = numericKeyboard
			log.Println("Reply keyboard opened")
		case "/close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			log.Println("Reply keyboard closed")
		case "LED Status":
			if esp.AskStatus("http://led.local/status") {
				msg.Text = "LED is ON"
				log.Println("LED status asked and it's on")
			} else {
				msg.Text = "LED is OFF"
				log.Println("LED status asked and it's off")
			}
		case "LED On":
			if esp.PinAction("http://led.local/on") {
				msg.Text = "Ok. LED turned on"
				log.Println("LED turned on")
			} else {
				msg.Text = "An error has occurred :("
				log.Println("An error has occurred while LED turning on")
			}
		case "LED Off":
			if esp.PinAction("http://led.local/off") {
				msg.Text = "Ok. LED turned off"
				log.Println("LED turned off")
			} else {
				msg.Text = "An error has occurred :("
				log.Println("An error has occurred while LED turning off")
			}
		case "Lamp Status":
			if esp.AskStatus("http://lamp.local/status") {
				msg.Text = "Lamp is ON"
				log.Println("Lamp status asked and it's on")
			} else {
				msg.Text = "Lamp is OFF"
				log.Println("Lamp status asked and it's off")
			}
		case "Lamp On":
			if esp.PinAction("http://lamp.local/on") {
				msg.Text = "Ok. Lamp turned on"
				log.Println("Lamp turned on")
			} else {
				msg.Text = "An error has occurred :("
				log.Println("An error has occurred while lamp turning on")
			}
		case "Lamp Off":
			if esp.PinAction("http://lamp.local/off") {
				msg.Text = "Ok. Lamp turned off"
				log.Println("Lamp turned off")
			} else {
				msg.Text = "An error has occurred :("
				log.Println("An error has occurred while lamp turning off")
			}
		case "Sensor On":
			b.chatID = update.Message.Chat.ID
			b.isSensorOn = true
			log.Println("Motion sensor set to on")
		case "Sensor Off":
			b.isSensorOn = false
			log.Println("Motion sensor set to off")
		}

		if _, err := b.TgBot.Send(msg); err != nil {
			log.Println("The message couldn't be sent! Message: ", msg.Text)
			log.Panic(err)
		}
	}
}

// Sends motion detection message
func (b *Bot) SendMotionMessage() {
	if b.isSensorOn {
		msg := tgbotapi.NewMessage(b.chatID, "Motion detected!")
		if _, err := b.TgBot.Send(msg); err != nil {
			log.Println("Motion detection message couldn't be sent! Message: ", msg.Text)
			return
		}
		log.Println("Motion detection message sent")
	}
}
