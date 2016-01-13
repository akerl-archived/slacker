package main

import (
	"fmt"
	"net"
	"os"
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
	fmt.Printf("%s\n", mac_address)
}
