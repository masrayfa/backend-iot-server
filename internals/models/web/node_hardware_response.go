package web

import "github.com/masrayfa/internals/models/domain"

type NodeByHardware struct {
	IdNode   int64  `json:"id_node"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type NodeByHardwareResponse struct {
	Node []NodeByHardware `json:"node"`
	Hardware domain.Hardware `json:"hardware"`
}