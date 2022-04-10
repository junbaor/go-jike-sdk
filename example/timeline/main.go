package main

import (
	"context"
	"flag"
	"fmt"
	"go-jike-sdk/jike"
	"log"
	"strings"
)

var (
	areaCode string
	phone    string
	password string
)

func init() {
	flag.StringVar(&areaCode, "areaCode", "+86", "areaCode for Jike.")
	flag.StringVar(&phone, "phone", "", "phone for Jike.")
	flag.StringVar(&password, "password", "", "password for Jike.")
	flag.Parse()
}

func main() {
	if phone == "" || password == "" {
		flag.PrintDefaults()
		return
	}

	content := context.Background()
	client := jike.NewJike(areaCode, phone)

	loginOutput, err := client.UserService.PasswordLogin(content, areaCode, phone, password)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Username: %s \nScreenName: %s\n", loginOutput.User.Username, loginOutput.User.ScreenName)

	timeline, err := client.UserService.FollowingTimeline(content, 10, jike.TimelineLoadMoreKey{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, item := range timeline.Data {
		println("⭐️ " + item.User.ScreenName + " " + strings.ReplaceAll(item.Content, "\n", "↲"))
	}
}
