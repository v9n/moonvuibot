package main

import (
	"fmt"
	"github.com/JustinBeckwith/go-yelp/yelp"
	"github.com/kureikain/moonvuibot/telegram"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {

	config := telegram.NewConfig(os.Getenv("TELEGRAM_TOKEN"))
	api, _ := telegram.NewApi(config)

	done := make(chan int)
	go func() {
		listen(api)
	}()

	user, _ := api.GetMe()
	fmt.Println("RRR = ", user)

	<-done
}

func listen(api *telegram.Api) {
	fmt.Println("Listen")
	//This code run async
	opt := make(map[string]string)
	opt["timeout"] = "120"

	var maxOffsetId int64
	maxOffsetId = 0
	for {
		fmt.Println("Pool")
		updates, err := api.GetUpdates(opt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(len(updates))
		for _, element := range updates {
			fmt.Printf("UpdateID = %s", element.UpdateId)
			fmt.Println(element.Message.Text)

			re := regexp.MustCompile("find (.+) in (.+)$")
			match := re.FindAllStringSubmatch(element.Message.Text, -1)
			if match != nil && len(match) >= 1 {
				findYelp(api, match[0][1], match[0][2])
				fmt.Println("FINDDD")
				continue
			}

			re = regexp.MustCompile("xin ko\\?")
			match = re.FindAllStringSubmatch(element.Message.Text, -1)
			if match != nil && len(match) >= 1 {
				api.SendMessage(map[string]string{
					"chat_id": "-28394494",
					"text":    "YES, I AM!!!",
				})
				continue
			}

			api.SendMessage(map[string]string{
				"chat_id": "-28394494",
				"text":    "DONT KNOW WHAT YOU ARE TALKING ABOUT. HERE IS A GOOGLE RESULT: ",
			})

			if element.UpdateId > maxOffsetId {
				maxOffsetId = element.UpdateId
			}
		}
		opt["offset"] = strconv.FormatInt(maxOffsetId+1, 10)
		fmt.Println("sasa= ", opt["offset"])
		time.Sleep(1000 * time.Millisecond)
	}
}

func findYelp(api *telegram.Api, term string, location string) {
	m := make(map[string]string)
	m["chat_id"] = "-28394494"
	m["text"] = fmt.Sprintf("Moonvuibot finds you some %s  shop near %s", term, location)
	message, _ := api.SendMessage(m)
	fmt.Println("MMM = ", message)

	o := &yelp.AuthOptions{
		ConsumerKey:       os.Getenv("YELP_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("YELP_CONSUMER_SECRET"),
		AccessToken:       os.Getenv("YELP_TOKEN"),
		AccessTokenSecret: os.Getenv("YELP_TOKEN_SECRET"),
	}
	yelpClient := yelp.New(o, nil)
	results, err := yelpClient.DoSimpleSearch(term, location)
	if err != nil {
		fmt.Println(err)
	}
	m["text"] = fmt.Sprintf("Found a total of %v results for \"%v\" in \"%v\". \n===\n", results.Total, term, location)
	for _, business := range results.Businesses {
		m["text"] = m["text"] + fmt.Sprintf("\n %v, rate: %v\n\n", business.Name, business.Rating)
	}
	api.SendMessage(m)

}
