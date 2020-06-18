package main

import (
	"fmt"
	"github.com/golang/glog"
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
	cfg := reddit.BotConfig{
		Agent: "graw:doc_demo_bot:0.3.1 by /u/skp1998",
		App: reddit.App{
			ID:     "0VpLcglAwwG_cg",
			Secret: "FQTn77zAMeMN67Bi8h4ooFVqs18",
			Username: "skp1998",
			Password: "reddit@123",
		},
	}
	bot, _ := reddit.NewBot(cfg)

	//list top post from golang group and send message
	harvest, err := bot.Listing("/r/golang", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/golang: ", err)
		return
	}

	//construct a max heap and fetch top five posts
	var topFivePosts []*reddit.Post
	h := BuildMaxHeap(harvest.Posts)
	topFivePosts=append(topFivePosts,h.post[0])
	glog.Infof("user:%v, Title:%v, Ups: %v",h.post[0].Author,h.post[0].Title,h.post[0].Ups)
	for i := 1; i <5; i++ {
		h = BuildMaxHeap(h.post[i:])
		glog.Infof("Id:%v, Title:%v, Ups: %v",h.post[0].Author,h.post[0].Title,h.post[0].Ups)
	}

	//for _, post := range harvest.Posts  {
	//
	//	fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	//}

	//send message to a user inbox on reddit
	//err = bot.SendMessage("skp1998", "Thanks for making this Reddit API! Mr. Shani", "It's ok.")
	//if err != nil {
	//	glog.Infof("error while sending message to user skp1998 %v",err)
	//}
}
