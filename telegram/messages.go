package telegram

import (
	"pipeline-bot/bitbucket"
	"log"
	"pipeline-bot/resources"
)

type WebHookMessage struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	ImageData []byte `json:"-"`
}

func (msg *WebHookMessage) String() string {
	return msg.Title + msg.Body
}

func bold(text string) string {
	return "**" + text + "**"
}

func italic(text string) string {
	return "__" + text + "__"
}

func href(name, url string) string {
	return "[" + url + "](" + name + ")"
}

func NewPipelineInProgressMessage(webHook *bitbucket.WebHook) *WebHookMessage {
	return &WebHookMessage{
		Title: webHook.Repository.FullName + "\n\n",
		Body: webHook.CommitStatus.Name + " is in progress now." + "\n" +
			"Initialized it by " + webHook.CommitStatus.Commit.Author.User.DisplayName,
		ImageData: nil,
	}
}

func NewPipelineSuccessfulMessage(webHook *bitbucket.WebHook) *WebHookMessage {
	data, err := resources.Asset("resources/sunny.jpg")
	if err != nil {
		log.Println("file " + "resources/sunny.jpg" + " not found")
	}

	return &WebHookMessage{
		Title: webHook.Repository.FullName + "\n\n",
		Body:  webHook.CommitStatus.Name + " is successful!",
		ImageData: data,
	}
}

func NewPipelineFailureMessage(webHook *bitbucket.WebHook) *WebHookMessage {
	data, err := resources.Asset("resources/cloudy.jpg")
	if err != nil {
		log.Println("file " + "resources/sunny.jpg" + " not found")
	}

	return &WebHookMessage{
		Title: webHook.Repository.FullName + "\n\n",
		Body:  webHook.CommitStatus.Name + " is failed!",
		ImageData: data,
	}
}
