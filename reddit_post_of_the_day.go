package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
	"strings"
	"time"
)
// graw (Go Reddit API Wrapper) is a tool for gathering data from Reddit and creating Reddit bots.

type reminderBot struct {
	bot reddit.Bot
}

func (r *reminderBot) Post(p *reddit.Post) error {
	if strings.Contains(p.SelfText, "remind me of this post") {
		<-time.After(10 * time.Second)
		return r.bot.SendMessage(
			p.Author,
			fmt.Sprintf("Reminder: %s", p.Title),
			"You've been reminded!",
		)
	}
	return nil
}

func main() {
	if bot, err := reddit.NewBotFromAgentFile("reminderbot.agent", 0); err != nil {
		glog.Infof("Failed to create bot handle: ", err)
	} else {
		cfg := graw.Config{Subreddits: []string{"bottesting"}}
		handler := &reminderBot{bot: bot}
		if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
			glog.Infof("Failed to start graw run: ", err)
		} else {
			glog.Infof("graw run failed: ", wait())
		}
	}

	bot, err := reddit.NewBotFromAgentFile("reminderbot.agent", 0)
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}

	harvest, err := bot.Listing("/r/golang", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/golang: ", err)
		return
	}

	for _, post := range harvest.Posts[:5] {
		fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	}
}