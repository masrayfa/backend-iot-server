package domain

import "time"

type Channel struct {
	Time time.Time `json:"time"`
	Value  []float64`json:"value"`
	IdNode int64 `json:"id_node"`
}