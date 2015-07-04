package telegram

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	//"reflect"
)

type (
	Api struct {
		config *Config
	}
)

/*
Create a new API object
*/
func NewApi(config *Config) (*Api, error) {
	api := new(Api)
	api.config = config
	return api, nil
}

/*
Build URL for an API method call. Note that when using GET, we will also append query parameter to URL
*/
func (api *Api) buildUrl(method string, option map[string]string) string {
	q := url.Values{}
	for k, v := range option {
		q.Add(k, v)
	}
	return api.config.Endpoint(method) + "?" + q.Encode()
}

/*
Generic function to make Get request to endpoint
*/
func (api *Api) get(method string, option map[string]string) ([]byte, error) {
	q := url.Values{}
	for k, v := range option {
		q.Add(k, v)
	}

	fmt.Println(api.buildUrl(method, option))
	response, err := http.Get(api.buildUrl(method, option))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

/*
Generic function to make Post request to endpoint
*/
func (api *Api) post(method string, form map[string]string) ([]byte, error) {
	q := url.Values{}
	for k, v := range form {
		q.Add(k, v)
	}

	response, err := http.PostForm(api.buildUrl(method, nil), q)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

/*
Composer API
*/
func (api *Api) Get() {

}

func (api *Api) Send() {

}

func (api *Api) Forward() {

}

func (api *Api) Set() {

}
