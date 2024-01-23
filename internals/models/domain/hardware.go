package domain

type Hardware struct {
	IdHardware  int64  `json:"id_hardware"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type HardwareWithSensor struct {
	Hardware Hardware
	Sensors  []Sensor
}

type HardwareWithNode struct {
	Hardware Hardware
	Nodes    []Node
}