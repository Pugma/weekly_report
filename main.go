package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/robfig/cron/v3"
	"github.com/traPtitech/go-traq"
)

var (
	TraqAccessToken      = os.Getenv("BOT_ACCESS_TOKEN")
	embed           bool = true
)

func main() {
	c := cron.New()
	c.AddFunc("* * * * *", gitHub)
}

func gitHub() {
	activity, err := Do()
	if err != nil {
		slog.Error("%d", err)
	}

	if err = sendToTraq(activity); err != nil {
		slog.Error("%d", err)
	}
}

func sendToTraq(a string) error {
	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, TraqAccessToken)

	client.MessageApi.PostMessage(auth, "").PostMessageRequest(traq.PostMessageRequest{
		Content: "",
		Embed: &embed,
	}).Execute()

	return nil
}
