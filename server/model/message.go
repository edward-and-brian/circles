package model

import (
	"context"
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
)

// CreateMessage creates a new Message with the given data and returns it as a MessageModel
func CreateMessage(ctx context.Context, gs generalStore, message *types.Message) (*MessageModel, error) {
	if err := gs.CreateMessage(ctx, message); err != nil {
		return nil, err

	} else if message, err = gs.FindMessage(ctx, message.ID); err != nil {
		return nil, err
	}

	return &MessageModel{message, gs}, nil
}

// DeleteMessage deletes the Message specified by ID and returns it as a MessageModel
func DeleteMessage(ctx context.Context, gs generalStore, id string) (*MessageModel, error) {
	Message, err := gs.FindMessage(ctx, id)
	if err != nil {
		return nil, err

	} else if err = gs.DeleteMessage(ctx, id); err != nil {
		return nil, err
	}

	return &MessageModel{Message, gs}, nil
}

// FindMessage retrieves the Message specified by id
func FindMessage(ctx context.Context, gs generalStore, id string) (*MessageModel, error) {
	circle, err := gs.FindMessage(ctx, id)
	if err != nil {
		return nil, err
	}

	return &MessageModel{circle, gs}, nil
}

// UpdateMessageInput ...
type UpdateMessageInput struct {
	ID      string
	Content *string
}

// UpdateMessage updates the message specified by ID with the given data and returns it as a MessageModel
func UpdateMessage(ctx context.Context, gs generalStore, input *UpdateMessageInput) (*MessageModel, error) {
	message, err := gs.FindMessage(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if input.Content != nil {
		message.Content = *input.Content
	}

	if err = gs.UpdateMessage(ctx, message); err != nil {
		return nil, err
	}

	return &MessageModel{message, gs}, nil
}

// CircleMessages retrieves all Messages for a Message
func CircleMessages(ctx context.Context, gs generalStore, ciid string) ([]*MessageModel, error) {
	circles, err := gs.AllMessagesByCircleID(ctx, ciid)
	if err != nil {
		return nil, err
	}

	var messageModels []*MessageModel
	for _, ch := range circles {
		messageModels = append(messageModels, &MessageModel{ch, gs})
	}

	return messageModels, nil
}

// MessageModel is the resolvable struct for the Message struct
type MessageModel struct {
	*types.Message
	store generalStore
}

// ID field resolver
func (m *MessageModel) ID() graphql.ID {
	return graphql.ID(m.Message.ID)
}

// CircleID field resolver
func (m *MessageModel) CircleID() graphql.ID {
	return graphql.ID(m.Message.CircleID)
}

// SenderID field resolver
func (m *MessageModel) SenderID() graphql.ID {
	return graphql.ID(m.Message.SenderID)
}

// Content field resolver
func (m *MessageModel) Content() string {
	return m.Message.Content
}

// CreatedAt field resolver
func (m *MessageModel) CreatedAt() (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, m.Message.CreatedAt)
	return graphql.Time{Time: t}, err
}
