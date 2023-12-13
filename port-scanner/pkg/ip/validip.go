package ip

import (
	"strconv"
	"strings"
)

func IsValidIp(ip string) bool {
	if ip == "localhost" {
		return true
	}

	splitted := strings.Split(ip, ".")
	if len(splitted) != 4 {
		return false
	}

	for _, s := range splitted {
		converted, err := strconv.Atoi(s)
		if err != nil {
			return false
		}

		if converted < 0 || converted > 255 {
			return false
		}
	}

	return true
}
