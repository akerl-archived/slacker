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

func flip_seventh_bit(mac_segment string) string {
	mac_segment_as_int, _ := strconv.Atoi(mac_segment)
	shifted := mac_segment_as_int ^ (1 << 1)
	return strconv.Itoa(shifted)
}

func trim_leading_zeros(mac_segment string) string {
	return strings.TrimLeft(mac_segment, "0")
}

func trim_address(mac_segments []string) []string {
	// Start by trimming the 1st, 3rd, and 5th mac_segments,
	// which are always safe to trim
	for i := 0; i < 5; i = i + 2 {
		mac_segments[i] = trim_leading_zeros(mac_segments[i])
		// The 2nd and 4th mac_segments can be trimmed
		// only if their preceding mac_segment is empty
		if i != 2 && mac_segments[i] == "" {
			mac_segments[i+1] = trim_leading_zeros(mac_segments[i+1])
		}
	}
	return mac_segments
}

func main() {
	mac_address := parse_args()

	mac_segments := strings.Split(mac_address.String(), ":")

	mac_segments[0] = flip_seventh_bit(mac_segments[0])
	mac_segments = trim_address(mac_segments)

	fmt.Printf(
		"%s%s:%sff:fe%s:%s%s\n",
		mac_segments[0],
		mac_segments[1],
		mac_segments[2],
		mac_segments[3],
		mac_segments[4],
		mac_segments[5],
	)
}
