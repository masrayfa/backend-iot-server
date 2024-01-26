package web

import (
	"github.com/masrayfa/internals/models/domain"
)

type NodeUpdateRequest struct {
	Name             string   `json:"name"`
	Location         string   `json:"location"`
	IdHardwareNode   int64    `json:"id_hardware_node"`
	IdHardwareSensor []int64  `json:"id_hardware_sensor"`
	FieldSensor      []string `json:"field_sensor"`
	IsPublic         int64 `json:"is_public"`
}

func (n *NodeUpdateRequest) ChangeSettedField(node *domain.Node) {
	if n.Name == "" {
		n.Name = node.Name
	}

	if n.Location == "" {
		n.Location = node.Location
	}

	if n.IdHardwareNode == 0 {
		n.IdHardwareNode = node.IdHardwareNode
	}

	if n.IdHardwareSensor == nil || len(n.IdHardwareSensor) == 0 {
		n.IdHardwareSensor = node.IdHardwareSensor
	}

	if n.FieldSensor == nil || len(n.FieldSensor) == 0 {
		n.FieldSensor = node.FieldSensor
	}

	if n.IsPublic == 0 {
		n.IsPublic = convertBoolToInt(node.IsPublic)
	}
}

func convertBoolToInt(isPublic bool) int64 {
	if isPublic {
		return 1
	}
	return 0
}