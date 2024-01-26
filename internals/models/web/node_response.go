package web

import "github.com/masrayfa/internals/models/domain"

type NodeResponse struct {
	Node domain.Node `json:"node"`
	Err error `json:"error"`
}