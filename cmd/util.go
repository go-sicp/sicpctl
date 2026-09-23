package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

// parseByte accepts decimal ("42") or hexadecimal ("0x2A", "0X2a") notation
// and returns a byte.
func parseByte(s string) (byte, error) {
	base := 10
	t := s
	if low := strings.ToLower(s); strings.HasPrefix(low, "0x") {
		base = 16
		t = low[2:]
	}
	v, err := strconv.ParseUint(t, base, 16)
	if err != nil {
		return 0, fmt.Errorf("not a number: %q", s)
	}
	if v > 0xFF {
		return 0, fmt.Errorf("value %s exceeds one byte", s)
	}
	return byte(v), nil
}
