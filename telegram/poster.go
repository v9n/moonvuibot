package telegram

import (
	"encoding/json"
	"fmt"
)

/*
===============================================================================
API Method Definition

These are method which map directly to API. Hopefully I can find a way to eliminate those and using only GET and POST to make request to end point directly and somehow parse their response data
===============================================================================
*/

func (api *Api) GetUpdates(handler UpdateHandler, option map[string]string) {
	jsonData, err := api.get("getUpdates", option)

	if nil != err {
		fmt.Println(err)
	}
	var r UpdateResult
	err = json.Unmarshal(jsonData, &r)
	if r.Ok {
		handler(r.Result)
	} else {
		fmt.Println(string(jsonData))
	}
}

/*
* Setup a webhook to receive update with your callback
 */
func (api *Api) getWebhook(url string) {
}

func (api *Api) GetMe() (*User, error) {
	jsonData, err := api.get("getMe", nil)

	if err != nil {
		return nil, err
	}

	var u User
	u1 := parseResultData(jsonData, &u)
	return u1.(*User), nil
}

func (api *Api) SendMessage(form map[string]string) (*Message, error) {
	jsonData, err := api.post("sendMessage", form)

	if err != nil {
		return nil, err
	}

	var m Message
	u1 := parseResultData(jsonData, &m)
	return u1.(*Message), nil
}

func (api *Api) SendPhoto(form map[string]string) (*Message, error) {
	jsonData, err := api.post("sendPhoto", form)

	if err != nil {
		return nil, err
	}

	var m Message
	u1 := parseResultData(jsonData, &m)
	return u1.(*Message), nil
}
