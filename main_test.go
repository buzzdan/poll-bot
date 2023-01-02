package main

import (
	bot "github.com/buzzdan/poll-bot/pkg/bot"
	"github.com/buzzdan/poll-bot/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePollCommand(t *testing.T) {
	b := bot.StartBotFor(types.User{
		ID:   "my-user-id",
		Name: "dandan",
	})
	result, err := b.ExecuteCommand("hey @bot, create a poll with options: aroma, coffee bean, cafe azaria")

	assert.NoError(t, err)
	assert.Equal(t, "hey @dandan, poll #1 created", result)
}
