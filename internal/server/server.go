package server

import (
	"container/list"
)

type Controller struct {
	upIDs   *list.List
	downIDs *list.List
	// mux     sync.Mutex
}

func NewController() *Controller {
	return &Controller{
		upIDs:   list.New(),
		downIDs: list.New(),
	}
}
