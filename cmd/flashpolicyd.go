package main

import (
	"flag"
	"github.com/bybzmt/flashpolicy"
)

var addr = flag.String("addr", ":843", "listen ip:port")
var domains = flag.String("domains", "*", "etc. www.example.com,www.example.dev:81")
var ports = flag.String("ports", "80", "etc. 507,516")

func main() {
	flag.Parse()

	flashpolicy.Server(*addr, *domains, *ports)
}
