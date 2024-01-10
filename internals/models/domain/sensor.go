package domain

type Sensor struct {
	IdSensor   int
	Name       string
	IdNode     int
	IdHardware int
}

type SensorWithChannel struct {
	Sensor  Sensor
	Channel []Channel
}
