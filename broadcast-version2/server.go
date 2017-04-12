package main

import (
	"encoding/binary"
	"gopkg.in/dedis/onet.v1/log"
	"net"
	"strconv"
	"time"
)

const UDP_PORT int = 10101
const MAX_UDP_SIZE int = 65507
const LOCALADDR string = "10.0.1.254:0"
const DESTADDR string = "10.0.1.255"

func main() {
	log.Info("Running broadcast server...")

	LocalAddr, err := net.ResolveUDPAddr("udp", LOCALADDR)
	if err != nil {
		log.Error("Broadcast: could not resolve Local address, error is", err.Error())
	}

	ServerAddr, err := net.ResolveUDPAddr("udp", DESTADDR+":"+strconv.Itoa(UDP_PORT))
	if err != nil {
		log.Error("Broadcast: could not resolve BCast address, error is", err.Error())
	}

	conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		log.Fatal("Broadcast: could not UDP Dial, error is", err.Error())
	}

	for {
		log.Info("Broadcast: Sending one message...")
		message := make([]byte, 4+2)
		binary.BigEndian.PutUint32(message[0:4], uint32(2))     //size
		binary.BigEndian.PutUint16(message[4:6], uint16(43690)) //magic number

		_, err = conn.Write(message)
		if err != nil {
			log.Error("Broadcast: could not write message, error is", err.Error())
		} else {
			log.Lvl3("Broadcast: broadcasted one message of length", len(message))
		}

		time.Sleep(10 * time.Second)
	}
}
