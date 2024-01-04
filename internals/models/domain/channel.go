package domain

import "time"

type Channel struct {
	Time time.Time `json:"time"`
	Value string `json:"value"`
	IdSensor int `json:"id_sensor"`
}