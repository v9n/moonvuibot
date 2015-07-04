# Moonvui bot

A telegram API client implentationa and a bot that use it


# API 


## Long polling async style
```
config := telegram.NewConfig('abcxyz')
api :=    telegram.NewApi(config)

//This code run async
api.getUpdates(func ([]*telegram.Message) {
})

//The program flow normally here
```

## Send Message

We work around map, compose a map of form value/query param and pass
into corresponding method

```
	m := make(map[string]string)
	m["chat_id"] = "chat_id_number"
	m["text"] = "test test"
	message, _ := api.SendMessage(m)
	fmt.Println("MMM = ", message)
```
