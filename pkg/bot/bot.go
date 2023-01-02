package bot

import (
	"errors"
	"fmt"
	"github.com/buzzdan/poll-bot/pkg/types"
	"strings"
)

var (
	ErrInvalidMessage = errors.New("invalid message")
)

type Bot struct {
	Commander types.User
}

func StartBotFor(commander types.User) *Bot {
	return &Bot{Commander: commander}
}

func (bot *Bot) ExecuteCommand(message string) (string, error) {
	executor, err := bot.AnalyseMessage(message)
	if err != nil {
		return "", err
	}
	ider, err := executor.Execute()
	if err != nil {
		return "", err
	}
	return ider.GetCreatedMessage(), nil
}

func (bot *Bot) AnalyseMessage(message string) (types.Executor, error) {
	message = strings.TrimSpace(message)
	split := strings.Split(message, ",")

	if len(split) < 2 {
		return nil, ErrInvalidMessage
	}

	heyBot := split[0]
	if heyBot != "hey @bot" {
		return nil, fmt.Errorf("%w: you should start a command by saying - hey @bot", ErrInvalidMessage)
	}

	command := strings.TrimSpace(strings.TrimPrefix(message, heyBot+","))
	if strings.HasPrefix(command, "create a poll") {
		createPollCommand, err := bot.parseCreatePoll(command)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInvalidMessage, err.Error())
		}
		return createPollCommand, nil
	}

	println("got message: " + message)
	return nil, errors.New("not implemented")
}
