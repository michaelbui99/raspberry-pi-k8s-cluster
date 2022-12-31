package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/michaelbui99/discord-alerthandler/internal/alertmanager"
	"github.com/michaelbui99/discord-alerthandler/internal/context"
	"github.com/michaelbui99/discord-alerthandler/internal/discord"
)

func onAlert(w http.ResponseWriter, r *http.Request) {
	context, err := context.ParseFromEnvironment()
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("Received new alert")
	defer r.Body.Close()

	bytes, _ := io.ReadAll(r.Body)

	var receivedAlert alertmanager.AlertManagerDTO
	json.Unmarshal(bytes, &receivedAlert)

	discordAlert := discord.BuildDiscordAlert(&receivedAlert)
	discord.SendDiscordAlert(*&context, discordAlert)
}

func main() {
	context, err := context.ParseFromEnvironment()
	if err != nil {
		log.Fatalf(err.Error())
	}

	http.HandleFunc(context.HandlerWebHookPath, onAlert)

	addr := fmt.Sprintf(":%s", context.HandlerPort)
	log.Println("Discord Alerthandler starting on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
