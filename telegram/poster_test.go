package telegram

import (
	//	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUpdateSendCorrectParam(t *testing.T) {
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

func TestSendMessageSendCorrectParam(t *testing.T) {

}
