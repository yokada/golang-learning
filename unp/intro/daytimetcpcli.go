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
	$ go run daytimetcpcli.go 98.175.203.200 13
	57480 16-04-02 13:33:59 50 0 0 905.5 UTC(NIST) *
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
	//clientIp := "127.0.0.1"
	//clientPort := 8889

	// create tcp socket
	/*
	serverAddr, err := net.ResolveTCPAddr("tcp", serverIp+":"+serverPort)
	if err != nil {
		log.Fatal("cannot resolve server addr")
	}
	*/

	/*
	clientAddr := new(net.TCPAddr)
	clientAddr.IP = net.ParseIP(clientIp)
	clientAddr.Port = clientPort
	*/

	/*
	conn, err := net.DialTCP("tcp", clientAddr, serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	*/
	conn, err := net.Dial("tcp", serverIp+":"+serverPort)
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
