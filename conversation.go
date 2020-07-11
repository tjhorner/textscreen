package main

// ConversationStorageAdapter is an interface that storage adapters
// should conform to in order to provide backing storage for
// conversation sessions.
type ConversationStorageAdapter interface {
	Save(conversation *Conversation) error
	Delete(conversation *Conversation) error
	Get(phoneNumber string) (*Conversation, error)
	Exists(phoneNumber string) (bool, error)
}

type ConversationState int

const (
	Started ConversationState = iota
	AskedForName
	AskedForPurpose
	Complete
)

type Conversation struct {
	PhoneNumber string
	State       ConversationState
	Name        string
	Purpose     string
}

func NewConversation(phoneNumber string) *Conversation {
	return &Conversation{
		PhoneNumber: phoneNumber,
		State:       Started,
	}
}
