package telegram

import (
	//	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetUpdate(t *testing.T) {
	var query, path string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query = r.URL.Query().Encode()
		path = r.URL.Path
		fmt.Fprint(w, `{"ok":true,"result":[{"update_id":141616126,
		"message":{"message_id":8,"from":{"id":57905015,"first_name":"Kurei","last_name":"Kain","username":"kurei"},"chat":{"id":-28394494,"title":"Moonvui"},"date":1435953204,"text":"@moonvuibot be dua"}},{"update_id":141616127,
		"message":{"message_id":9,"from":{"id":57905015,"first_name":"Kurei","last_name":"Kain","username":"kurei"},"chat":{"id":-28394494,"title":"Moonvui"},"date":1435957006,"text":"@moonvuibot be hu dau nanh"}}]}`)
	}))
	defer ts.Close()

	config := NewConfig("abcxyz", func(c *Config) {
		c.WithUrl(ts.URL)
	})

	api, _ := NewApi(config)
	q := make(map[string]string)
	q["foo"] = "bar"
	q["foo2"] = "bar2"
	update, err := api.GetUpdates(q)

	assert.Equal(t, nil, err, "should have no errors")
	assert.Equal(t, "@moonvuibot be dua", update[0].Message.Text, "should parse result properly")
	assert.Equal(t, "foo=bar&foo2=bar2", query, "should pass corret params")
	assert.Equal(t, "/abcxyz/getUpdates", path, "should request correct endpoint")
}

func TestSendMessage(t *testing.T) {
	var query, path string
	var form url.Values
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query = r.URL.Query().Encode()
		path = r.URL.Path

		err := r.ParseForm()
		form = r.Form
		if err != nil {
			fmt.Fprint(w, "Cannot parse form")
		}

		fmt.Fprint(w, `{"ok":true,"result":{"message_id":49,"from":{"id":1,"first_name":"Moonvui","username":"moonvuibot"},"chat":{"id":-28394494,"title":"Moonvui"},"date":1436047996,"text":"Moonvuibot finds you some cake  shop near san jose"}}`)
	}))
	defer ts.Close()

	config := NewConfig("abcxyz", func(c *Config) {
		c.WithUrl(ts.URL)
	})

	api, _ := NewApi(config)
	m := make(map[string]string)
	m["chat_id"] = "1"
	m["text"] = fmt.Sprintf("Moonvuibot finds you some %s  shop near %s", "coffee", "san jose")
	message, err := api.SendMessage(m)

	assert.Equal(t, nil, err, "should have no errors")
	assert.Equal(t, "Moonvuibot finds you some cake  shop near san jose", message.Text, "should parse result properly")
	assert.Equal(t, "chat_id=1&text=Moonvuibot+finds+you+some+coffee++shop+near+san+jose", form.Encode(), "should pass corret params")
	assert.Equal(t, "/abcxyz/sendMessage", path, "should request correct endpoint")
}
