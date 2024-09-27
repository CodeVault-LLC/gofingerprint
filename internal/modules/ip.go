package modules

import (
	"fmt"
	"net"
	"net/http"
)

type IpResponse struct {
	IP      string `json:"ip"`
	IsLocal bool   `json:"is_local"`
}

// GetIP extracts the correct IP address from the request, accounting for proxies and reverse proxies.
func GetIP(req *http.Request) IpResponse {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)

	if err != nil {
		fmt.Printf("Unable to parse RemoteAddr: %v\n", err)
		return IpResponse{
			IP:      "",
			IsLocal: false,
		}
	}

	return IpResponse{
		IP:      ip,
		IsLocal: isLocalIP(ip),
	}
}

func isLocalIP(ip string) bool {
	// Loopback addresses
	if ip == "127.0.0.1" || ip == "::1" {
		return true
	}

	// Check if it is a private IP address
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	privateIPBlocks := []*net.IPNet{
		// IPv4 private address blocks
		{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(8, 32)},
		{IP: net.ParseIP("172.16.0.0"), Mask: net.CIDRMask(12, 32)},
		{IP: net.ParseIP("192.168.0.0"), Mask: net.CIDRMask(16, 32)},

		// IPv6 private address blocks
		{IP: net.ParseIP("fc00::"), Mask: net.CIDRMask(7, 128)},
	}

	for _, block := range privateIPBlocks {
		if block.Contains(parsedIP) {
			return true
		}
	}

	return false
}
