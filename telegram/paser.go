package telegram

import (
	"encoding/json"
	"fmt"
)

func parseResultData(jsonData []byte, to interface{}) interface{} {
	fmt.Println(string(jsonData))
	switch to.(type) {
	case *User:
		var r struct {
			Ok     bool
			Result *User
		}
		err := json.Unmarshal(jsonData, &r)

		if err == nil {
			return r.Result
		}
	case *Message:
		var r struct {
			Ok     bool
			Result *Message
		}
		err := json.Unmarshal(jsonData, &r)
		fmt.Println("result = ", r)

		if err == nil {
			return r.Result
		}

	}

	return nil
}
