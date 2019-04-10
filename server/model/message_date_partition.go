package model

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

// MessageDatePartitionModel is the resolvable struct for the MessageDatePartition struct
type MessageDatePartitionModel struct {
	*MessageDatePartition
}

// MessageDatePartition type
type MessageDatePartition struct {
	CircleID      string
	date          time.Time
	MessageGroups []*MessageGroupModel
}

// CircleID field resolver
func (mdp *MessageDatePartitionModel) CircleID() graphql.ID {
	return graphql.ID(mdp.MessageDatePartition.CircleID)
}

// MessageGroups field resolver
func (mdp *MessageDatePartitionModel) MessageGroups() []*MessageGroupModel {
	return mdp.MessageDatePartition.MessageGroups
}

// Date field resolver
func (mdp *MessageDatePartitionModel) Date() graphql.Time {
	return graphql.Time{Time: mdp.MessageDatePartition.date}
}
