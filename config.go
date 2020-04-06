package main

// Config holds program configuration
type Config struct {
	Min           float64
	Max           float64
	Generator     string
	MessageIDFile string
	Sensor        string
	Unit          string
	MQTTEndpoint  string
	ClientID      string
	Topic         string
	Username      string
	Password      string
}
