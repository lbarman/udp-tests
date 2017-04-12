package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"gopkg.in/dedis/onet.v1/log"
	"net"
	"strconv"
)

const UDP_PORT int = 10101
const MAX_UDP_SIZE int = 65507
const LOCALADDR string = "10.0.1.254:0"
const DESTADDR string = "10.0.1.255"

func main() {
	var ip = flag.String("ip", "10.0.1.1", "listen on this IP")
	flag.Parse()

	listeningAddr := *ip

	log.Info("Running listening client on IP", listeningAddr+":"+strconv.Itoa(UDP_PORT), "...")

	ServerAddr, err := net.ResolveUDPAddr("udp", listeningAddr+":"+strconv.Itoa(UDP_PORT))
	if err != nil {
		log.Error("Listener: could not resolve BCast address, error is", err.Error())
	}

	conn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		log.Fatal("Listener: could not UDP Dial, error is", err.Error())
	}

	for {
		buf := make([]byte, MAX_UDP_SIZE)
		log.Info("Listener: Ready to receive")

		n, addr, err := conn.ReadFromUDP(buf)
		log.Info("Listener: Received a header from", addr, "gonna read message of length...", n, "size is", len(buf))
		sizeAdvertised := int(binary.BigEndian.Uint32(buf[0:4]))

		if sizeAdvertised+4 != n {
			log.Error("Listener: could not receive read the ", string(sizeAdvertised+4), ", only", n, ", error is", err.Error())
		}
		message := make([]byte, sizeAdvertised)
		copy(message[:], buf[4:sizeAdvertised+4])

		log.Info(hex.Dump(message))
	}
}
