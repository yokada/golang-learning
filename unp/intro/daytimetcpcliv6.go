package main

import (
	"fmt"
	"flag"
	"log"
	"net"
	"time"
)

/*
Usage example:
	$ go run daytimetcpcliv6.go 2610:20:6F15:15::27 13
	57480 16-04-02 14:23:42 50 0 0 694.7 UTC(NIST) *
You can find valid daytime server URLs and IP Addresses from here:
	[ NIST Internet Time Service ]( http://tf.nist.gov/tf-cgi/servers.cgi )
*/
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("Usage: daytimetcpcli.go <IP ADDRESS> <PORT>")
		return
	}
	serverIp := args[0]
	serverPort := args[1]

	conn, err := net.Dial("tcp6", "["+serverIp+"]:"+serverPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	readBuf := make([]byte, 1024)
	// set timeout duration as absolute soconds from now for reading data.
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	readLen, err := conn.Read(readBuf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(readBuf[:readLen]))
}
