package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func PrintCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	bot := slacker.NewClient(os.Getenv("SLACK_APP_TOKEN"), os.Getenv("SLACK_BOT_TOKEN"))

	go PrintCommandEvents(bot.CommandEvents())

	bot.Command("", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}
