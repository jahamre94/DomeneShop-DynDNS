package domeneshop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Domain struct {
	Domain string `json:"domain"`
	ID     int    `json:"id"`
}

func GetDomains() ([]Domain, error) {

	method := "GET"
	url := BaseURL + "/domains"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+Token)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var domains []Domain
	if err := json.NewDecoder(resp.Body).Decode(&domains); err != nil {
		return nil, err
	}

	return domains, nil

}

func VerifyDomain(domains []string) (bool, error) {
	// Create a set of unique root domains from cfg.Domains
	rootDomains := make(map[string]struct{})
	for _, domain := range domains {
		rootDomain := extractRootDomain(domain)
		rootDomains[rootDomain] = struct{}{}
	}

	d, err := GetDomains()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return false, err
	}

	// Create a set of fetched domains
	fetchedDomains := make(map[string]struct{})
	for _, domain := range d {
		fetchedDomains[domain.Domain] = struct{}{}
	}

	// Compare the root domains with fetched domains
	for rootDomain := range rootDomains {
		if _, exists := fetchedDomains[rootDomain]; !exists {
			return false, nil
		}
	}

	return true, nil

}

// Helper function to extract the root domain from a domain or subdomain
func extractRootDomain(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) > 2 {
		return strings.Join(parts[len(parts)-2:], ".")
	}
	return domain
}

func UpdateDynDNS(domain string, ip string) error {

	// /dyndns/update?hostname=<hostname>&myip=<string>
	method := "GET"
	url := BaseURL + "/dyndns/update?hostname=" + domain + "&myip=" + ip

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Basic "+Token)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil

}
