package describe

import (
	"net"
)

// ValidIP check string is valid IP or not by net.ParseIP
func ValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsPrivateIP if addr string is in RFC1918 private network or not
func IsPrivateIP(ip net.IP) bool {
	ip = ip.To4()
	if ip == nil {
		return false
	}
	return (ip[0] == 10) ||
		(ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31) ||
		(ip[0] == 192 && ip[1] == 168)
}

// IsPublicIP if addr string is in public network or not
func IsPublicIP(ip net.IP) bool {
	return !IsPrivateIP(ip) && ip.IsGlobalUnicast()
}
