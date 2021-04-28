package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/eclipse/paho.golang/paho"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1883")

	if err != nil {
		fmt.Println("error initializing connection")
		return
	}

	mqttClient := paho.NewClient()

	mqttClient.Conn = conn
	mqttClient.Connect(context.Background(), &paho.Connect{})

	mqttClient.Publish(context.Background(), &paho.Publish{Topic: "smoker/1/temperature/1"})

	for {
		time.Sleep(time.Second)
		writeMessages(context.Background(), mqttClient)
	}
}

func writeMessages(ctx context.Context, mqtt *paho.Client) {
	for smoker := 1; smoker <= 2; smoker++ {
		for temperature := 1; temperature <= 8; temperature++ {
			topic := fmt.Sprintf("smoker/%v/temperature/%v", smoker, temperature)
			value := rand.Float32()*5 + float32(temperature) + float32(temperature%4*10) + float32(smoker*5)

			msg := []byte(fmt.Sprintf("{\"value\":%f}", value))

			mqtt.Publish(ctx, &paho.Publish{Topic: topic, Payload: msg})
		}
	}
}
