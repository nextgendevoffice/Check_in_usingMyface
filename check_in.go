package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gocv.io/x/gocv"
)

const telegramBotToken = "INPUT_YOUR_TOKEN"

const targetGroupID = int64("INPUT_YOUR_GROUPID")

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Photo != nil {
			processPhoto(bot, update)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please send a photo of your face for check-in.")
			bot.Send(msg)
		}
	}
}

func processPhoto(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	photo := (*update.Message.Photo)[0]
	fileURL, _ := bot.GetFileDirectURL(photo.FileID)

	res, err := http.Get(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	imgBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		log.Fatal(err)
	}

	mat, err := gocv.ImageToMatRGB(img)
	if err != nil {
		log.Fatal(err)
	}
	defer mat.Close()

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load("haarcascade_frontalface_default.xml")

	faces := classifier.DetectMultiScale(mat)
	if len(faces) == 1 {
		// Forward the user's picture to the target group
		forward := tgbotapi.NewForward(targetGroupID, update.Message.Chat.ID, update.Message.MessageID)
		bot.Send(forward)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Check-in successful. Welcome to work!")
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Face not detected. Please try again.")
		bot.Send(msg)
	}
}
