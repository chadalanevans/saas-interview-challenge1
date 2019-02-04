package pubsub

import . "github.com/chadalanevans/saas-interview-challenge1/types"

type Publisher interface {
	Initialize(url string) error
	Publish(event Event) error
	Finish() error
}