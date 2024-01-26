package web

type NodeCreateRequest struct {
	Name             string   `json:"name"`
	Location         string   `json:"location"`
	FieldSensor      []string `json:"field_sensor"`
	IdHardwareNode   int64    `json:"id_hardware_node"`
	IdHardwareSensor []int64  `json:"id_hardware_sensor"`
	IsPublic         bool     `json:"is_public"`
}