package utils

import (
	"strings"

	"github.com/likexian/whois-go"
)

// Returns country, owner, error get by WhoIs for the given address
func SendWhoIsRequest(address string) (string, string, error) {

	var country string
	var owner string

	whoIsResponse, err := whois.Whois(address)

	if err != nil {
		return country, owner, err
	}

	whoisLines := strings.Split(whoIsResponse, "\n")

	for i := 0; i < len(whoisLines); i++ {
		line := strings.TrimSpace(whoisLines[i])

		if !validateLine(line) {
			continue
		}

		lines := strings.SplitN(line, ":", 2)
		name := strings.TrimSpace(lines[0])
		value := strings.TrimSpace(lines[1])

		if name == "Country" {
			country = value
		} else if name == "OrgName" {
			owner = value
		}
	}

	return country, owner, err

}

// Returns true if it's a valid line, otherwise returns false
func validateLine(line string) bool {
	if len(line) < 5 || !strings.Contains(line, ":") || line[len(line)-1:] == ":" {
		return false
	}

	fChar := line[:1]
	if fChar == ">" || fChar == "%" || fChar == "*" || fChar == "#" {
		return false
	}

	return true
}
