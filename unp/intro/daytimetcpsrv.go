package main

import (
	"log"
	"net"
	"time"
)

/*
Usage example:
	$ go run daytimetcpcli.go 98.175.203.200 13
	57480 16-04-02 13:33:59 50 0 0 905.5 UTC(NIST)

You can find valid daytime server URLs and IP Addresses from here:
	[ NIST Internet Time Service ]( http://tf.nist.gov/tf-cgi/servers.cgi )

daytime protocol reference:
	[ RFC 867 - Daytime Protocol ]( https://tools.ietf.org/html/rfc867 )
*/
func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleListener(conn)
	}
}


func handleListener(conn net.Conn) {
	defer conn.Close()
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	t := time.Now()
	//msg := t.Format("Mon Jan 2 15:04:05 2006 -0700 MST")
	msg := t.Format(time.RFC1123)
	conn.Write([]byte(msg))
}