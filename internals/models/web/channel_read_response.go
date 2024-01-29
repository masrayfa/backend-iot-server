package web

import "time"

type ChannelReadResponse struct {
	Time   time.Time `json:"time"`
	Value  []float64    `json:"value"`
	IdNode int64     `json:"id_node"`
}
