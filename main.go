package main

import (
	"github.com/buzzdan/poll-bot/pkg/bot"
	"github.com/buzzdan/poll-bot/pkg/types"
	"os"
	"strings"
)

//message parser - split to words, discover tokens - normal word, command (create, show-results, vote) tag user, poll options

//objects:

//bot - able to take commands, login user
//command - commander (user), command-type, poll-options
//user - id, name
//poll
//result per command - action id, poll-id / vote-id, command

// use cases:
// A user asks: ‘hey @bot, create a poll with options: <option 1>, <option 2>, <option 3>’
// The bot replies: ‘hey @<username>, poll #1 created.’

func main() {
	user := os.Args[1]
	message := strings.Join(os.Args[2:], " ")
	b := bot.StartBotFor(types.User{
		ID:   "my-user-id",
		Name: user,
	})

	result, err := b.ExecuteCommand(message)

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(result)
}
