package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func main() {
	var (
		dataPath string
		data     []byte
	)

	flag.StringVar(
		&dataPath, "m", "",
		fmt.Sprintf(
			"json file for data payload. Example: %s. File content example is: %s",
			`./fcm-send -m message.json`,
			`{"notification": { "title": "Hello", "body": "World"} }`,
		),
	)

	if len(os.Args) == 1 {
		fmt.Println("require arguments")
		fmt.Println()
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if dataPath != "" {
		d, err := os.ReadFile(dataPath)
		if err != nil {
			fmt.Printf("failed load data from %s\n", dataPath)
			os.Exit(1)
		}
		data = d
	}

	err := sendMessage(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sendMessage(data []byte) error {
	opt := option.WithCredentialsFile("service-account-file.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	client, err := app.Messaging(context.Background())
	if err != nil {
		return err
	}

	var message messaging.Message
	err = json.Unmarshal(data, &message)
	if err != nil {
		return err
	}

	_, err = client.Send(context.Background(), &message)
	if err != nil {
		return err
	}

	return nil
}
