package messaging

import (
	"encoding/json"
	"net/http"
	"os"

	result "github.com/heaptracetechnology/microservice-intercom/result"
	intercom "gopkg.in/intercom/intercom-go.v2"
)

type Message struct {
	Subject string      `json:"subject,omitempty"`
	Body    string      `json:"body,omitempty"`
	UserID  string      `json:"user_id,omitempty"`
	From    json.Number `json:"from,omitempty"`
	To      string      `json:"to,omitempty"`
	Email   string      `json:"email,omitempty"`
}

//Create User
func CreateUser(responseWriter http.ResponseWriter, request *http.Request) {

	responseWriter.Header().Set("Content-Type", "application/json")
	var accessToken = os.Getenv("ACCESS_TOKEN")

	ic := intercom.NewClient(accessToken, "")

	decoder := json.NewDecoder(request.Body)

	var param *intercom.User
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	savedUser, err := ic.Users.Save(param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	bytes, _ := json.Marshal(savedUser)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//InApp Messaage
func SendInAppMessage(responseWriter http.ResponseWriter, request *http.Request) {

	responseWriter.Header().Set("Content-Type", "application/json")
	var accessToken = os.Getenv("ACCESS_TOKEN")

	ic := intercom.NewClient(accessToken, "")

	decoder := json.NewDecoder(request.Body)
	var param *Message
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	msg := intercom.NewInAppMessage(intercom.Admin{ID: param.From}, intercom.Contact{Email: param.To}, param.Body)
	savedMessage, err := ic.Messages.Save(&msg)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(savedMessage)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Email Message
func SendEmailMessage(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	ic := intercom.NewClient(accessToken, "")

	decoder := json.NewDecoder(request.Body)
	var param *Message
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	msg := intercom.NewEmailMessage(intercom.PERSONAL_TEMPLATE, intercom.User{UserID: param.UserID}, intercom.User{Email: param.To}, param.Subject, param.Body)
	savedMessage, err := ic.Messages.Save(&msg)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(savedMessage)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//User Message
func SendUserMessage(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	ic := intercom.NewClient(accessToken, "")

	decoder := json.NewDecoder(request.Body)
	var param *Message
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	msg := intercom.NewUserMessage(intercom.User{Email: param.Email}, param.Body)
	savedMessage, err := ic.Messages.Save(&msg)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(savedMessage)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
