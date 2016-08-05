package FlashSocketPolicyFile

import (
	"io"
	"log"
	"net"
	"strings"
	"time"
)

var xmlfile []byte

func Server(add, domains, ports string) {
	lsn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	buildPolicyXml(domains, ports)

	for {
		if conn, err := lsn.Accept(); err == nil {
			go writePolicyFile(conn)
		}
	}
}

var request = []byte("<policy-file-request/>\x00")

func writePolicyFile(conn net.Conn) {
	defer func() {
		conn.Close()

		e := recover()
		if e != nil {
			log.Println(e)
		}
	}()

	conn.SetDeadline(time.Now().Add(10 * time.Second))

	buff := make([]byte, len(request))

	if n, err := io.ReadFull(conn, buff); err == nil && n == len(request) {
		conn.Write(xmlfile)
	}
}

func buildPolicyXml(domains, addr string) {
	x := strings.Split(addr, ":")

	var port string

	if len(x) == 1 {
		port = "80"
	} else {
		port = x[1]
	}

	ss := "<?xml version=\"1.0\"?>\n"
	ss += "<cross-domain-policy>\n"
	for _, y := range strings.Split(domains, ",") {
		z := strings.Split(y, ":")
		ss += "\t<allow-access-from domain=\"" + z[0] + "\" to-ports=\"" + port + "\" />\n"
	}
	ss += "</cross-domain-policy>"

	xmlfile = []byte(ss)
}
