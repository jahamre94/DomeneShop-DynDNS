package network

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func GetPublicIP() (string, string, error) {
	// API to get the public IP address
	url := "https://api.ipify.org?format=text"

	// Make an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("failed to get public IP: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	ipBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response: %v", err)
	}

	// Convert the response to a string
	ipStr := string(ipBytes)

	// Validate the IP address
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "", "", fmt.Errorf("invalid IP address format: %s", ipStr)
	}

	// Determine if it's IPv4 or IPv6
	ipType := "unknown"
	if ip.To4() != nil {
		ipType = "IPv4"
	} else if ip.To16() != nil {
		ipType = "IPv6"
	}

	return ipStr, ipType, nil
}
