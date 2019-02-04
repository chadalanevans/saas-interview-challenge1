package services

import . "github.com/chadalanevans/saas-interview-challenge1/types"

type Launcher interface {
	Launch(work Work) error
}

