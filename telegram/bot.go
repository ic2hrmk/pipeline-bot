package telegram

import (
	"gopkg.in/telegram-bot-api.v4"

	"log"
	"strconv"
)

const timeout = 60

var bot *tgbotapi.BotAPI
var chatMap map[int64]struct{}

func NewBot(token string) (err error) {
	//	Bot initializing
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//	Chat map initializing
	chatMap = make(map[int64]struct{})

	//	Subscribing for updates
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return
	}

	go handleUpdates(updates)

	return
}

func handleUpdates(updates <-chan tgbotapi.Update) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if _, isRegisteredChat := chatMap[update.Message.Chat.ID]; !isRegisteredChat {
			chatMap[update.Message.Chat.ID] = struct{}{}
			log.Println("-> New chat registered: " + strconv.FormatInt(update.Message.Chat.ID, 10))
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func Broadcast(hookMessage *WebHookMessage) (err error) {
	for chatId, _ := range chatMap {
		switch {
		case hookMessage.ImageData != nil:
			messageConfig, err := NewImageMessage(chatId, hookMessage.String(), hookMessage.ImageData)
			if err != nil {
				return err
			}

			_, err = bot.Send(messageConfig)
			if err != nil {
				return err
			}
		default:
			messageConfig, err := NewTextMessage(chatId, hookMessage.String())
			if err != nil {
				return err
			}

			_, err = bot.Send(messageConfig)
			if err != nil {
				return err
			}
		}
	}

	return
}

func NewImageMessage(chatId int64, caption string, image []byte) (messageConfig tgbotapi.PhotoConfig, err error) {
	messageConfig = tgbotapi.NewPhotoUpload(
		chatId,
		tgbotapi.FileBytes{
			Name:  "image.jpg",
			Bytes: image,
		})

	messageConfig.Caption = caption

	return
}

func NewTextMessage(chatId int64, text string) (messageConfig tgbotapi.MessageConfig, err error) {
	messageConfig = tgbotapi.NewMessage(chatId, text)
	messageConfig.ParseMode = tgbotapi.ModeMarkdown

	return
}
