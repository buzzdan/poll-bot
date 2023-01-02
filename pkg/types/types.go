package types

type User struct {
	ID, Name string
}

type Poll struct {
	ID        int
	Message   string
	Options   map[int]string
	CreatedBy User
}

type Command struct {
	Commander   User
	CommandType CommandType
}

type Executor interface {
	Execute() (IDer, error)
}

type IDer interface {
	GetID() int
	GetCreatedMessage() string
}

type CommandType int

const (
	CommandCreatePoll = iota
	CommandVote
	CommandShow
)
