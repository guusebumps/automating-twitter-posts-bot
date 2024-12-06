package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	_, _, err := client.Statuses.Update("Hello, Twitter!", nil)
	if err != nil {
		fmt.Println("Error posting tweet:", err)
		return
	}
	fmt.Println("Tweet posted successfully")
}
