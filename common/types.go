package common

import (
	"github.com/google/uuid"
)

type (
	NodeId        = uuid.UUID
	NodeContentId = uuid.UUID
	CommitId      = uuid.UUID
	DocumentId    = uuid.UUID
	BranchId      = uuid.UUID
	RevisionId    = uuid.UUID
	Color         = *string
)

type RootNodeJson struct {
	Filename   string `json:"filename"`
	RootNodeId NodeId `json:"rootNodeId"`
}

type NodeContentJson struct {
	Text            *string `json:"text"`
	BackgroundColor *string `json:"backgroundColor"`
}

type NodeContent struct {
	NodeContentId     NodeContentId
	Text              *string
	Note              *string
	BackgroundColor   *string
	CreatedTimestamp  *int64
	ModifiedTimestamp *int64
	FreemindId        *string
	Url               *string
}
