package domain

import "time"

type Channel struct {
	Time time.Time `json:"time"`
	Value string `json:"value"`
	IdNode int `json:"id_node"`
}