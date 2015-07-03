package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Api struct {
	config *Config
}

type (
	UpdateParam struct {
		offset  int
		limit   int
		timeout int
	}
)
type UpdateHandler func([]*Message)

func NewApi(config *Config) (*Api, error) {
	api := new(Api)
	api.config = config
	return api, nil
}

func (api *Api) buildUrl(method string, option map[string]string) string {
	q := url.Values{}
	for k, v := range option {
		q.Add(k, v)
	}
	return api.config.Endpoint(method) + "?" + q.Encode()
}

/*
* Setup a webhook to receive update with your callback
 */
func (api *Api) getWebhook(url string) {
}

func (api *Api) GetUpdates(handler UpdateHandler, option map[string]string) {
	q := url.Values{}
	for k, v := range option {
		q.Add(k, v)
	}

	response, err := http.Get(api.buildUrl("getUpdates", option))
	defer response.Body.Close()
	jsonData, err := ioutil.ReadAll(response.Body)
	if nil != err {
		fmt.Println(err)
	}
	var m UpdateResult
	err = json.Unmarshal(jsonData, &m)
	fmt.Println(m)
	fmt.Println(string(jsonData))
}
