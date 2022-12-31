package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/michaelbui99/discord-alerthandler/internal/alertmanager"
	"github.com/michaelbui99/discord-alerthandler/internal/context"
)

type DiscordAlert struct {
	Content   string         `json:"content"`
	Username  string         `json:"username"`
	AvatarUrl string         `json:"avatar_url"`
	Embeds    []DiscordEmbed `json:"embeds"`
}

type DiscordEmbed struct {
	Type        string             `json:"type"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Color       int                `json:"color"`
	Url         string             `json:"url"`
	Author      DiscordEmbedAuthor `json:"author"`
}

type DiscordEmbedAuthor struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func NewDiscordEmbed(
	embedType string,
	title string,
	description string,
	color int,
	url string,
	author DiscordEmbedAuthor) *DiscordEmbed {

	return &DiscordEmbed{
		Type:        embedType,
		Title:       title,
		Description: description,
		Color:       color,
		Url:         url,
		Author:      author,
	}
}

func NewDiscordAlert(
	content string,
	username string,
	avatarUrl string,
	embeds []DiscordEmbed) *DiscordAlert {

	return &DiscordAlert{
		Content:   content,
		Username:  username,
		AvatarUrl: avatarUrl,
		Embeds:    embeds,
	}
}

func NewDiscordEmbedAuthor(name string, url string) *DiscordEmbedAuthor {
	return &DiscordEmbedAuthor{Name: name, Url: url}
}

func SendDiscordAlert(context *context.Context, alert *DiscordAlert) {
	alertJson, _ := json.Marshal(*alert)
	http.Post(*&context.DiscordWebHookUrl, "application/json", bytes.NewReader(alertJson))
}

func BuildDiscordAlert(alert *alertmanager.AlertManagerAlert) *DiscordAlert {
	embedAuthor := NewDiscordEmbedAuthor(fmt.Sprintf("Discord Alerthandler - %v - %v", alert.Labels["instance"], alert.Labels["job"]), "")
	embed := NewDiscordEmbed("rich", fmt.Sprintf("[%v] %v", alert.Status, alert.Labels["alertname"]), alert.Annotations.Description, 0x00FFFF, alert.GeneratorURL, *embedAuthor)
	embeds := []DiscordEmbed{*embed}
	return &DiscordAlert{
		Content: fmt.Sprintf("A new alert has been triggered: [%v] - [%v].", alert.StartsAt, alert.EndsAt),
		Embeds:  embeds,
	}
}
