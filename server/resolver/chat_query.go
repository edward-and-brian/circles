package resolver

import (
	"context"

	"circles/server/types"
)

// Chat retrieves the chat specified by id
func (r *Resolver) Chat(ctx context.Context, args *IDArgs) (*ChatResolver, error) {
	var chat *types.Chat

	if err := r.Store.Open(); err != nil {
		return nil, err

	} else if chat, err = r.Store.FindChat(&args.ID); err != nil {
		return nil, err
	}

	return &ChatResolver{chat, r.Store}, nil
}

// UserChats retrieves all chats for a user
func (r *Resolver) UserChats(ctx context.Context, args *IDArgs) ([]*ChatResolver, error) {
	var (
		chats         []*types.Chat
		chatsResolver []*ChatResolver
	)

	if err := r.Store.Open(); err != nil {
		return nil, err

	} else if chats, err = r.Store.AllUserChats(&args.ID); err != nil {
		return nil, err
	}

	for _, ch := range chats {
		chatsResolver = append(chatsResolver, &ChatResolver{ch, r.Store})
	}

	return chatsResolver, nil
}
