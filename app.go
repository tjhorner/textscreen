package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/kevinburke/twilio-go"
)

type TextScreen struct {
	Storage ConversationStorageAdapter
	Twilio  *twilio.Client
	Config  *Config
}

func (app *TextScreen) HandleError(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(app.Config.Responses.Error))
}

func (app *TextScreen) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	defer req.Body.Close()
	if err != nil {
		app.HandleError(w, req)
		return
	}

	phoneNumber := req.Form.Get("From")

	var conversation *Conversation
	exists, err := app.Storage.Exists(phoneNumber)
	if err != nil {
		app.HandleError(w, req)
		return
	}

	if exists {
		conversation, err = app.Storage.Get(phoneNumber)
		if err != nil {
			app.HandleError(w, req)
			return
		}
	} else {
		conversation = NewConversation(phoneNumber)
	}

	response := app.Respond(&req.Form, conversation)

	err = app.Storage.Save(conversation)
	if err != nil {
		app.HandleError(w, req)
		return
	}

	if response == "" {
		// Empty response
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<Response></Response>"))
	} else {
		w.Write([]byte(response))
	}
}

func (app *TextScreen) Respond(form *url.Values, conversation *Conversation) string {
	switch conversation.State {
	case Started:
		conversation.State = AskedForName
		go app.Twilio.Messages.SendMessage(
			app.Config.Twilio.FromPhoneNumber,
			app.Config.Twilio.ToPhoneNumber,
			fmt.Sprintf(
				"Starting SMS screen from %s",
				conversation.PhoneNumber,
			), nil,
		)
		return app.Config.Responses.AskForName
	case AskedForName:
		conversation.Name = form.Get("Body")
		conversation.State = AskedForPurpose
		return fmt.Sprintf(app.Config.Responses.AskForPurpose, conversation.Name)
	case AskedForPurpose:
		conversation.Purpose = form.Get("Body")
		conversation.State = Complete
		go app.Twilio.Messages.SendMessage(
			app.Config.Twilio.FromPhoneNumber,
			app.Config.Twilio.ToPhoneNumber,
			fmt.Sprintf(
				"SMS Screen from %s (%s): %s",
				conversation.Name,
				conversation.PhoneNumber,
				conversation.Purpose,
			), nil,
		)
		return app.Config.Responses.Complete
	case Complete:
		return ""
	}

	return app.Config.Responses.UnknownState
}
