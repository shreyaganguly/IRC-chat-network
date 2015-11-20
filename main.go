package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

type IRCServer struct {
	Addr string
}
type Client struct {
	RemoteAddress string
	UserName      string
	HostIP        string
	RealName      string
}
type ChannelDetails struct {
	Topic string
}

var ClientMap map[string]Client
var ChannelMap map[string]ChannelDetails

func NewIRCServer(host string, port int) *IRCServer {
	return &IRCServer{Addr: fmt.Sprintf("%s:%d", host, port)}
}

func (s *IRCServer) ListenAndServe() {

	ln, err := net.Listen("tcp", s.Addr)

	if err != nil {
		log.Fatal("ERROR:", err)
	}

	for {
		con, err := ln.Accept()

		if err != nil {
			log.Println("ERROR:", err)
		}

		go ConnectionHandler(con)

	}

}

var (
	host = flag.String("b", "0.0.0.0", "Host name of the TCP Server")
	port = flag.Int("t", 0, "port to listen for connections")
)

func main() {
	ClientMap = make(map[string]Client)
	ChannelMap = make(map[string]ChannelDetails)
	log.Println("Starting IRC Server..")
	flag.Parse()
	server := NewIRCServer(*host, *port)
	server.ListenAndServe()
}
