package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Printf("Usage: %s MAC_ADDRESS\n", os.Args[0])
		os.Exit(1)
	}
	mac_address, err := net.ParseMAC(os.Args[1])
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	chunks := strings.Split(mac_address.String(), ":")
	fmt.Printf(
		"%s%s:%sff:fe%s:%s%s\n",
		chunks[0],
		chunks[1],
		chunks[2],
		chunks[3],
		chunks[4],
		chunks[5],
	)

	a := 0x04
	fmt.Printf("%b\n", a)
}
