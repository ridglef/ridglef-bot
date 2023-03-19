package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/5elenay/revoltgo"
)

func main() {
	client := revoltgo.Client{
		Token: "",
	}
	client.OnMessage(func(msg *revoltgo.Message) {
		if msg.AuthorId == "ehvsSkIKY5GF59fk8Vuf_JCrIUEAK8Il645LQDrUh_" {
			return
		} else {
			var messagestring string = msg.Content.(string)
			if (strings.Contains(strings.ToLower(messagestring), "what would you do")) || (strings.Contains(strings.ToLower(messagestring), "what would you use")) {
				sendMsg := &revoltgo.SendMessage{}
				sendMsg.SetContent("HASHMAP")
				msg.Reply(true, sendMsg)
				sendMsgTwo := &revoltgo.SendMessage{}
				sendMsgTwo.SetContent("I would use a hasmap.")
				msg.Reply(true, sendMsgTwo)
			}
			if strings.Contains(strings.ToLower(messagestring), "seedcrackerx") {
				sendMsg := &revoltgo.SendMessage{}
				sendMsg.SetContent("https://media.discordapp.net/attachments/1081690345995247646/1081986210966290573/893615002015-1-316249021.jpg")
				msg.Reply(true, sendMsg)
			}
			if messagestring == "!one" {
				sendMsg := &revoltgo.SendMessage{}
				sendMsg.SetContent("https://www.youtube.com/watch?v=M9J6DKJXoKk")
				msg.Reply(true, sendMsg)
			}
			if messagestring == "!help" {
				sendMsg := &revoltgo.SendMessage{}
				sendMsg.SetContent("Commands: !help, !cat, !one")
				msg.Reply(true, sendMsg)
			}
			if messagestring == "!cat" {
				resp, err := http.Get("https://aws.random.cat/meow")
				if err != nil {
					fmt.Println("Error getting cat image:", err)
					return
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading body:", err)
					return
				}

				var data map[string]interface{}
				err = json.Unmarshal(body, &data)
				if err != nil {
					fmt.Println("Error parsing JSON:", err)
					return
				}

				fileURL, ok := data["file"].(string)
				if !ok {
					fmt.Println("Error getting URL")
					return
				}

				sendMsg := &revoltgo.SendMessage{}
				sendMsg.SetContent(fileURL)
				msg.Reply(true, sendMsg)
			}
		}
	})

	client.Start()
	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)
	<-sc
	client.Destroy()

}
