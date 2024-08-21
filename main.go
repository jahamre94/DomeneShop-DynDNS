package main

import (
	"fmt"
	cfg "main/config"
	"main/domeneshop"
	"main/network"
)

func main() {

	cfg := cfg.Load()
	domeneshop.Init(cfg)

	// Verify the domains
	if ok, err := domeneshop.VerifyDomain(cfg.Domains); ok {
		fmt.Println("All domains are verified.")
	} else {
		fmt.Printf("Error: %v\n", err)
	}

	ip, _, err := network.GetPublicIP()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for _, domain := range cfg.Domains {
		fmt.Printf("Updating DynDNS for domain %s\n", domain)
		err := domeneshop.UpdateDynDNS(domain, ip)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

}
