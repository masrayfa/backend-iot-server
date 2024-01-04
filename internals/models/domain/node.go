package domain

type Node struct {
	IdNode     int    `json:"id_node"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	IdHardware int    `json:"id_hardware"`
	IdUser     int64  `json:"id_user"`
}