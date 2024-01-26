package web

import "github.com/masrayfa/internals/models/domain"

type NodeWithFeedResponse struct {
	Node domain.Node
	Feed []domain.Channel
}