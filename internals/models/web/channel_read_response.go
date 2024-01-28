package web

import "time"

type ChannelReadResponse struct {
	Time   time.Time `json:"time"`
	Value  string    `json:"value"`
	IdNode int       `json:"id_node"`
}
