package main

import (
	"flag"
	"github.com/bybzmt/FlashSocketPolicyFile"
	"log"
	"net/http"
	"strings"
)

var addr = flag.String("addr", ":843", "listen ip:port")
var domains = flag.String("domains", "*", "etc. www.example.com,www.example.dev:81")
var ports = flag.String("ports", "80", "etc. 507,516")

func main() {
	flag.Parse()

	FlashSocketPolicyFile.Server(*addr, *domains, *ports)
}
