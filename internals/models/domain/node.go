package domain

type Node struct {
	IdNode           int64    `json:"id_node"`
	Name             string   `json:"name"`
	Location         string   `json:"location"`
	FieldSensor      []string `json:"field_sensor"`
	IdHardwareSensor []int64  `json:"id_hardware"`
	IdHardwareNode   int64    `json:"id_hardware_node"`
	IdUser           int64    `json:"id_user"`
	IsPublic         bool     `json:"is_public"`
}

type NodeWithFeed struct {
	Node Node
	Feed []Channel
}

func convertBoolToInt(isPublic bool) int64 {
	if isPublic {
		return 1
	}
	return 0
}