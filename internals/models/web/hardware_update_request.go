package web

type HardwareUpdateRequest struct {
	IdHardware  int64  `json:"id_hardware"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}