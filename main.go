package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/pelletier/go-toml"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/mqtt"
)

func main() {
	configFile := "virtsense.toml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	var config Config
	configTree, err := toml.LoadFile(configFile)
	if err != nil {
		log.Fatalln("Failed to read config", err)
	}

	err = configTree.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Failed to unmarshal config", err)
	}

	generator := Generators[config.Generator](config.Min, config.Max)

	adaptor := mqtt.NewAdaptorWithAuth(config.MQTTEndpoint, config.ClientID, config.Username, config.Password)

	work := func() {
		gobot.Every(15*time.Second, func() {
			dp := generator()
			msgID, err := GenerateMessageID(config.MessageIDFile)
			if err != nil {
				log.Println("Failed to generate message ID", err)
				return
			}

			payload := map[string]interface{}{
				"msg_id":    msgID,
				"unit_id":   config.Unit,
				"sensor":    config.Sensor,
				"value":     strconv.FormatFloat(dp, 'f', 2, 64),
				"timestamp": time.Now().Unix(),
			}

			rawPayload, err := json.Marshal(&payload)
			if err != nil {
				log.Println("Failed to marshal payload", err)
				return
			}

			adaptor.Publish(config.Topic, rawPayload)
			log.Println("Sending reading", dp)
		})
	}

	robot := gobot.NewRobot("virtsense",
		[]gobot.Connection{adaptor},
		work,
	)

	robot.Start()
}
