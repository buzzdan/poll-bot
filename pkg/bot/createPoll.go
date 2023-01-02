package bot

import (
	"fmt"
	"github.com/buzzdan/poll-bot/pkg/types"
	"strings"
)

var (
	polls []types.Poll
)

type CreatePollCommand struct {
	types.Command
	Options []string
}

type CreatePollResult struct {
	OriginalCommand CreatePollCommand
	PollID          int
}

func (cp CreatePollResult) GetID() int {
	return cp.PollID
}

func (cp CreatePollResult) GetCreatedMessage() string {
	return fmt.Sprintf("hey @%s, poll #%d created", cp.OriginalCommand.Commander.Name, cp.PollID)
}

func (bot *Bot) parseCreatePoll(command string) (*CreatePollCommand, error) {
	optionsSplit := strings.Split(command, ":")
	if len(optionsSplit) < 1 {
		return nil, fmt.Errorf("create poll command should have options")
	}
	options := strings.Split(optionsSplit[1], ",")

	trimmed := []string{}
	for _, opt := range options {
		trimmed = append(trimmed, strings.TrimSpace(opt))
	}
	return &CreatePollCommand{
		Command: types.Command{Commander: bot.Commander, CommandType: types.CommandCreatePoll},
		Options: trimmed,
	}, nil
}

func (cp CreatePollCommand) Execute() (types.IDer, error) {
	options := map[int]string{}
	for i, opt := range cp.Options {
		options[i] = opt
	}

	newPoll := types.Poll{
		ID:        len(polls) + 1, //in DB we'd use UUID
		Options:   options,
		CreatedBy: cp.Commander,
	}

	polls = append(polls, newPoll)
	return &CreatePollResult{OriginalCommand: cp, PollID: newPoll.ID}, nil
}
