package model

import (
	"context"
	"time"

	"circles/server/types"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rs/xid"
)

// AllMessages retrieves all messages and returns them as MessageModels
func AllMessages(ctx context.Context, gs generalStore) ([]*MessageModel, error) {
	messages, err := gs.AllMessages(ctx)
	if err != nil {
		return nil, err
	}

	var messageModels []*MessageModel
	for _, message := range messages {
		messageModels = append(messageModels, &MessageModel{message, gs})
	}

	return messageModels, nil
}

// CreateMessage creates a new Message with the given data
func CreateMessage(ctx context.Context, gs generalStore, message *types.Message) (*MessageModel, error) {
	message.ID = xid.New().String()

	if err := gs.CreateMessage(ctx, message); err != nil {
		return nil, err

	} else if message, err = gs.FindMessage(ctx, message.ID); err != nil {
		return nil, err
	}

	return &MessageModel{message, gs}, nil
}

// DeleteMessage deletes the Message specified by ID and returns it as a MessageModel
func DeleteMessage(ctx context.Context, gs generalStore, id string) (*MessageModel, error) {
	message, err := gs.FindMessage(ctx, id)
	if err != nil {
		return nil, err

	} else if err = gs.DeleteMessage(ctx, id); err != nil {
		return nil, err
	}

	return &MessageModel{message, gs}, nil
}

// FindMessage retrieves the Message specified by id
func FindMessage(ctx context.Context, gs generalStore, id string) (*MessageModel, error) {
	message, err := gs.FindMessage(ctx, id)
	if err != nil {
		return nil, err
	}

	return &MessageModel{message, gs}, nil
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

// CircleMessageDatePartitions retrieves all messages, puts them in small groups, and partitions them by date
func CircleMessageDatePartitions(ctx context.Context, gs generalStore, ciid string) ([]*MessageDatePartitionModel, error) {
	var (
		messageGroup         *MessageGroup
		messageDatePartition *MessageDatePartition

		messageDatePartitionModels = []*MessageDatePartitionModel{}
		messageGroupModels         = []*MessageGroupModel{}
		messageModels              = []*MessageModel{}
	)

	messages, err := gs.AllMessagesByCircleID(ctx, ciid)
	if err != nil {
		return nil, err
	}

	for _, m := range messages {
		messageCreatedAt, err := time.Parse("ANSCI", m.CreatedAt)
		if err != nil {
			return nil, err
		}

		if messageDatePartition == nil {
			messageDatePartition = &MessageDatePartition{m.CircleID, messageCreatedAt, nil}

		} else if !(messageCreatedAt.YearDay() == messageDatePartition.date.YearDay() && messageCreatedAt.Year() == messageDatePartition.date.Year()) {
			messageDatePartition.MessageGroups = messageGroupModels
			messageDatePartitionModels = append(messageDatePartitionModels, &MessageDatePartitionModel{messageDatePartition})
			messageDatePartition = &MessageDatePartition{m.CircleID, messageCreatedAt, nil}
			messageGroupModels = []*MessageGroupModel{}
		}

		if messageGroup == nil {
			messageGroup = &MessageGroup{m.SenderID, messageCreatedAt, messageCreatedAt, nil}

		} else if messageCreatedAt.Sub(messageGroup.LastMessageAt).Minutes() <= 10 {
			messageGroup.LastMessageAt = messageCreatedAt

		} else {
			messageGroup.Messages = messageModels
			messageGroupModels = append(messageGroupModels, &MessageGroupModel{messageGroup})
			messageGroup = &MessageGroup{m.SenderID, messageCreatedAt, messageCreatedAt, nil}
		}

		messageModels = append(messageModels, &MessageModel{m, gs})
	}

	messageGroup.Messages = messageModels
	messageGroupModels = append(messageGroupModels, &MessageGroupModel{messageGroup})
	messageDatePartition.MessageGroups = messageGroupModels
	messageDatePartitionModels = append(messageDatePartitionModels, &MessageDatePartitionModel{messageDatePartition})

	return messageDatePartitionModels, nil
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
