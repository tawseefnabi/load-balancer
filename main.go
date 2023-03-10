package main

import (
	"flag"
	"fmt"

	"github.com/tawseefnabi/load-balancer/internal/lb"
	"github.com/tawseefnabi/load-balancer/internal/utils"
)

var (
	flagURL utils.FlagURL
	port    int
)

func init() {
	flag.Var(&flagURL, "servers", "Use commas to separate")
	flag.IntVar(&port, "port", 8090, "Port to serve")
	flag.Parse()
	// if len(flagURL.URLs) == 0 {
	// 	log.Fatal("Please provide one or more servers to load balance")
	// }
}
func main() {
	fmt.Println("simple lb")
	l := lb.New()
	l.Register(flagURL.URLs...)
}
