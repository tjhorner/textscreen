package main

import "fmt"

type MemoryStorageAdapter struct {
	Conversations map[string]Conversation
}

func (c *MemoryStorageAdapter) Delete(conversation *Conversation) error {
	delete(c.Conversations, conversation.PhoneNumber)
	return nil
}

func (c *MemoryStorageAdapter) Save(conversation *Conversation) error {
	c.Conversations[conversation.PhoneNumber] = *conversation
	return nil
}

func (c *MemoryStorageAdapter) Get(phoneNumber string) (*Conversation, error) {
	conversation, found := c.Conversations[phoneNumber]
	if !found {
		return nil, fmt.Errorf("conversation was not found")
	}

	return &conversation, nil
}

func (c *MemoryStorageAdapter) Exists(phoneNumber string) (bool, error) {
	_, found := c.Conversations[phoneNumber]
	return found, nil
}

func NewMemoryStorageAdapter() *MemoryStorageAdapter {
	return &MemoryStorageAdapter{
		Conversations: map[string]Conversation{},
	}
}
