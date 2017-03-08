package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
	"log"
	"messenger-sdk"
	"net/http"
	"os"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/messenger", handleMessengerWebHook)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func handleMessengerWebHook(w http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)

	bot := msdk.BotAPI{
		Token:       os.Getenv("PAGE_TOKEN"),
		VerifyToken: os.Getenv("HUB_VERIFY_TOKEN"),
		Debug:       true,
		Client:      client,
	}

	switch req.Method {
	case "GET":
		if req.FormValue("hub.verify_token") == bot.VerifyToken {
			w.Write([]byte(req.FormValue("hub.challenge")))
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		return

	case "POST":
		defer req.Body.Close()

		body, _ := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewReader(body))

		if bot.Debug {
			log.Printf("[INFO]%s", body)
		}

		var userResponse msdk.Response

		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&userResponse)

		if userResponse.Object == "page" {
			for _, e := range userResponse.Entries {
				for _, event := range e.Messaging {
					receivedMessage(event, bot)
				}
			}
		}
		w.WriteHeader(http.StatusOK)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func receivedMessage(event msdk.Callback, bot msdk.BotAPI) {
	senderID := event.Sender.ID
	recipientID := event.Recipient.ID
	timeOfMessage := event.Timestamp
	message := event.Message

	if event.IsMessage(){
		if bot.Debug {
			log.Printf("[INFO]Received message for user %d and page %d at %d with message %s", senderID, recipientID, timeOfMessage, message)
		}

		messageText := message.Text
		messageAttachments := message.Attachments

		if &messageText != nil {
			switch messageText {
			case "generic":
				sendGenericMessage(event.Sender, bot)

			default:
				sendTextMessage(event.Sender, messageText, bot)
			}

		} else if &messageAttachments != nil {
			sendTextMessage(event.Sender, "Message with attachement received", bot)
		}
	}else if event.IsPostback(){
		log.Print("[INFO] Received postback for user %d and page %d with payload %s at %d", senderID, recipientID, event.Postback.Payload, timeOfMessage)
		sendTextMessage(event.Sender, "Postback Called", bot)
	}
}
func sendTextMessage(recipient msdk.User, messageText string, bot msdk.BotAPI) {
	bot.Send(recipient, msdk.NewMessage(messageText), msdk.RegularNotif)
}
func receivedPostback(event msdk.Callback){

}

func sendGenericMessage(recipient msdk.User, bot msdk.BotAPI) {

	genericTemp := msdk.NewGenericTemplate()

	riftElement := msdk.Element{
		Title:    "rift",
		Subtitle: "Next-generation virtual reality",
		URL:      "https://www.oculus.com/en-us/rift/",
		ImageURL: "http://messengerdemo.parseapp.com/img/rift.png",
	}

	riftElement.AddButton(
		msdk.NewURLButton("Open web url", "https://www.oculus.com/en-us/rift/"),
		msdk.NewPostbackButton("Call PostBack", "Payload for First bubble"),
	)
	touchElement := msdk.Element{
		Title:    "touch",
		Subtitle: "Your Hands, Now in VR",
		URL:      "https://www.oculus.com/en-us/touch/",
		ImageURL: "http://messengerdemo.parseapp.com/img/touch.png",
	}

	touchElement.AddButton(
		msdk.NewURLButton("Open web url", "https://www.oculus.com/en-us/touch/"),
		msdk.NewPostbackButton("Call PostBack", "Payload for second bubble"),
	)


	genericTemp.AddElement(riftElement,touchElement)

	bot.Send(recipient, genericTemp, msdk.RegularNotif)
}
