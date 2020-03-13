package block_dns

import (
	"os"
	"bufio"
)

type BlacklistDomain struct {
	domains []string
}

func (b BlacklistDomain) IsBlocked(domain string) bool {
	for _, d := range b.domains {
		if d == domain {
			return true
		}
	}
	return false
}

func NewBlacklistDomain() BlacklistDomain {
	file, err := os.Open("domains.txt")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	var blackListedDomains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		blackListedDomains = append(blackListedDomains, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}

	return BlacklistDomain{
		domains: blackListedDomains,
	}
}
