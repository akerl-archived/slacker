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
	return strconv.Itoa(shifted)
}

func trim_leading_zeros(chunk string) string {
	return strings.TrimLeft(chunk, "0")
}

func trim_address(chunks []string) []string {
	// Start by trimming the 1st, 3rd, and 5th chunks,
	// which are always safe to trim
	for i := 0; i < 5; i = i + 2 {
		chunks[i] = trim_leading_zeros(chunks[i])
		// The 2nd and 4th chunks can be trimmed
		// only if their preceding chunk is empty
		if i != 2 && chunks[i] == "" {
			chunks[i+1] = trim_leading_zeros(chunks[i+1])
		}
	}
	return chunks
}

func main() {
	mac_address := parse_args()

	chunks := strings.Split(mac_address.String(), ":")

	chunks[0] = flip_seventh_bit(chunks[0])
	chunks = trim_address(chunks)

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
