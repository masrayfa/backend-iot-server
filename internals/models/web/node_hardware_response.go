package web

type NodeByHardwareResponse struct {
	IdNode   int64  `json:"id_node"`
	Name     string `json:"name"`
	Location string `json:"location"`
}