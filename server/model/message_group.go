package model

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// MessageGroupModel is the resolvable struct for the MessageGroup struct
type MessageGroupModel struct {
	*MessageGroup
}

// MessageGroup type
type MessageGroup struct {
	SenderID      string
	LastMessageAt time.Time
	CreatedAt     time.Time
	Messages      []*MessageModel
}

// SenderID field resolver
func (mg *MessageGroupModel) SenderID() graphql.ID {
	return graphql.ID(mg.MessageGroup.SenderID)
}

// Messages field resolver
func (mg *MessageGroupModel) Messages() []*MessageModel {
	return mg.MessageGroup.Messages
}

// CreatedAt field resolver
func (mg *MessageGroupModel) CreatedAt() graphql.Time {
	return graphql.Time{Time: mg.MessageGroup.CreatedAt}
}
