package web

// NodeReadResponse is a struct that represent the response of node read
// it's a temporary struct, it will be removed when the node read response is implemented
// because the node read response will be the same as node create response
type NodeCreateResponse struct {
	Name             string   `json:"name"`
	Location         string   `json:"location"`
	FieldSensor      []string `json:"field_sensor"`
	IdHardwareNode   int64    `json:"id_hardware_node"`
	IdHardwareSensor []int64  `json:"id_hardware_sensor"`
	IsPublic         bool     `json:"is_public"`
}