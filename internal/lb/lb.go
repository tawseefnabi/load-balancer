package lb

import (
	"fmt"

	"github.com/tawseefnabi/load-balancer/internal/server"
)

type LoadBalancer struct {
	controller *server.Controller
}

func New() *LoadBalancer {
	fmt.Println("New func from lb called...")
	return &LoadBalancer{
		controller: server.NewController(),
	}
}
