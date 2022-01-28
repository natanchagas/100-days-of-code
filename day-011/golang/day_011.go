package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IsValidIP(address string) bool {
	blocks := strings.Split(address, ".")

	for _, block := range blocks {
		if len(block) > 1 && block[0:1] == "0" {
			return false
		}

		value, err := strconv.Atoi(block)

		if err != nil {
			return false
		} else if value > 255 {
			return false
		}

	}

	return true
}

func GeIPAddresses(address string) []string {

	var addresses []string

	if len(address) < 4 || len(address) > 12 {
		return []string{}
	}

	for i := 0; i < 3; i++ {
		for j := i; j < len(address) && j <= i+3; j++ {
			for k := j; k < len(address) && k <= j+3; k++ {
				for l := k; l < len(address) && l <= k+3; l++ {
					ip_address := fmt.Sprintf("%v%v%v%v%v%v%v", address[:i+1], ".", address[i+1:j+1], ".", address[j+1:k+1], ".", address[k+1:l+1])
					if len(ip_address) == len(address)+3 {
						if IsValidIP(ip_address) {
							addresses = append(addresses, ip_address)
						}
					}

				}
			}
		}
	}

	return addresses
}
