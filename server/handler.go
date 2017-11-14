package server

import (
	"pipeline-bot/bitbucket"
	"pipeline-bot/telegram"

	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

func WebHookHandler(rw http.ResponseWriter, req *http.Request) {
	log.Println("New web hook from: " + req.RequestURI)

	//	Read request body
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//	Deserialize request
	webHook := &bitbucket.WebHook{}
	err = json.Unmarshal(body, &webHook)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//	Build message
	var message *telegram.WebHookMessage

	switch webHook.CommitStatus.State {
	case bitbucket.PIPELINE_IN_PROGRESS:
		message = telegram.NewPipelineInProgressMessage(webHook)
	case bitbucket.PIPELINE_SUCCESSFUL:
		message = telegram.NewPipelineSuccessfulMessage(webHook)
	default:
		message = telegram.NewPipelineFailureMessage(webHook)
	}

	//	Broadcast message
	err = telegram.Broadcast(message)
	if err != nil {
		log.Println(err.Error())
		return
	}
}