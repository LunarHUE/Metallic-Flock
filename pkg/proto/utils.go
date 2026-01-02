package proto

import (
	"log"
	"net"
	"strconv"
)

func FindOpenPort(defaultPort int) int {
	const maxPort = 32767 // Max positive value for int16

	for port := defaultPort; port <= maxPort; port++ {
		addr := net.JoinHostPort("", strconv.Itoa(port))
		listener, err := net.Listen("tcp", addr)

		// If err is nil, the port is available
		if err == nil {
			listener.Close()
			return port
		}
	}

	return 0
}

func CurrentLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
