package describe

import (
	"net"
)

// ValidIP check string is valid IP or not by net.ParseIP
func ValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
