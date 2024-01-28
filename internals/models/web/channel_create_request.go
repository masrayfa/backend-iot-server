package web

type ChannelCreateRequest struct {
	IdNode int64  `json:"id_node"`
	Value  string `json:"value"`
}