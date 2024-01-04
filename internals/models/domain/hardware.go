package domain

type Hardware struct {
	IdHardware  int    `json:"id_hardware"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}