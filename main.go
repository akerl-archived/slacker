package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func parse_args() net.HardwareAddr {
	if len(os.Args[1:]) < 1 {
		fmt.Printf("Usage: %s MAC_ADDRESS\n", os.Args[0])
		os.Exit(1)
	}
	mac_address, err := net.ParseMAC(os.Args[1])
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	return mac_address
}

func flip_seventh_bit(chunk string) string {
	chunk_as_int, _ := strconv.Atoi(chunk)
	shifted := chunk_as_int ^ (1 << 1)
	new_chunk_as_string := strconv.Itoa(shifted)
	return new_chunk_as_string
}

func main() {
	mac_address := parse_args()

	chunks := strings.Split(mac_address.String(), ":")
	chunks[0] = flip_seventh_bit(chunks[0])

	fmt.Printf(
		"%s%s:%sff:fe%s:%s%s\n",
		chunks[0],
		chunks[1],
		chunks[2],
		chunks[3],
		chunks[4],
		chunks[5],
	)
}
