# Moonvui bot

A telegram API client implentationa and a bot that uses it

[![wercker
status](https://app.wercker.com/status/0c494370686732a327599fb1b58b3e52/m
"wercker
status")](https://app.wercker.com/project/bykey/0c494370686732a327599fb1b58b3e52)

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

## Composer& Transport API (Pending)

This API allows us to create an object of type of Telegram then we call
`Transport` method

  * send
  * forward
  * set
  * get


Ideally, we do this:

```
  var m := composer.NewMessage()
  api.Transport.Send(m)

  var photo := composer.NewPhoto()
  api.Transport.Send(photo)
```
