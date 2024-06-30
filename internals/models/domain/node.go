package domain

type Node struct {
	IdNode           int64    `json:"id_node"`
	Name             string   `json:"name"`
	Location         string   `json:"location"`
	FieldSensor      []string `json:"field_sensor"`
	IdHardwareSensor []int64  `json:"id_hardware_sensor"`
	IdHardwareNode   int64    `json:"id_hardware_node"`
	IdUser           int64    `json:"id_user"`
	IsPublic         bool     `json:"is_public"`
}

type NodeWithFeed struct {
	Node Node
	Feed []Channel
}

type NodeWithFeedRow struct {
	Node Node
	Feed []ChannelRow
}

type FieldSensor struct {
	IdFieldSensor int64  `json:"id_field_sensor"`
	Name          string `json:"name"`
}

type NodeDetail struct {
	IdNode           int64    `json:"id_node"`
	Name             string   `json:"name"`
	Location         string   `json:"location"`
	IdHardwareNode   int64    `json:"id_hardware"`
	FieldSensor      []string `json:"field_sensor"`
	IdHardwareSensor []int64  `json:"id_hardware_sensor"`
	Hardware         Hardware `json:"hardware"`
}

type NodeWithFeedV2 struct {
	IdNode     int64    `json:"id_node"`
	Name       string   `json:"name"`
	Location   string   `json:"location"`
	IdHardware int64    `json:"id_hardware"`
	IdUser     int64    `json:"id_user"`
	IsPublic   bool     `json:"is_public"`
	Hardware   Hardware `json:"hardware"`
	Sensor     []Sensor `json:"sensor"`
}

func convertBoolToInt(isPublic bool) int64 {
	if isPublic {
		return 1
	}
	return 0
}
