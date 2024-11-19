package main

import (
	"net/url"
	"webscraper/app"
	"webscraper/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.API_TOKEN)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		userText := update.Message.Text

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, userText)

		switch update.Message.Command() {
		case "start":
			msg.Text = app.StartMessage
		default:
			if isValidURL(userText) {
				msg.Text = "It's valid link"
				urls := app.Grabber(userText)
				for _, url := range urls {
					photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(url))
					if _, err := bot.Send(photo); err != nil {
						panic(err)
					}
				}

			} else {
				msg.Text = "It's not a link, maybe jsut a text or media"
			}
		}

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}

	}
}

func isValidURL(str string) bool {

	u, err := url.ParseRequestURI(str)

	return err == nil && u.Scheme != "" && u.Host != ""
}
