package countIPAddresses

import (
	"fmt"
	"strconv"
	"strings"
)

func result(ip string) uint32 {
	octetMultipliers := []uint32{16777216, 65536, 256, 1}

	var startIPParts []uint32

	parts := strings.Split(ip, ".")
	for i, part := range parts {
		intValue, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error al convertir octeto:", err)
			return 0
		}
		segmentValue := uint32(intValue) * octetMultipliers[i]
		startIPParts = append(startIPParts, segmentValue)
	}

	var totalSum uint32
	for _, part := range startIPParts {
		totalSum += part
	}
	return totalSum
}

func IpsBetween(start string, end string) uint32 {
	return result(end) - result(start)
}
