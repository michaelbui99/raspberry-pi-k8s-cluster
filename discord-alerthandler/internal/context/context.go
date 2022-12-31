package context

import (
	"errors"
	"os"
	"strings"
)

type Context struct {
	DiscordWebHookUrl  string
	HandlerWebHookPath string
	HandlerPort        string
}

func NewContext(discordWebHookUrl string, handlerWebHookPath string, handlerPort string) *Context {
	return &Context{DiscordWebHookUrl: discordWebHookUrl,
		HandlerWebHookPath: handlerWebHookPath,
		HandlerPort:        handlerPort,
	}
}

func ParseFromEnvironment() (*Context, error) {
	discordWebHookUrl := os.Getenv("DISCORD_WEBHOOK_URL")
	handlerWebHookPath := os.Getenv("HANDLER_WEBHOOK_PATH")
	handlerPort := os.Getenv("HANDLER_PORT")

	if strings.TrimSpace(discordWebHookUrl) == "" {
		return nil, errors.New("No discord webhook has been provided")
	}

	if strings.TrimSpace(handlerWebHookPath) == "" {
		handlerWebHookPath = "/"
	}

	if strings.TrimSpace(handlerPort) == "" {
		handlerPort = "8089"
	}

	return NewContext(discordWebHookUrl, handlerWebHookPath, handlerPort), nil
}
